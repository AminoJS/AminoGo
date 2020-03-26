package test

import (
	"fmt"
	"github.com/AminoJS/AminoGo/aminogo"
	"os"
	"testing"
)

func TestGetChatBeforeLogin(t *testing.T) {
	_, err := aminogo.GetChat(0, "")
	if err == nil {
		t.Error("Should throw error since store have store any SID token yet")
	}
}

func TestGetChatWithInvalidIDs(t *testing.T) {
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	_, err := aminogo.GetChat(0, "")
	if err == nil {
		t.Error("Expect error but got none")
	}
}

func TestGetChat(t *testing.T) {
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	yourFirstCommunity, err := aminogo.GetJoinedCommunities(&aminogo.GetJoinedCommunitiesOptions{
		Start: 0,
		Size:  1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	chatrooms, err := aminogo.GetJoinedChatrooms(&aminogo.GetJoinedChatroomsOptions{
		CommunityID: yourFirstCommunity.CommunityList[0].NdcID,
		Start:       0,
		Size:        1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	chatroom := chatrooms.ThreadList[0]
	chats, err := aminogo.GetChat(yourFirstCommunity.CommunityList[0].NdcID, chatroom.ThreadID)
	if err != nil {
		t.Error(err)
	}

	if len(chats.MessageList) <= 0 {
		t.Error("MessageList is empty, it shouldn't be")
	}

}
