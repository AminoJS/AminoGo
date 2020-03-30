package structs

import "time"

// All below structures are generated using https://mholt.github.io/json-to-go/

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

type UploadedMedia struct {
	APIStatuscode int       `json:"api:statuscode"`
	APIDuration   string    `json:"api:duration"`
	APIMessage    string    `json:"api:message"`
	MediaValue    string    `json:"mediaValue"`
	APITimestamp  time.Time `json:"api:timestamp"`
}

type CommunityBlogsFeed struct {
	Paging struct {
		NextPageToken string `json:"nextPageToken"`
	} `json:"paging"`
	APIMessage    string `json:"api:message"`
	APIStatuscode int    `json:"api:statuscode"`
	BlogList      []struct {
		GlobalVotesCount   int         `json:"globalVotesCount"`
		GlobalVotedValue   int         `json:"globalVotedValue"`
		VotedValue         int         `json:"votedValue"`
		Keywords           string      `json:"keywords"`
		StrategyInfo       string      `json:"strategyInfo"`
		MediaList          interface{} `json:"mediaList"`
		Style              int         `json:"style"`
		TotalQuizPlayCount int         `json:"totalQuizPlayCount"`
		Title              string      `json:"title"`
		TipInfo            struct {
			TipOptionList []struct {
				Value int    `json:"value"`
				Icon  string `json:"icon"`
			} `json:"tipOptionList"`
			TipMaxCoin      int  `json:"tipMaxCoin"`
			TippersCount    int  `json:"tippersCount"`
			Tippable        bool `json:"tippable"`
			TipMinCoin      int  `json:"tipMinCoin"`
			TipCustomOption struct {
				Value interface{} `json:"value"`
				Icon  string      `json:"icon"`
			} `json:"tipCustomOption"`
			TippedCoins int `json:"tippedCoins"`
		} `json:"tipInfo"`
		ContentRating         int         `json:"contentRating"`
		Content               string      `json:"content"`
		NeedHidden            bool        `json:"needHidden"`
		GuestVotesCount       int         `json:"guestVotesCount"`
		Type                  int         `json:"type"`
		Status                int         `json:"status"`
		GlobalCommentsCount   int         `json:"globalCommentsCount"`
		ModifiedTime          time.Time   `json:"modifiedTime"`
		WidgetDisplayInterval interface{} `json:"widgetDisplayInterval"`
		TotalPollVoteCount    int         `json:"totalPollVoteCount"`
		BlogID                string      `json:"blogId"`
		ViewCount             int         `json:"viewCount"`
		Author                struct {
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
		} `json:"author"`
		Extensions struct {
			Style struct {
				BackgroundMediaList [][]interface{} `json:"backgroundMediaList"`
			} `json:"style"`
			FansOnly bool `json:"fansOnly"`
		} `json:"extensions"`
		VotesCount    int         `json:"votesCount"`
		NdcID         int         `json:"ndcId"`
		CreatedTime   time.Time   `json:"createdTime"`
		EndTime       interface{} `json:"endTime"`
		CommentsCount int         `json:"commentsCount"`
	} `json:"blogList"`
	APIDuration  string    `json:"api:duration"`
	APITimestamp time.Time `json:"api:timestamp"`
}

