package aminogo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"io/ioutil"
	"net/http"
	"time"
)

// Return a complete REST respond as a struct, all the info are contain inside the "Account" field
func MyProfile() (profile structs.MyProfile, err error) {
	SID := stores.Get("SID")
	endpoint := routes.GetRoutes()["MyProfile"]
	if SID == nil {
		return structs.MyProfile{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return structs.MyProfile{}, err
	}
	req.Header.Add("NDCAUTH", fmt.Sprintf("sid=%s", SID))
	client := &http.Client{Timeout: time.Second * 10}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return structs.MyProfile{}, err
	}

	var bodyMap structs.MyProfile

	jStr, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return structs.MyProfile{}, err
	}

	err = json.Unmarshal(jStr, &bodyMap)
	if err != nil {
		return structs.MyProfile{}, err
	}

	return bodyMap, nil
}
