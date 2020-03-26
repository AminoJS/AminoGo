package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func formReturnMessage(msg string, aminoMsg interface{}) error {
	msg += fmt.Sprintf("\nAmino Said: %v", aminoMsg)
	return errors.New(msg)
}

// Return a error message for some common HTTP error with some tips
func ThrowHttpErrorIfFail(res *http.Response) (err error) {

	if res.StatusCode == 0 {
		return errors.New("empty or 0 HTTP status is not allowed")
	}

	var aminoMsg interface{}

	if res.Body != nil {
		var body map[string]interface{}

		jStr, _ := ioutil.ReadAll(res.Body)

		if string(jStr) != "" {
			json.Unmarshal(jStr, body)

			aminoMsg = body["api:message"].(interface{})
		} else {
			aminoMsg = "(AminoGo Could Not Parse Error Message)"
		}
	}

	switch res.StatusCode {
	case 404:
		return formReturnMessage("fail to login API call due to resource not found, resulted in a none 404 status code", aminoMsg)
	case 400:
		return formReturnMessage("fail to login API call due to bad request (perhaps you are giving the wrong arguments), resulted in a none 400 status code", aminoMsg)
	case 401:
		return formReturnMessage("fail to login API call due to unauthorized (perhaps you could try re-login thus generating a new session-token), resulted in a none 400 status code", aminoMsg)
	case 405:
		return formReturnMessage("fail to login API call due to bad method not allowed (there must be a bug in AminoGo, please ensure AminoGo are update-to-date, if you are on the latest version, please repo this on GitHub at https://github.com/AminoJS/AminoGo/issues), resulted in a none 400 status code", aminoMsg)
	case 500:
		return formReturnMessage("fail to login API call due to internal server error(not your fault), resulted in a none 500 status code", aminoMsg)
	}
	if res.StatusCode != 200 && res.StatusCode != 201 {
		return formReturnMessage(fmt.Sprintf("fail to login API call, resulted in a none %d status code", res.StatusCode), aminoMsg)
	}
	return nil
}
