package aminogo

import (
	"AminoJS/AminoGo/routes"
	"AminoJS/AminoGo/stores"
	"AminoJS/AminoGo/structs"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func MyProfile() (profile structs.MyProfile, err error) {
	SID := stores.Get("SID")
	endpoint := routes.GetRoutes()["MyProfile"]
	if SID == nil {
		return structs.MyProfile{}, errors.New("Missing SID in state, try using aminogo.Login() first")
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

	json.Unmarshal(jStr, &bodyMap)

	return bodyMap, nil
}
