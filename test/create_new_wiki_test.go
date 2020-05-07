package test

import (
	"fmt"
	"github.com/AminoJS/AminoGo/aminogo"
	"os"
	"testing"
)

func TestCreateNewWiki(t *testing.T) {
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

	picture := "../image.png"
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	mc, err := aminogo.UploadMedia(picture)
	if err != nil {
		t.Error(err)
	}

	media, err := mc.Local(&aminogo.PathInterface{
		BaseDirectory: os.Getenv("PWD"),
		FileName:      fmt.Sprintf("./test/%s", picture),
	})
	if err != nil {
		t.Error(err)
	}

	wiki, err := aminogo.CreateNewWiki(myCom.CommunityList[0].NdcID, "This is a test wiki OwO", "Hope this works UwU", media.MediaValue)
	if err != nil {
		t.Error(err)
	}
	t.Log(wiki.Item.ItemID)

}
