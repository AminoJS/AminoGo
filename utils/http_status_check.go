package utils

import (
	"errors"
	"fmt"
)

// Return a error message for some common HTTP error with some tips
func ThrowHttpErrorIfFail(status_code int) (err error) {

	if status_code == 0 {
		return errors.New("empty or 0 HTTP status is not allowed")
	}

	switch status_code {
	case 404:
		return errors.New("fail to login API call due to resource not found, resulted in a none 404 status code")
	case 400:
		return errors.New("fail to login API call due to bad request (perhaps you are giving the wrong arguments), resulted in a none 400 status code")
	case 401:
		return errors.New("fail to login API call due to unauthorized (perhaps you could try re-login thus generating a new session-token), resulted in a none 400 status code")
	case 405:
		return errors.New("fail to login API call due to bad method not allowed (there must be a bug in AminoGo, please ensure AminoGo are update-to-date, if you are on the latest version, please repo this on GitHub at https://github.com/AminoJS/AminoGo/issues), resulted in a none 400 status code")
	case 500:
		return errors.New("fail to login API call due to internal server error(not your fault), resulted in a none 500 status code")
	}
	if status_code != 200 && status_code != 201 {
		return errors.New(fmt.Sprintf("fail to login API call, resulted in a none %d status code", status_code))
	}
	return nil
}
