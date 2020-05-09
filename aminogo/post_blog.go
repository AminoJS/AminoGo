package aminogo

import (
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/utils"
	"github.com/imroc/req"
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

	if len(*mediaList) >= 1 {
		var tmp []*utils.Media
		for _, media := range *mediaList {
			var img = utils.Media{
				URL:          media.FinalDes,
				Caption:      media.Captions,
				ReferenceKey: media.referenceKey,
			}
			tmp = append(tmp, &img)
		}
		mediaList, newContext := utils.CreateMediaList(tmp, content)
		if *newContext != "" {
			data["content"] = *newContext
		}
		data["mediaList"] = *mediaList
	}

	req.SetTimeout(30 * time.Second)
	res, err := req.Post(endpoint, header, req.BodyJSON(data))
	if err != nil {
		return &structs.PostedBlog{}, err
	}
	err = utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.PostedBlog{}, err
	}

	resMap := structs.PostedBlog{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.PostedBlog{}, err
	}

	return &resMap, nil

}
