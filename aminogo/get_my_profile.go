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

// Return a complete REST respond as a struct, all the info are contain inside the "Account" field
func MyProfile() (profile *structs.MyProfile, err error) {
	SID := stores.Get("SID")

	if SID == nil || SID == "" {
		return &structs.MyProfile{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}

	endpoint := routes.MyProfile()

	utils.DebugLog("get_my_profile.go", fmt.Sprintf("URL: %s", endpoint))

	req.SetTimeout(30 * time.Second)
	res, err := req.Get(endpoint, header)
	if err != nil {
		return &structs.MyProfile{}, err
	}

	resMap := structs.MyProfile{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.MyProfile{}, err
	}
	err = utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.MyProfile{}, err
	}

	return &resMap, nil
}
