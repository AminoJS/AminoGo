package test

import (
	"errors"
	"github.com/AminoJS/AminoGo/aminogo"
	"github.com/AminoJS/AminoGo/test_utils"
	"os"
	"testing"
)

func TestRequestBlogsFromCommunityBeforeLoggingIn(t *testing.T) {
	_, err := aminogo.GetUserBlogsFromCommunity(&aminogo.GetUserBlogFromComOptions{})
	if err == nil {
		t.Error("There should be an error since we haven't obtain a session token yet")
	}
	test_utils.ExpectError(errors.New("missing SID in state, try using aminogo.Login() first"), err, t)
}

func TestInvalidCommunityId(t *testing.T) {
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	myProfile, err := aminogo.MyProfile()
	if err != nil {
		t.Error(err)
	}

	_, err = aminogo.GetUserBlogsFromCommunity(&aminogo.GetUserBlogFromComOptions{
		CommunityID: 0,
		UUID:        myProfile.Account.UID,
		Start:       0,
		Size:        1,
	})

	expectedError := errors.New("CommunityID cannot be 0 or empty")
	test_utils.ExpectError(expectedError, err, t)

	if err == nil {
		t.Error("There should be an error since argument 'CommunityID' is incorrect")
	}
}

func TestInvalidCommunityUUID(t *testing.T) {
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}
	_, err := aminogo.GetUserBlogsFromCommunity(&aminogo.GetUserBlogFromComOptions{
		CommunityID: 0,
		UUID:        "",
		Start:       0,
		Size:        1,
	})
	if err == nil {
		t.Error("There should be an error since argument 'UUID' is incorrect")
	}
}

func TestGreppingNothing(t *testing.T) {
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	myProfile, err := aminogo.MyProfile()
	if err != nil {
		t.Error(err)
	}

	myCom, err := aminogo.GetJoinedCommunities(&aminogo.GetJoinedCommunitiesOptions{
		Start: 0,
		Size:  0,
	})
	if err != nil {
		t.Error(err)
	}

	blogs, err := aminogo.GetUserBlogsFromCommunity(&aminogo.GetUserBlogFromComOptions{
		CommunityID: myCom.CommunityList[0].NdcID,
		UUID:        myProfile.Account.UID,
		Start:       0,
		Size:        0,
	})
	if err != nil {
		t.Error(err)
	}

	if len(blogs.BlogList) <= 0 {
		t.Errorf("BlogList should not be emtpy since this testing account has join a testing community and posted a blog already")
	}

}
