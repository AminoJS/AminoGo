package main

import (
	"fmt"
	"github.com/AminoJS/AminoGo/aminogo"
	"os"
)

func main() {
	err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		fmt.Println(err)
		return
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
	fmt.Printf("Chatroom Title: %s\nFrom Community: %s\nID: %s\n", chatroom.Title, yourFirstCommunity.CommunityList[0].Name, chatroom.ThreadID)
}
