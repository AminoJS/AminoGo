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
	myProfile, err := aminogo.MyProfile()
	if err != nil {
		fmt.Println(err)
		return
	}

	myCom, err := aminogo.GetJoinedCommunities(&aminogo.GetJoinedCommunitiesOptions{
		Start: 0,
		Size:  1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	myBlogs, err := aminogo.GetUserBlogsFromCommunity(&aminogo.GetUserBlogFromComOptions{
		CommunityID: myCom.CommunityList[0].NdcID,
		UUID:        myProfile.Account.UID,
		Start:       0,
		Size:        1,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(myBlogs.BlogList) <= 0 {
		fmt.Printf("Sorry, there is no blogs you have written in this selected community (%s)\n", myCom.CommunityList[0].Name)
		return
	}

	fmt.Printf("First Blog Title: %s", myBlogs.BlogList[0].Title)
}
