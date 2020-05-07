package test

import (
	"fmt"
	"github.com/AminoJS/AminoGo/aminogo"
	"os"
	"testing"
)

func TestPostComment(t *testing.T) {
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

	comId := com.CommunityList[0].NdcID

	blog, err := aminogo.GetBlogFeed(comId, 0, 1)
	if err != nil {
		t.Error(err)
	}

	msg := "This is a testing comment from a test case"

	comment, err := aminogo.CommentingOnPost(comId, blog.BlogList[0].BlogID, msg)
	if err != nil {
		t.Error(err)
	}
	t.Logf("New comment ID: %s", comment.Comment.CommentID)
}

func TestEmptyCommentMessage(t *testing.T) {
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

	comId := com.CommunityList[0].NdcID

	blog, err := aminogo.GetBlogFeed(comId, 0, 1)
	if err != nil {
		t.Error(err)
	}

	msg := ""

	_, err = aminogo.CommentingOnPost(comId, blog.BlogList[0].BlogID, msg)
	if err == nil {
		t.Error("There should be an error since to comment message is empty")
	}
}
