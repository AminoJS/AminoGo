package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func formReturnMessage(msg string, res *http.Response) (interface{}, error) {

	var aminoMsg interface{}
	var body map[string]interface{}

	if res.Body != nil {

		jStr, _ := ioutil.ReadAll(res.Body)

		if string(jStr) != "" && jStr != nil {
			err := json.Unmarshal(jStr, &body)
			if err != nil {
				aminoMsg = "(AminoGo Could Not Parse Error Message)"
			} else {
				aminoMsg = body["api:message"].(interface{})
			}
		} else {
			aminoMsg = "(AminoGo Said Nothing)"
		}
	}

	msg += fmt.Sprintf("\nAmino Said: %v", aminoMsg)
	return body, errors.New(msg)
}

// Return a error message for some common HTTP error with some tips
func ThrowHttpErrorIfFail(res *http.Response) (interface{}, error) {

	if res.StatusCode == 0 {
		return errors.New("empty or 0 HTTP status is not allowed"), nil
	}

	switch res.StatusCode {
	case 404:
		body, err := formReturnMessage(fmt.Sprintf("fail to call %s due to resource not found, resulted in a 404 status code", res.Request.URL), res)
		if err != nil {
			return nil, err
		}
		return body, nil
	case 400:
		body, err := formReturnMessage(fmt.Sprintf("fail to call %s due to bad request (perhaps you are giving the wrong arguments), resulted in a 400 status code", res.Request.URL), res)
		if err != nil {
			return nil, err
		}
		return body, nil
	case 401:
		body, err := formReturnMessage(fmt.Sprintf("fail to call %s due to unauthorized (perhaps you could try re-login thus generating a new session-token), resulted in a 401 status code", res.Request.URL), res)
		if err != nil {
			return nil, err
		}
		return body, nil
	case 405:
		body, err := formReturnMessage(fmt.Sprintf("fail to call %s due to bad method not allowed (there must be a bug in AminoGo, please ensure AminoGo are update-to-date, if you are on the latest version, please repo this on GitHub at https://github.com/AminoJS/AminoGo/issues), resulted in a 405 status code", res.Request.URL), res)
		if err != nil {
			return nil, err
		}
		return body, nil
	case 500:
		body, err := formReturnMessage(fmt.Sprintf("fail to call %s due to internal server error(not your fault), resulted in a 500 status code", res.Request.URL), res)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	if res.StatusCode != 200 && res.StatusCode != 201 {
		body, err := formReturnMessage(fmt.Sprintf("fail to call %s resulted in a %d status code", res.Request.URL, res.StatusCode), res)
		if err != nil {
			return nil, err
		}
		return body, nil
	}
	body, _ := formReturnMessage("", res)
	return body, nil
}
