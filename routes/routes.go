package routes

import "fmt"

var ENDPOINT = "https://service.narvii.com/"

func Login() string {
	return fmt.Sprintf("%s/api/v1/g/s/auth/login", ENDPOINT)
}

func MyProfile() string {
	return fmt.Sprintf("%s/api/v1/g/s/account", ENDPOINT)
}

func JoinedCommunities(start int, size int) string {
	return fmt.Sprintf("%s/api/v1/g/s/community/joined?start=%d&size=%d", ENDPOINT, start, size)
}
