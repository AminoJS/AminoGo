package routes

import "fmt"

var ENDPOINT = "https://service.narvii.com"

func Login() string {
	return fmt.Sprintf("%s/api/v1/g/s/auth/login", ENDPOINT)
}

func MyProfile() string {
	return fmt.Sprintf("%s/api/v1/g/s/account", ENDPOINT)
}

func JoinedCommunities(start int, size int) string {
	return fmt.Sprintf("%s/api/v1/g/s/community/joined?start=%d&size=%d", ENDPOINT, start, size)
}

func UploadMedia() string {
	return fmt.Sprintf("%s/api/v1/g/s/media/upload", ENDPOINT)
}

func GetCommunityBlogs(communityID int, start int, size int) string {
	return fmt.Sprintf("%s/api/v1/x%d/s/feed/blog-all?pagingType=t&start=%d&size=%d", ENDPOINT, communityID, start, size)
}

func GetJoinedChatrooms(communityID int, start int, size int) string {
	return fmt.Sprintf("%s/api/v1/x%d/s/chat/thread?type=joined-me&start=%d&size=%d", ENDPOINT, communityID, start, size)
}

func GetUserBlogsFromCommunity(communityID int, UUID string, start int, size int) string {
	return fmt.Sprintf("%s/api/v1/x%d/s/blog?type=user&q=%s&start=%d&size=%d", ENDPOINT, communityID, UUID, start, size)
}

func GetChat(communityID int, threadID string) string {
	return fmt.Sprintf("%s/api/v1/x%d/s/chat/thread/%s/message", ENDPOINT, communityID, threadID)
}

func PostBlog(communityID int) string {
	return fmt.Sprintf("%s/api/v1/v1/x%d/s/blog/", ENDPOINT, communityID)
}
