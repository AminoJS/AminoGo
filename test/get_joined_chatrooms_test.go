package test

import (
	"errors"
	"github.com/AminoJS/AminoGo/aminogo"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/test_utils"
	"os"
	"testing"
)

func TestRequestJoinedChatroomsforeLogin(t *testing.T) {
	stores.Remove("SID")
	_, err := aminogo.GetJoinedChatrooms(&aminogo.GetJoinedChatroomsOptions{
		CommunityID: 0,
		Start:       0,
		Size:        0,
	})
	if err == nil {
		t.Error("There should be an error since we haven't obtain a session token yet")
	}
	test_utils.ExpectError(errors.New("missing SID in state, try using aminogo.Login() first"), err, t)
}

func TestInvalidCommunityIdForChatrooms(t *testing.T) {
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}
	_, err := aminogo.GetJoinedChatrooms(&aminogo.GetJoinedChatroomsOptions{
		CommunityID: 0,
		Start:       0,
		Size:        0,
	})
	if err == nil {
		t.Error("There should be an error since we haven't obtain a session token yet")
	}
	expectedErr := errors.New("fail to login API call due to bad request (perhaps you are giving the wrong arguments), resulted in a none 400 status code")
	test_utils.ExpectError(expectedErr, err, t)
}

func TestGetJoinedChatrooms(t *testing.T) {
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	myCom, err := aminogo.GetJoinedCommunities(&aminogo.GetJoinedCommunitiesOptions{
		Start: 0,
		Size:  1,
	})
	if err != nil {
		t.Error(err)
	}

	chatroms, err := aminogo.GetJoinedChatrooms(&aminogo.GetJoinedChatroomsOptions{
		CommunityID: myCom.CommunityList[0].NdcID,
		Start:       0,
		Size:        0,
	})
	if err != nil {
		t.Error(err)
	}

	if len(chatroms.ThreadList) <= 0 {
		t.Error("Length of chatrooms list shouldn't result in less then or equal 0 since this test Amino has indeed joined a chat room")
	}
}