type JoinedChatrooms struct {
	ThreadList []struct {
		UserAddedTopicList []interface{} `json:"userAddedTopicList"`
		UID                string        `json:"uid"`
		MembersQuota       int           `json:"membersQuota"`
		MembersSummary     []struct {
			Status           int    `json:"status"`
			UID              string `json:"uid"`
			MembershipStatus int    `json:"membershipStatus"`
			Role             int    `json:"role"`
			Nickname         string `json:"nickname"`
			Icon             string `json:"icon"`
		} `json:"membersSummary"`
		ThreadID           string      `json:"threadId"`
		Keywords           interface{} `json:"keywords"`
		MembersCount       int         `json:"membersCount"`
		StrategyInfo       string      `json:"strategyInfo"`
		IsPinned           bool        `json:"isPinned"`
		Title              interface{} `json:"title"`
		MembershipStatus   int         `json:"membershipStatus"`
		Content            interface{} `json:"content"`
		NeedHidden         bool        `json:"needHidden"`
		AlertOption        int         `json:"alertOption"`
		LastReadTime       time.Time   `json:"lastReadTime"`
		Type               int         `json:"type"`
		Status             int         `json:"status"`
		PublishToGlobal    int         `json:"publishToGlobal"`
		ModifiedTime       interface{} `json:"modifiedTime"`
		LastMessageSummary struct {
			UID         string      `json:"uid"`
			MediaType   int         `json:"mediaType"`
			Content     string      `json:"content"`
			MessageID   string      `json:"messageId"`
			CreatedTime time.Time   `json:"createdTime"`
			Type        int         `json:"type"`
			MediaValue  interface{} `json:"mediaValue"`
		} `json:"lastMessageSummary"`
		Condition          int         `json:"condition"`
		Icon               interface{} `json:"icon"`
		LatestActivityTime time.Time   `json:"latestActivityTime"`
		Author             struct {
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
		} `json:"author"`
		Extensions struct {
			LastMembersSummaryUpdateTime int `json:"lastMembersSummaryUpdateTime"`
		} `json:"extensions"`
		NdcID       int         `json:"ndcId"`
		CreatedTime interface{} `json:"createdTime"`
	} `json:"threadList"`
	APIMessage           string `json:"api:message"`
	APIStatuscode        int    `json:"api:statuscode"`
	APIDuration          string `json:"api:duration"`
	PlaylistInThreadList struct {
	} `json:"playlistInThreadList"`
	APITimestamp time.Time `json:"api:timestamp"`
}

type BlogsFromCommunity struct {
	APIStatuscode int    `json:"api:statuscode"`
	APIDuration   string `json:"api:duration"`
	APIMessage    string `json:"api:message"`
	BlogList      []struct {
		GlobalVotesCount   int         `json:"globalVotesCount"`
		GlobalVotedValue   int         `json:"globalVotedValue"`
		VotedValue         int         `json:"votedValue"`
		Keywords           interface{} `json:"keywords"`
		StrategyInfo       string      `json:"strategyInfo"`
		MediaList          interface{} `json:"mediaList"`
		Style              int         `json:"style"`
		TotalQuizPlayCount int         `json:"totalQuizPlayCount"`
		Title              string      `json:"title"`
		TipInfo            struct {
			TipOptionList []struct {
				Value int    `json:"value"`
				Icon  string `json:"icon"`
			} `json:"tipOptionList"`
			TipMaxCoin      int  `json:"tipMaxCoin"`
			TippersCount    int  `json:"tippersCount"`
			Tippable        bool `json:"tippable"`
			TipMinCoin      int  `json:"tipMinCoin"`
			TipCustomOption struct {
				Value interface{} `json:"value"`
				Icon  string      `json:"icon"`
			} `json:"tipCustomOption"`
			TippedCoins int `json:"tippedCoins"`
		} `json:"tipInfo"`
		ContentRating         int         `json:"contentRating"`
		Content               string      `json:"content"`
		NeedHidden            bool        `json:"needHidden"`
		GuestVotesCount       int         `json:"guestVotesCount"`
		Type                  int         `json:"type"`
		Status                int         `json:"status"`
		GlobalCommentsCount   int         `json:"globalCommentsCount"`
		ModifiedTime          time.Time   `json:"modifiedTime"`
		WidgetDisplayInterval interface{} `json:"widgetDisplayInterval"`
		TotalPollVoteCount    int         `json:"totalPollVoteCount"`
		BlogID                string      `json:"blogId"`
		ViewCount             int         `json:"viewCount"`
		Author                struct {
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
		} `json:"author"`
		Extensions    interface{} `json:"extensions"`
		VotesCount    int         `json:"votesCount"`
		NdcID         int         `json:"ndcId"`
		CreatedTime   time.Time   `json:"createdTime"`
		EndTime       interface{} `json:"endTime"`
		CommentsCount int         `json:"commentsCount"`
	} `json:"blogList"`
	APITimestamp time.Time `json:"api:timestamp"`
}

