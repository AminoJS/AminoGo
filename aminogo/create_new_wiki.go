package aminogo

import (
	"errors"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/utils"
	"time"
)

// Get a list of blogs from a targeted community
func CreateNewWiki(communityID int, title string, content string, icon string) (*structs.CreatedWikiEntry, error) {
	SID := stores.Get("SID")
	if SID == nil || SID == "" {
		return &structs.CreatedWikiEntry{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	endpoint := routes.CreateWiki(communityID)

	data := make(map[string]interface{})

	data["content"] = content
	data["label"] = title
	data["eventSource"] = "GlobalComposeMenu"
	data["eventSource"] = icon
	data["timestamp"] = time.Now().UTC().Unix()

	mediaList := []interface{}{}
	var image [2]interface{}
	image[0] = 100
	image[1] = icon

	mediaList = append(mediaList, image)
	data["mediaList"] = mediaList

	res, err := utils.PostJSON(endpoint, data)
	if err != nil {
		return &structs.CreatedWikiEntry{}, err
	}
	err = utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.CreatedWikiEntry{}, err
	}

	resMap := structs.CreatedWikiEntry{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.CreatedWikiEntry{}, err
	}

	return &resMap, nil
}
