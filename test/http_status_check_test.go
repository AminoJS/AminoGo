package test

import (
	_ "errors"
	_ "fmt"

	"github.com/AminoJS/AminoGo/utils"
	"net/http"
	"testing"
)

func TestEmptyCode(t *testing.T) {
	res := http.Response{StatusCode: 0}
	err := utils.ThrowHttpErrorIfFail(&res)
	if err == nil {
		t.Errorf("There should be an error since the HTTP status code is 0 but Got:\n%v", err)
	}

}

func TestOKResponse(t *testing.T) {
	res := http.Response{StatusCode: 200}
	err := utils.ThrowHttpErrorIfFail(&res)
	if err != nil {
		t.Errorf("Test should pass, since the mokc code is 200 OK, somehow got error:\n%v\n", err)
	}
}

func TestResourceNotFound(t *testing.T) {
	res := http.Response{StatusCode: 404}
	err := utils.ThrowHttpErrorIfFail(&res)
	if err == nil {
		t.Errorf("Test should fail, since the mokc code is 404 OK, somehow got error:\n%v\n", err)
	}
}

func TestUnauthorized(t *testing.T) {
	res := http.Response{StatusCode: 401}
	err := utils.ThrowHttpErrorIfFail(&res)
	if err == nil {
		t.Errorf("Test should fail, since the mokc code is 401 OK, somehow got error:\n%v\n", err)
	}
}

func TestBadRequest(t *testing.T) {
	res := http.Response{StatusCode: 400}
	err := utils.ThrowHttpErrorIfFail(&res)
	if err == nil {
		t.Errorf("Test should fail, since the mokc code is 400 OK, somehow got error:\n%v\n", err)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	res := http.Response{StatusCode: 405}
	err := utils.ThrowHttpErrorIfFail(&res)
	if err == nil {
		t.Errorf("Test should fail, since the mokc code is 405 OK, somehow got error:\n%v\n", err)
	}
}

func TestInternalServerError(t *testing.T) {
	res := http.Response{StatusCode: 500}
	err := utils.ThrowHttpErrorIfFail(&res)
	if err == nil {
		t.Errorf("Test should fail, since the mokc code is 500 OK, somehow got error:\n%v\n", err)
	}
}

func TestHttpErrorCapturing(t *testing.T) {
	codes := [6]int{301, 302, 403, 406, 502, 504}

	for _, code := range codes {
		res := http.Response{StatusCode: code}
		err := utils.ThrowHttpErrorIfFail(&res)
		if err == nil {
			t.Errorf("Test should fail, since the mokc code is %d OK, somehow got error:\n%v\n", code, err)
		}
	}
}