type ChatRecords struct {
	APIStatuscode int    `json:"api:statuscode"`
	APIDuration   string `json:"api:duration"`
	APIMessage    string `json:"api:message"`
	MessageList   []struct {
		IncludedInSummary bool   `json:"includedInSummary"`
		UID               string `json:"uid"`
		Author            struct {
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
		} `json:"author"`
		IsHidden    bool        `json:"isHidden"`
		MessageID   string      `json:"messageId"`
		MediaType   int         `json:"mediaType"`
		Content     interface{} `json:"content"`
		ClientRefID int         `json:"clientRefId"`
		ThreadID    string      `json:"threadId"`
		CreatedTime time.Time   `json:"createdTime"`
		Extensions  struct {
		} `json:"extensions"`
		Type       int    `json:"type"`
		MediaValue string `json:"mediaValue"`
	} `json:"messageList"`
	APITimestamp time.Time `json:"api:timestamp"`
}

type PostedBlog struct {
	Blog struct {
		GlobalVotesCount   int         `json:"globalVotesCount"`
		GlobalVotedValue   int         `json:"globalVotedValue"`
		VotedValue         int         `json:"votedValue"`
		Keywords           string      `json:"keywords"`
		MediaList          interface{} `json:"mediaList"`
		Style              int         `json:"style"`
		TotalQuizPlayCount int         `json:"totalQuizPlayCount"`
		Title              string      `json:"title"`
		TipInfo            struct {
			TipOptionList []struct {
				Value int    `json:"value"`
				Icon  string `json:"icon"`
			} `json:"tipOptionList"`
			TipMaxCoin      int  `json:"tipMaxCoin"`
			TippersCount    int  `json:"tippersCount"`
			Tippable        bool `json:"tippable"`
			TipMinCoin      int  `json:"tipMinCoin"`
			TipCustomOption struct {
				Value interface{} `json:"value"`
				Icon  string      `json:"icon"`
			} `json:"tipCustomOption"`
			TippedCoins int `json:"tippedCoins"`
		} `json:"tipInfo"`
		ContentRating         int         `json:"contentRating"`
		Content               string      `json:"content"`
		NeedHidden            bool        `json:"needHidden"`
		GuestVotesCount       int         `json:"guestVotesCount"`
		Type                  int         `json:"type"`
		Status                int         `json:"status"`
		GlobalCommentsCount   int         `json:"globalCommentsCount"`
		ModifiedTime          time.Time   `json:"modifiedTime"`
		WidgetDisplayInterval interface{} `json:"widgetDisplayInterval"`
		TotalPollVoteCount    int         `json:"totalPollVoteCount"`
		BlogID                string      `json:"blogId"`
		ViewCount             int         `json:"viewCount"`
		Language              interface{} `json:"language"`
		Author                struct {
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
		} `json:"author"`
		Extensions    interface{} `json:"extensions"`
		VotesCount    int         `json:"votesCount"`
		NdcID         int         `json:"ndcId"`
		CreatedTime   time.Time   `json:"createdTime"`
		EndTime       interface{} `json:"endTime"`
		CommentsCount int         `json:"commentsCount"`
	} `json:"blog"`
	APIStatuscode int       `json:"api:statuscode"`
	APIDuration   string    `json:"api:duration"`
	APIMessage    string    `json:"api:message"`
	APITimestamp  time.Time `json:"api:timestamp"`
}

type SentChat struct {
	APIStatuscode int `json:"api:statuscode"`
	Message       struct {
		IncludedInSummary bool   `json:"includedInSummary"`
		UID               string `json:"uid"`
		Author            struct {
			Status                  int         `json:"status"`
			IsNicknameVerified      bool        `json:"isNicknameVerified"`
			UID                     string      `json:"uid"`
			Level                   int         `json:"level"`
			AccountMembershipStatus int         `json:"accountMembershipStatus"`
			MembershipStatus        interface{} `json:"membershipStatus"`
			Reputation              int         `json:"reputation"`
			Role                    int         `json:"role"`
			Nickname                string      `json:"nickname"`
			Icon                    string      `json:"icon"`
		} `json:"author"`
		IsHidden    bool      `json:"isHidden"`
		MessageID   string    `json:"messageId"`
		MediaType   int       `json:"mediaType"`
		Content     string    `json:"content"`
		ClientRefID int       `json:"clientRefId"`
		ThreadID    string    `json:"threadId"`
		CreatedTime time.Time `json:"createdTime"`
		Extensions  struct {
		} `json:"extensions"`
		Type       int         `json:"type"`
		MediaValue interface{} `json:"mediaValue"`
	} `json:"message"`
	APIMessage   string    `json:"api:message"`
	APIDuration  string    `json:"api:duration"`
	APITimestamp time.Time `json:"api:timestamp"`
}
