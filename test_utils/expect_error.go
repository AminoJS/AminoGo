package test_utils

import (
	"strings"
	"testing"
)

func ExpectError(expect error, got error, t *testing.T) {
	if strings.Contains(got.Error(), expect.Error()) == false {
		t.Errorf("Error message is difference from intended, \nGot:\n%v\nExpect:\n%v\n", got, expect)
	}
}
