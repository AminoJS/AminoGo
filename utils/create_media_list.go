package utils

import (
	"fmt"
	"strings"
)

type Media struct {
	URL          string
	Caption      interface{}
	ReferenceKey string
}

func CreateMediaList(medias []*Media, context string) (*[]interface{}, *string) {

	var newMediaList []interface{}
	var newContext string = ""

	for _, media := range medias {

		var image [4]interface{}

		image[0] = 100
		image[1] = media.URL

		if media.Caption == "" {
			image[2] = nil
		} else {
			image[2] = media.Caption
		}

		if media.ReferenceKey != "" {

			newContext = strings.ReplaceAll(context, media.ReferenceKey, fmt.Sprintf("[IMG=%s]", media.ReferenceKey))

			image[3] = media.ReferenceKey

		} else {
			image[3] = nil
		}

		newMediaList = append(newMediaList, image)

	}

	return &newMediaList, &newContext

}
