package test

import (
	"fmt"
	"github.com/AminoJS/AminoGo/aminogo"
	"os"
	"testing"
)

func TestSendChat(t *testing.T) {
	err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		fmt.Println(err)
		return
	}

	com, err := aminogo.GetJoinedCommunities(&aminogo.GetJoinedCommunitiesOptions{
		Start: 0,
		Size:  1,
	})
	if err != nil {
		t.Error(err)
	}

	chatR, err := aminogo.GetJoinedChatrooms(&aminogo.GetJoinedChatroomsOptions{
		CommunityID: com.CommunityList[0].NdcID,
		Start:       0,
		Size:        0,
	})
	if err != nil {
		t.Error(err)
	}

	_, err = aminogo.SendChat(com.CommunityList[0].NdcID, chatR.ThreadList[0].ThreadID, "This is a test message boi")
	if err != nil {
		t.Error(err)
	}

}

func TestSendEmptyTextMessage(t *testing.T) {
	err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		fmt.Println(err)
		return
	}

	com, err := aminogo.GetJoinedCommunities(&aminogo.GetJoinedCommunitiesOptions{
		Start: 0,
		Size:  1,
	})
	if err != nil {
		t.Error(err)
	}

	chatR, err := aminogo.GetJoinedChatrooms(&aminogo.GetJoinedChatroomsOptions{
		CommunityID: com.CommunityList[0].NdcID,
		Start:       0,
		Size:        0,
	})
	if err != nil {
		t.Error(err)
	}

	_, err = aminogo.SendChat(com.CommunityList[0].NdcID, chatR.ThreadList[0].ThreadID, "")
	if err == nil {
		t.Error("Should throw since empty Message")
	}

}
func TestSendingMessageWithEmptyCommunityID(t *testing.T) {
	err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		fmt.Println(err)
		return
	}

	com, err := aminogo.GetJoinedCommunities(&aminogo.GetJoinedCommunitiesOptions{
		Start: 0,
		Size:  1,
	})
	if err != nil {
		t.Error(err)
	}

	chatR, err := aminogo.GetJoinedChatrooms(&aminogo.GetJoinedChatroomsOptions{
		CommunityID: com.CommunityList[0].NdcID,
		Start:       0,
		Size:        0,
	})
	if err != nil {
		t.Error(err)
	}

	_, err = aminogo.SendChat(0, chatR.ThreadList[0].ThreadID, "Test Message")
	if err == nil {
		t.Error("Should throw since empty CommunityID")
	}

}
func TestSendingMessageWithEmptyThreadID(t *testing.T) {
	err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		fmt.Println(err)
		return
	}

	com, err := aminogo.GetJoinedCommunities(&aminogo.GetJoinedCommunitiesOptions{
		Start: 0,
		Size:  1,
	})
	if err != nil {
		t.Error(err)
	}

	_, err = aminogo.GetJoinedChatrooms(&aminogo.GetJoinedChatroomsOptions{
		CommunityID: com.CommunityList[0].NdcID,
		Start:       0,
		Size:        0,
	})
	if err != nil {
		t.Error(err)
	}

	_, err = aminogo.SendChat(com.CommunityList[0].NdcID, "", "Test Message")
	if err == nil {
		t.Error("Should throw since empty ThreadID")
	}

}
