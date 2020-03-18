package routes

import "fmt"

var ENDPOINT = "https://service.narvii.com/"

func GetRoutes() map[string]string {
	routes := map[string]string{
		"Login":     fmt.Sprintf("%s/api/v1/g/s/auth/login", ENDPOINT),
		"MyProfile": fmt.Sprintf("%s/api/v1/g/s/account", ENDPOINT),
	}
	return routes
}
