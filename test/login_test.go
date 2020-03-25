package test

import (
	"errors"
	"github.com/AminoJS/AminoGo/aminogo"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/test_utils"
	"os"
	"strings"
	"testing"
)

func TestEmptyEmailAddress(t *testing.T) {
	err := aminogo.Login("", "PWD")

	if err == nil {
		t.Error("This test should fail, since the email argument is empty")
	}

	expectErr := errors.New("email address MUST be provided as a argument of this function call")
	test_utils.ExpectError(expectErr, err, t)

}

func TestEmptyPassword(t *testing.T) {
	err := aminogo.Login("EMAIL", "")

	if err == nil {
		t.Error("This test should fail, since the password argument is empty")
	}

	expectErr := errors.New("password MUST be provided as a argument of this function call")
	test_utils.ExpectError(expectErr, err, t)

}

func TestEmptyAllField(t *testing.T) {
	err := aminogo.Login("", "")
	if err == nil {
		t.Error("Fail to check both empty password and email address, might lead to nil pointer exception")
	}
}

func TestUUID(t *testing.T) {

	username := os.Getenv("AMINO_USERNAME")
	password := os.Getenv("AMINO_PASSWORD")

	// Check if environment variable exits

	if username == "" {
		t.Errorf("Environment variable AMINO_USERNAME is missing")
	}

	if password == "" {
		t.Errorf("Environment variable AMINO_PASSWORD is missing")
	}

	err := aminogo.Login(username, password)
	if err != nil {
		t.Error(err)
	}

	myProfile, err := aminogo.MyProfile()
	if err != nil {
		t.Error(err)
	}

	emptyStruct := structs.MyProfile{}

	if myProfile == &emptyStruct {
		t.Error("API call in result of an empty struct, which is bad needless to say")
	}

	UUID := myProfile.Account.UID

	slitted := strings.Split(UUID, "-")

	if len(slitted) != 5 {
		t.Errorf("UUID is malformed,\nExpect format: XXX-XXX-XXX-XXX-XXX (splited with - and a set of 5 sets)\n Got: %v", UUID)
	}
}
