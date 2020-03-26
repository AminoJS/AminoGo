package test

import (
	"github.com/AminoJS/AminoGo/aminogo"
	"os"
	"testing"
)

func TestRequestProfileBeforeLogin(t *testing.T) {
	_, err := aminogo.MyProfile()
	if err == nil {
		t.Error("There should be an error since we have obtain a session token yet")
	}
}

func TestRequestingResource(t *testing.T) {
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	myProfile, err := aminogo.MyProfile()
	if err != nil {
		t.Error(err)
	}

	expectedUsername := "ProjectAmino"
	if myProfile.Account.Nickname != expectedUsername {
		t.Errorf("Expect nickname to be %s, but got %v", expectedUsername, myProfile.Account.Nickname)
	}
}
