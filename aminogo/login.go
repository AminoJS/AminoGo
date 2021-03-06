package aminogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/utils"
	"github.com/imroc/req"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	deviceID        = "015051B67B8D59D0A86E0F4A78F47367B749357048DD5F23DF275F05016B74605AAB0D7A6127287D9C"
	clientType int8 = 100
	action          = "normal"
)

// Get authorize, and returns a session token
func Login(email string, password string) error {

	if email == "" {
		return errors.New("email address MUST be provided as a argument of this function call")
	}

	if password == "" {
		return errors.New("password MUST be provided as a argument of this function call")
	}

	// Create a new map for the post body

	postAuthBody := make(map[string]interface{})

	postAuthBody["email"] = email
	postAuthBody["secret"] = fmt.Sprintf("0 %s", password)
	postAuthBody["clientType"] = clientType
	postAuthBody["deviceID"] = deviceID
	postAuthBody["action"] = action
	postAuthBody["timestamp"] = time.Now().Unix()
	postAuthBody["version"] = 2
	jStr, _ := json.Marshal(postAuthBody)
	data := bytes.NewReader(jStr)

	req.SetTimeout(30 * time.Second)
	res, err := http.Post(routes.Login(), "application/json", data)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var bodyMap map[string]interface{}
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		return err
	}

	err = utils.ThrowHttpErrorIfFail(res)
	if err != nil {
		return err
	}

	SID := bodyMap["sid"].(string)
	stores.Set("SID", SID)

	utils.DebugLog("login.go", fmt.Sprintf("SID %s", SID))

	return nil
}
