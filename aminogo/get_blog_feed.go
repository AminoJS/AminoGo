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

// Get a list of blogs from a targeted community
func GetBlogFeed(communityID int, start int, size int) (blogsFeed *structs.CommunityBlogsFeed, err error) {
	SID := stores.Get("SID")
	if SID == nil || SID == "" {
		return &structs.CommunityBlogsFeed{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	endpoint := routes.GetCommunityBlogs(communityID, start, size)

	utils.DebugLog("get_blog_feed.go", fmt.Sprintf("URL: %s", endpoint))

	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}

	req.SetTimeout(30 * time.Second)
	res, err := req.Get(endpoint, header)
	if err != nil {
		return &structs.CommunityBlogsFeed{}, err
	}
	err = utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.CommunityBlogsFeed{}, err
	}

	resMap := structs.CommunityBlogsFeed{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.CommunityBlogsFeed{}, err
	}

	return &resMap, nil
}
