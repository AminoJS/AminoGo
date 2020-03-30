package aminogo

import (
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/utils"
)

// Return a complete REST respond as a struct, all the info are contain inside the "Account" field
func MyProfile() (profile *structs.MyProfile, err error) {
	SID := stores.Get("SID")

	if SID == nil || SID == "" {
		return &structs.MyProfile{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	endpoint := routes.MyProfile()

	res, err := utils.Get(endpoint)
	if err != nil {
		return &structs.MyProfile{}, err
	}

	resMap, err := utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.MyProfile{}, err
	}
	tmp := resMap.(structs.MyProfile)
	fmt.Printf("Hello")
	return &tmp, nil
}
