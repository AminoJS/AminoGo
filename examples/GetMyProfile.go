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
	fmt.Println(myProfile.Account.UID)
}
