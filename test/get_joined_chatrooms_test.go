package test

import (
	_ "errors"
	"github.com/AminoJS/AminoGo/aminogo"

	"os"
	"testing"
)

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
