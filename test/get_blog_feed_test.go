package test

import (
	"github.com/AminoJS/AminoGo/aminogo"
	"os"
	"testing"
)

func TestInvalidCommunityID(t *testing.T) {
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	_, err := aminogo.GetUserBlogsFromCommunity(&aminogo.GetUserBlogFromComOptions{
		CommunityID: 0,
		UUID:        "",
		Start:       0,
		Size:        0,
	})
	if err == nil {
		t.Error("There should be an error since the CommunityID is invalid")
	}
}
func TestGetUserBlogsFromCommunity(t *testing.T) {
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

	blogs, err := aminogo.GetBlogFeed(myCom.CommunityList[0].NdcID, 0, 1)
	if err != nil {
		t.Error(err)
	}
	if len(blogs.BlogList) <= 0 {
		t.Error("Length of blog list shouldn't result in less then or equal 0 since this test Amino has indeed posted a blog at the testing Amino Community")
	}
}
