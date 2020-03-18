package aminogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	deviceID   string = "015051B67B8D59D0A86E0F4A78F47367B749357048DD5F23DF275F05016B74605AAB0D7A6127287D9C"
	clientType int8   = 100
	action     string = "normal"
)

// Get authorize, and returns a session token
func Login(email string, password string) error {
	var endpoint string = routes.GetRoutes()["Login"]
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

	res, err := http.Post(endpoint, "application/json", data)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var bodyMap map[string]interface{}
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		return err
	}

	SID := bodyMap["sid"].(string)
	stores.Set("SID", SID)
	return nil
}
