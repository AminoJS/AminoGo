package aminogo

import (
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/utils"
	"time"

	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/imroc/req"
)

type GetJoinedCommunitiesOptions struct {
	Start int
	Size  int
}

// Get a list of user's joined communities
func GetJoinedCommunities(argument *GetJoinedCommunitiesOptions) (joinedCommunities *structs.JoinedCommunities, err error) {
	SID := stores.Get("SID")
	if SID == nil || SID == "" {
		return &structs.JoinedCommunities{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	// Send the actual request to the Amino API endpoint

	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}

	endpoint := routes.JoinedCommunities(argument.Start, argument.Size)

	utils.DebugLog("get_joined_communities.go", fmt.Sprintf("URL: %s", endpoint))

	req.SetTimeout(30 * time.Second)
	res, err := req.Get(endpoint, header)
	if err != nil {
		return &structs.JoinedCommunities{}, err
	}

	resMap := structs.JoinedCommunities{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.JoinedCommunities{}, err
	}
	err = utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.JoinedCommunities{}, err
	}

	return &resMap, nil

}
