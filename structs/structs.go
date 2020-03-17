package structs

import "time"

var ENDPOINT string = "https://service.narvii.com/"

// All structs are generated using https://mholt.github.io/json-to-go/

type MyProfile struct {
	APIStatuscode int `json:"api:statuscode"`
	Account       struct {
		Username              interface{}     `json:"username"`
		Status                int             `json:"status"`
		UID                   string          `json:"uid"`
		ModifiedTime          time.Time       `json:"modifiedTime"`
		TwitterID             interface{}     `json:"twitterID"`
		Activation            int             `json:"activation"`
		PhoneNumberActivation int             `json:"phoneNumberActivation"`
		EmailActivation       int             `json:"emailActivation"`
		FacebookID            interface{}     `json:"facebookID"`
		Nickname              string          `json:"nickname"`
		MediaList             [][]interface{} `json:"mediaList"`
		GoogleID              interface{}     `json:"googleID"`
		Icon                  string          `json:"icon"`
		SecurityLevel         int             `json:"securityLevel"`
		PhoneNumber           interface{}     `json:"phoneNumber"`
		Membership            interface{}     `json:"membership"`
		AdvancedSettings      struct {
			AmplitudeAnalyticsEnabled int         `json:"amplitudeAnalyticsEnabled"`
			AmplitudeAppID            interface{} `json:"amplitudeAppId"`
		} `json:"advancedSettings"`
		Role            int       `json:"role"`
		AminoIDEditable bool      `json:"aminoIdEditable"`
		AminoID         string    `json:"aminoId"`
		CreatedTime     time.Time `json:"createdTime"`
		Extensions      struct {
			AdsLevel   int `json:"adsLevel"`
			DeviceInfo struct {
				LastClientType int `json:"lastClientType"`
			} `json:"deviceInfo"`
			PopupConfig struct {
				Ads struct {
					Status int `json:"status"`
				} `json:"ads"`
			} `json:"popupConfig"`
			AdsFlags int `json:"adsFlags"`
		} `json:"extensions"`
		Email string `json:"email"`
	} `json:"account"`
	APIMessage   string    `json:"api:message"`
	APIDuration  string    `json:"api:duration"`
	APITimestamp time.Time `json:"api:timestamp"`
}
