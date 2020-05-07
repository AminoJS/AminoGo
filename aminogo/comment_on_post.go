package aminogo

import (
	"errors"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/utils"
	"time"
)

// Send a text message to one of the selected chat room
func CommentingOnPost(CommunityID int, BlogID string, Message string) (*structs.PostedComment, error) {
	SID := stores.Get("SID")
	if SID == nil || SID == "" {
		return &structs.PostedComment{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	if CommunityID == 0 {
		return &structs.PostedComment{}, errors.New("argument CommunityID cannot be empty")
	}

	if BlogID == "" {
		return &structs.PostedComment{}, errors.New("argument ThreadID cannot be empty")
	}

	if Message == "" {
		return &structs.PostedComment{}, errors.New("argument Message cannot be empty")
	}
	endpoint := routes.CommentingOnPost(CommunityID, BlogID)

	data := make(map[string]interface{})

	data["content"] = Message
	data["mediaList"] = []interface{}{}
	data["eventSource"] = "PostDetailView"
	data["timestamp"] = time.Now().UTC().Unix()

	res, err := utils.PostJSON(endpoint, data)
	if err != nil {
		return &structs.PostedComment{}, err
	}
	err = utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.PostedComment{}, err
	}

	resMap := structs.PostedComment{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.PostedComment{}, err
	}

	return &resMap, nil
}
