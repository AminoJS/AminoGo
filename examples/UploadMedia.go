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
	imageSource := "image.jpg"

	mediaContainer, err := aminogo.UploadMedia(imageSource)
	if err != nil {
		fmt.Println(err)
		return
	}

	media, err := mediaContainer.Local(&aminogo.PathInterface{
		BaseDirectory: os.Getenv("PWD"),
		FileName:      "./test/image.jpg",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(media.MediaValue)
}
