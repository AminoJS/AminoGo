package aminogo

import (
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/utils"
	"github.com/imroc/req"
	"strings"
	"time"
)

// Post a blog post to a selected community
// Post a blog post to a selected community
func PostBlog(communityID int, title string, content string, mediaList *[]*MediaContainer) (*structs.PostedBlog, error) {

	SID := stores.Get("SID")
	if SID == nil {
		return &structs.PostedBlog{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	if title == "" {
		return &structs.PostedBlog{}, errors.New("post title cannot be empty")
	}

	if content == "" {
		return &structs.PostedBlog{}, errors.New("post content cannot be empty")
	}

	endpoint := routes.PostBlog(communityID)

	header := req.Header{
		"NDCAUTH":      fmt.Sprintf("sid=%s", SID),
		"Content-Type": "application/json",
	}

	var data = make(map[string]interface{})

	data["content"] = content
	data["latitude"] = 0
	data["longitude"] = 0
	data["title"] = title
	data["clientRefId"] = 43196704
	data["eventSource"] = "eventSource"
	data["timestamp"] = time.Now().Unix()

	// Replace reference key
	var postMediaList []interface{}

	if len(*mediaList) >= 1 {
		for _, media := range *mediaList {
			var image [4]interface{}

			image[0] = 100
			image[1] = media.FinalDes
			if media.Captions != "" {
				image[2] = media.Captions
			}
			if media.referenceKey != "" {
				data["content"] = strings.ReplaceAll(content, media.referenceKey, fmt.Sprintf("[IMG=%s]", media.referenceKey))
				image[3] = media.referenceKey
			}
			postMediaList = append(postMediaList, image)
		}

		data["mediaList"] = postMediaList
	}

	req.SetTimeout(30 * time.Second)
	res, err := req.Post(endpoint, header, req.BodyJSON(data))
	if err != nil {
		return &structs.PostedBlog{}, err
	}

	resMap, err := utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.PostedBlog{}, err
	}
	tmp := resMap.(structs.PostedBlog)
	return &tmp, nil

}
