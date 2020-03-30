package aminogo

import (
	"errors"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/utils"
)

// Get a detailed chat log from a selected chat room
func GetChat(communityID int, threadID string) (*structs.ChatRecords, error) {
	SID := stores.Get("SID")
	if SID == nil || SID == "" {
		return &structs.ChatRecords{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	endpoint := routes.GetChat(communityID, threadID)

	res, err := utils.Get(endpoint)
	if err != nil {
		return &structs.ChatRecords{}, err
	}

	resMap, err := utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.ChatRecords{}, err
	}
	tmp := resMap.(structs.ChatRecords)
	return &tmp, err

}
