package aminogo

import (
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/utils"
	"github.com/imroc/req"
)

type GetUserBlogFromComOptions struct {
	CommunityID int
	UUID        string
	Start       int
	Size        int
}

// Get a list of bloggs that are written from a selected users from a selected community
func GetUserBlogsFromCommunity(argument *GetUserBlogFromComOptions) (blogsFeed *structs.GetUserBlogsFromCommunity, err error) {
	SID := stores.Get("SID")
	if SID == nil {
		return &structs.GetUserBlogsFromCommunity{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	endpoint := routes.GetUserBlogsFromCommunity(argument.CommunityID, argument.UUID, argument.Start, argument.Size)

	utils.DebugLog("get_user_blogs_from_community.go", fmt.Sprintf("URL: %s", endpoint))

	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}

	res, err := req.Get(endpoint, header)
	if err != nil {
		return &structs.GetUserBlogsFromCommunity{}, err
	}

	resMap := structs.GetUserBlogsFromCommunity{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.GetUserBlogsFromCommunity{}, err
	}

	return &resMap, nil
}
