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

// Get a detailed chat log from a selected chat room
func GetChat(communityID int, threadID string) (*structs.ChatRecords, error) {
	SID := stores.Get("SID")
	if SID == nil || SID == "" {
		return &structs.ChatRecords{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}

	endpoint := routes.GetChat(communityID, threadID)

	utils.DebugLog("get_chat.go", fmt.Sprintf("URL: %s", endpoint))
	req.SetTimeout(30 * time.Second)
	res, err := req.Get(endpoint, header)
	if err != nil {
		return &structs.ChatRecords{}, err
	}

	resMap := structs.ChatRecords{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.ChatRecords{}, err
	}
	err = utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.ChatRecords{}, err
	}

	return &resMap, err

}
