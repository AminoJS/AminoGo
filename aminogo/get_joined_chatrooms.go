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

type GetJoinedChatroomsOptions struct {
	CommunityID int
	Start       int
	Size        int
}

// Get a list of joined chat rooms based ona selected community
func GetJoinedChatrooms(argument *GetJoinedChatroomsOptions) (joinedChatrooms *structs.JoinedChatrooms, err error) {
	SID := stores.Get("SID")
	if SID == nil {
		return &structs.JoinedChatrooms{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}

	endpoint := routes.GetJoinedChatrooms(argument.CommunityID, argument.Start, argument.Size)

	utils.DebugLog("get_joined_chatrooms.go", fmt.Sprintf("URL: %s", endpoint))

	res, err := req.Get(endpoint, header)
	if err != nil {
		return &structs.JoinedChatrooms{}, err
	}

	//fmt.Println(res.ToString())

	resMap := structs.JoinedChatrooms{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.JoinedChatrooms{}, err
	}
	err = utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.JoinedChatrooms{}, err
	}

	return &resMap, nil

}
