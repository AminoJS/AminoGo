package aminogo

import (
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/utils"
	"github.com/imroc/req"
	"net/http"
	"net/url"
	"os"
	"path"
)

type mediaContainer struct {
	// The URL of the tagged source
	DES string
	// Flag for IDing the source file whether it is a local file or a remote one
	IsRemoteResource bool

	// The final public URL
	RemoteURL string
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func getLocalFileContent(filePath string) (file interface{}, err error) {
	file, fsErr := os.Open(filePath)
	if fsErr != nil {
		return nil, fsErr
	}
	return file, nil
}

// Upload a remote resource or a local binary file
func UploadMedia(des string) (media *structs.UploadedMedia, err error) {

	SID := stores.Get("SID")
	if SID == nil {
		return &structs.UploadedMedia{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	utils.DebugLog("upload_media.go", fmt.Sprintf("DES: %s", des))

	desContainer := mediaContainer{
		DES:              des,
		IsRemoteResource: isValidUrl(des),
	}

	endpoint := routes.UploadMedia()

	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}

	var uploadContent interface{}

	// Handle remote content
	if desContainer.IsRemoteResource == true {

		utils.DebugLog("upload_media.go", "Grepping REMOTE resource")

		desRes, err := http.Get(des)
		defer desRes.Body.Close()
		if err != nil {
			return &structs.UploadedMedia{}, err
		}
		uploadContent = desRes.Body
		utils.DebugLog("upload_media.go", "Done grepping REMOTE resource")
	}

	// Handle local content
	if desContainer.IsRemoteResource == false {

		utils.DebugLog("upload_media.go", "Grepping LOCAL resource")

		isAbso := path.IsAbs(desContainer.DES)

		if isAbso == false {
			return &structs.UploadedMedia{}, errors.New("at this moment, AminoGo only supports absolute path for local file")
		}

		file, err := getLocalFileContent(desContainer.DES)
		if err != nil {
			return &structs.UploadedMedia{}, err
		}
		uploadContent = file
		utils.DebugLog("upload_media.go", "Done grepping LOCAL resource")
	}

	res, err := req.Post(endpoint, header, uploadContent)
	if err != nil {
		return &structs.UploadedMedia{}, err
	}

	resMap := &structs.UploadedMedia{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.UploadedMedia{}, err
	}

	return resMap, nil

}
