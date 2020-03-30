package aminogo

import (
	"errors"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/utils"
)

type GetJoinedChatroomsOptions struct {
	CommunityID int
	Start       int
	Size        int
}

// Get a list of joined chat rooms based ona selected community
func GetJoinedChatrooms(argument *GetJoinedChatroomsOptions) (joinedChatrooms *structs.JoinedChatrooms, err error) {
	SID := stores.Get("SID")
	if SID == nil || SID == "" {
		return &structs.JoinedChatrooms{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	endpoint := routes.GetJoinedChatrooms(argument.CommunityID, argument.Start, argument.Size)

	res, err := utils.Get(endpoint)
	if err != nil {
		return &structs.JoinedChatrooms{}, err
	}

	resMap, err := utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.JoinedChatrooms{}, err
	}
	tmp := resMap.(structs.JoinedChatrooms)
	return &tmp, nil

}
