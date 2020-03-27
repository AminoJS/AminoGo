package test

import (
	"fmt"
	"github.com/AminoJS/AminoGo/aminogo"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestEmptyPostTitle(t *testing.T) {
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

	_, err = aminogo.PostBlog(myCom.CommunityList[0].NdcID, "", "Test Post From post_blog_test.go", nil)
	if err == nil {
		t.Error("There should be a error due to a empty post title")
	}
}

func TestEmptyPostContent(t *testing.T) {
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

	_, err = aminogo.PostBlog(myCom.CommunityList[0].NdcID, "Test Case From AminoGo", "", nil)
	if err == nil {
		t.Error("There should be a error due to a empty post title")
	}
}

func TestPostBlogWithCaptionAndInsertedImage(t *testing.T) {

	useCustomContent := true
	gitUsernameCmd := exec.Command("git", "config", "user.name")

	title := "Test Case From AminoGo"

	output, err := gitUsernameCmd.Output()
	if err != nil {
		log.Print(err)
		useCustomContent = false
	}

	if useCustomContent {
		gitUsername := fmt.Sprintf("%v", string(output))
		firstName := strings.Split(gitUsername, " ")
		title = fmt.Sprintf("Test blog from ur boi %s [Test Case Runner]", firstName[0])
	}

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

	picture := "image.jpg"
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	mc, err := aminogo.UploadMedia(picture)
	if err != nil {
		t.Error(err)
	}

	_, err = mc.Local(&aminogo.PathInterface{
		BaseDirectory: os.Getenv("PWD"),
		FileName:      picture,
	})
	if err != nil {
		t.Error(err)
	}

	mc.Captions = "Nick Nazzaro Space Blue"
	refKey := mc.GenerateReferenceKey()

	var mediaList = []*aminogo.MediaContainer{mc}

	blog, err := aminogo.PostBlog(myCom.CommunityList[0].NdcID, title, fmt.Sprintf("This is a picture\n%s\nOf A Wonder Wallpaoer From PoPOS", refKey), &mediaList)
	if err != nil {
		t.Error(err)
	}
	t.Logf("New Blog ID: %v", blog.Blog.BlogID)
}
