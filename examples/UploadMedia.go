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

	/*
	   Original image is from one of the wallpapers from PoPOS
	   https://github.com/pop-os/wallpapers/blob/master/original/nick-nazzaro-space-blue.png
	*/
	imageSource := "http://pm1.narvii.com/7511/262dc66e4d7e3256b1ddbd10bf216a17b85abb69r1-2048-1152v2_00.jpg"

	mediaContainer, err := aminogo.UploadMedia(imageSource)
	if err != nil {
		fmt.Println(err)
		return
	}

	media, err := mediaContainer.Remote()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(media.MediaValue)
}
