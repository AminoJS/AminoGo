package structs

import "time"

// All struct are generated using https://mholt.github.io/json-to-go/

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

type JoinedCommunities struct {
	CommunityList []struct {
		UserAddedTopicList interface{} `json:"userAddedTopicList"`
		Agent              struct {
			Status                  interface{} `json:"status"`
			IsNicknameVerified      bool        `json:"isNicknameVerified"`
			UID                     string      `json:"uid"`
			Level                   int         `json:"level"`
			FollowingStatus         int         `json:"followingStatus"`
			AccountMembershipStatus int         `json:"accountMembershipStatus"`
			IsGlobal                bool        `json:"isGlobal"`
			MembershipStatus        int         `json:"membershipStatus"`
			Reputation              int         `json:"reputation"`
			Role                    interface{} `json:"role"`
			NdcID                   interface{} `json:"ndcId"`
			MembersCount            int         `json:"membersCount"`
			Nickname                interface{} `json:"nickname"`
			Icon                    interface{} `json:"icon"`
		} `json:"agent"`
		ListedStatus    int `json:"listedStatus"`
		ProbationStatus int `json:"probationStatus"`
		ThemePack       struct {
			ThemeColor        string `json:"themeColor"`
			ThemePackHash     string `json:"themePackHash"`
			ThemePackRevision int    `json:"themePackRevision"`
			ThemePackURL      string `json:"themePackUrl"`
		} `json:"themePack"`
		MembersCount    int     `json:"membersCount"`
		PrimaryLanguage string  `json:"primaryLanguage"`
		CommunityHeat   float64 `json:"communityHeat"`
		StrategyInfo    string  `json:"strategyInfo"`
		Tagline         string  `json:"tagline"`
		JoinType        int     `json:"joinType"`
		Status          int     `json:"status"`
		LaunchPage      struct {
			MediaList [][]interface{} `json:"mediaList"`
			Title     string          `json:"title"`
		} `json:"launchPage"`
		ModifiedTime time.Time `json:"modifiedTime"`
		NdcID        int       `json:"ndcId"`
		ActiveInfo   struct {
		} `json:"activeInfo"`
		Link                 string          `json:"link"`
		Icon                 string          `json:"icon"`
		Endpoint             string          `json:"endpoint"`
		Name                 string          `json:"name"`
		TemplateID           int             `json:"templateId"`
		CreatedTime          time.Time       `json:"createdTime"`
		PromotionalMediaList [][]interface{} `json:"promotionalMediaList"`
	} `json:"communityList"`
	UserInfoInCommunities struct {
		Num227549816 struct {
			UserProfile struct {
				Status                  int    `json:"status"`
				IsNicknameVerified      bool   `json:"isNicknameVerified"`
				UID                     string `json:"uid"`
				Level                   int    `json:"level"`
				FollowingStatus         int    `json:"followingStatus"`
				AccountMembershipStatus int    `json:"accountMembershipStatus"`
				IsGlobal                bool   `json:"isGlobal"`
				MembershipStatus        int    `json:"membershipStatus"`
				Reputation              int    `json:"reputation"`
				Role                    int    `json:"role"`
				NdcID                   int    `json:"ndcId"`
				MembersCount            int    `json:"membersCount"`
				Nickname                string `json:"nickname"`
				Icon                    string `json:"icon"`
			} `json:"userProfile"`
		} `json:"227549816"`
	} `json:"userInfoInCommunities"`
	APIMessage    string    `json:"api:message"`
	APIStatuscode int       `json:"api:statuscode"`
	APIDuration   string    `json:"api:duration"`
	APITimestamp  time.Time `json:"api:timestamp"`
}
