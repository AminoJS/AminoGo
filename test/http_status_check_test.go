package test

import (
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/test_utils"
	"github.com/AminoJS/AminoGo/utils"
	"testing"
)

func TestEmptyCode(t *testing.T) {
	code := 0
	err := utils.ThrowHttpErrorIfFail(code)
	if err == nil {
		t.Errorf("There should be a error since the HTTP status code is 0 but Got:\n%v", err)
	}

}

func TestOKResponse(t *testing.T) {
	code := 200
	err := utils.ThrowHttpErrorIfFail(code)
	if err != nil {
		t.Errorf("Test should pass, since the mokc code is 200 OK, somehow got error:\n%v\n", err)
	}
}

func TestResourceNotFound(t *testing.T) {
	code := 404
	expectErr := errors.New("fail to login API call due to resource not found, resulted in a none 404 status code")
	err := utils.ThrowHttpErrorIfFail(code)
	if err == nil {
		t.Errorf("Test should fail, since the mokc code is 404 OK, somehow got error:\n%v\n", err)
	}
	test_utils.ExpectError(expectErr, err, t)
}

func TestUnauthorized(t *testing.T) {
	code := 401
	expectErr := errors.New("fail to login API call due to unauthorized (perhaps you could try re-login thus generating a new session-token), resulted in a none 400 status code")
	err := utils.ThrowHttpErrorIfFail(code)
	if err == nil {
		t.Errorf("Test should fail, since the mokc code is 401 OK, somehow got error:\n%v\n", err)
	}
	test_utils.ExpectError(expectErr, err, t)
}

func TestBadRequest(t *testing.T) {
	code := 400
	expectErr := errors.New("fail to login API call due to bad request (perhaps you are giving the wrong arguments), resulted in a none 400 status code")
	err := utils.ThrowHttpErrorIfFail(code)
	if err == nil {
		t.Errorf("Test should fail, since the mokc code is 400 OK, somehow got error:\n%v\n", err)
	}
	test_utils.ExpectError(expectErr, err, t)
}

func TestMethodNotAllowed(t *testing.T) {
	code := 405
	expectErr := errors.New("fail to login API call due to bad method not allowed (there must be a bug in AminoGo, please ensure AminoGo are update-to-date, if you are on the latest version, please repo this on GitHub at https://github.com/AminoJS/AminoGo/issues), resulted in a none 400 status code")
	err := utils.ThrowHttpErrorIfFail(code)
	if err == nil {
		t.Errorf("Test should fail, since the mokc code is 405 OK, somehow got error:\n%v\n", err)
	}
	test_utils.ExpectError(expectErr, err, t)
}

func TestInternalServerError(t *testing.T) {
	code := 500
	expectErr := errors.New("fail to login API call due to internal server error(not your fault), resulted in a none 500 status code")
	err := utils.ThrowHttpErrorIfFail(code)
	if err == nil {
		t.Errorf("Test should fail, since the mokc code is 500 OK, somehow got error:\n%v\n", err)
	}
	test_utils.ExpectError(expectErr, err, t)
}

func TestHttpErrorCapturing(t *testing.T) {
	codes := [6]int{301, 302, 403, 406, 502, 504}

	for _, code := range codes {
		expectErr := errors.New(fmt.Sprintf("fail to login API call, resulted in a none %d status code", code))
		err := utils.ThrowHttpErrorIfFail(code)
		if err == nil {
			t.Errorf("Test should fail, since the mokc code is %d OK, somehow got error:\n%v\n", code, err)
		}
		test_utils.ExpectError(expectErr, err, t)
	}
}
