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
	imageSource := "http://pm1.narvii.com/7502/17fe54011759e3ced794abb6e569028620faa81ar1-400-400v2_00.jpg"

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
