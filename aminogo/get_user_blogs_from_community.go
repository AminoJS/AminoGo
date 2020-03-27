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

type GetUserBlogFromComOptions struct {
	CommunityID int
	UUID        string
	Start       int
	Size        int
}

// Get a list of blogs that are written from a selected users from a selected community
func GetUserBlogsFromCommunity(argument *GetUserBlogFromComOptions) (blogsFeed *structs.BlogsFromCommunity, err error) {
	SID := stores.Get("SID")
	if SID == nil {
		return &structs.BlogsFromCommunity{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	if argument.CommunityID == 0 {
		return &structs.BlogsFromCommunity{}, errors.New("CommunityID cannot be 0 or empty")
	}

	endpoint := routes.GetUserBlogsFromCommunity(argument.CommunityID, argument.UUID, argument.Start, argument.Size)

	utils.DebugLog("get_user_blogs_from_community.go", fmt.Sprintf("URL: %s", endpoint))

	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}

	req.SetTimeout(30 * time.Second)
	res, err := req.Get(endpoint, header)
	if err != nil {
		return &structs.BlogsFromCommunity{}, err
	}
	err = utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.BlogsFromCommunity{}, err
	}

	resMap := structs.BlogsFromCommunity{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.BlogsFromCommunity{}, err
	}

	return &resMap, nil
}
