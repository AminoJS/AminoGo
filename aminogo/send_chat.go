package aminogo

import (
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/utils"
	"time"
)

// Send a text message to one of the selected chat room
func SendChat(CommunityID int, ThreadID string, Message string) (*structs.SentChat, error) {
	SID := stores.Get("SID")
	if SID == nil || SID == "" {
		return &structs.SentChat{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	endpoint := routes.SendChat(CommunityID, ThreadID)

	utils.DebugLog("get_blog_feed.go", fmt.Sprintf("URL: %s", endpoint))

	data := make(map[string]interface{})

	data["content"] = Message
	data["type"] = 0
	data["clientRefId"] = 43196704
	data["timestamp"] = time.Now().UTC().Unix()

	res, err := utils.PostJSON(endpoint, data)
	if err != nil {
		return &structs.SentChat{}, err
	}
	err = utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.SentChat{}, err
	}

	resMap := structs.SentChat{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.SentChat{}, err
	}

	return &resMap, nil
}
