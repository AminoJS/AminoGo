package main

import (
	"fmt"
	"os"

	"github.com/AminoJS/AminoGo/aminogo"
)

func main() {
	err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		fmt.Println(err)
		return
	}
	joinedCommunities, err := aminogo.GetJoinedCommunities(&aminogo.GetJoinedCommunitiesOptions{
		Start: 0,
		Size:  100,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	firstCommunity := joinedCommunities.CommunityList[0]
	fmt.Printf("First Community ID: %d\nName: %s", firstCommunity.NdcID, firstCommunity.Name)
}
