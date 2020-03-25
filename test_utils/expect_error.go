package test_utils

import "testing"

func ExpectError(expect error, got error, t *testing.T) {
	if expect.Error() != got.Error() {
		t.Errorf("Error message is difference from intended, \nGot:\n%v\nExpect:\n%v\n", got, expect)
	}
}
