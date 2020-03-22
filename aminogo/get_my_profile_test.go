package aminogo

import (
	"github.com/AminoJS/AminoGo/structs"
	"os"
	"strings"
	"testing"
)

func TestEmptyEmailAddress(t *testing.T) {
	err := Login("", "PWD")
	if err == nil {
		t.Errorf("Fail to optain a session-token, error are following:\n%v", err)
	}
}

func TestEmptyPassword(t *testing.T) {
	err := Login("EMAIL", "")
	if err == nil {
		t.Error("Fail to check a empty password, might lead to nil pointer exception")
	}
}

func TestEmptyAllField(t *testing.T) {
	err := Login("", "")
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

	if os.Getenv("CI") == "true" {
		t.Logf("USERNAME: %s", username[1-2])
		t.Logf("PASSWORD: %s", password[1-2])
	}

	err := Login(username, password)
	if err != nil {
		t.Error(err)
	}

	myProfile, err := MyProfile()
	if err != nil {
		t.Error(err)
	}

	emptyStruct := structs.MyProfile{}

	if myProfile == &emptyStruct {
		t.Error("API call in result of an empty struct, which is bad needless to say")
	}

	UUID := myProfile.Account.UID

	spliited := strings.Split(UUID, "-")

	if len(spliited) != 5 {
		t.Errorf("UUID is malformed,\nExpect format: XXX-XXX-XXX-XXX-XXX (splited with - and a set of 5 sets)\n Got: %v", UUID)
	}
}
