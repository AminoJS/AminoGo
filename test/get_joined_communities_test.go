package test

import (
	"github.com/AminoJS/AminoGo/aminogo"
	"os"
	"testing"
)

func TestGetJoinedCommunities(t *testing.T) {
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

	communities, err := aminogo.GetBlogFeed(myCom.CommunityList[0].NdcID, 0, 1)
	if err != nil {
		t.Error(err)
	}
	if len(communities.BlogList) <= 0 {
		t.Error("Length of communities list shouldn't result in less then or equal 0 since this test Amino has joined one and more communities")
	}
}
