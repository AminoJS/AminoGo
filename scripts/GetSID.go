package main

import (
	"fmt"
	"github.com/AminoJS/AminoGo/aminogo"
	"github.com/AminoJS/AminoGo/stores"
	"os"
)

func main() {
	if os.Getenv("CI") != "true" {
		return
	}
	username := os.Getenv("AMINO_USERNAME")
	password := os.Getenv("AMINO_PASSWORD")
	if username == "" {
		panic("Environment variable AMINO_USERNAME is missing")
	}
	if password == "" {
		panic("Environment variable AMINO_PASSWORD is missing")
	}
	err := aminogo.Login(username, password)
	if err != nil {
		panic(err)
	}
	SID := stores.Get("SID")
	if SID == nil {
		panic("SID is nil")
	}
	fmt.Fprintf(os.Stdout, fmt.Sprintf("%v", SID))
	return
}
