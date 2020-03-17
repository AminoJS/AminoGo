package main

import (
	"AminoJS/AminoGo/aminogo"
	"fmt"
)

func main() {
	err := aminogo.Login("USERNAME", "PASSWORD")
	if err != nil {
		fmt.Println(err)
		return
	}
	myProfile, err := aminogo.MyProfile()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(myProfile.Account.Nickname)
}
