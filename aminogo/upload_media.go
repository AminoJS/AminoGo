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
	"path/filepath"
	"strconv"
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

	var uploadContent interface{}
	var doneUploading chan bool

	if desContainer.IsRemoteResource == false {
		a, s, err3 := uploadLocalFile(desContainer, uploadContent)
		if err3 != nil {
			return s, err3
		}
		uploadContent = a
	}

	if desContainer.IsRemoteResource == true {
		a, uploadedMedia, err2, d := uploadRemoteFile(des, uploadContent)
		if err2 != nil {
			return uploadedMedia, err2
		}
		uploadContent = a
		doneUploading = d
	}

	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}
	res, err := req.Post(endpoint, header, uploadContent)
	if err != nil {
		return &structs.UploadedMedia{}, err
	}

	// Telling the remote upload function it's done uploading, it's time to close out that precious HTTP connection
	if desContainer.IsRemoteResource {
		doneUploading <- true
	}

	err = utils.ThrowHttpErrorIfFail(res.Response().StatusCode)
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

func uploadLocalFile(desContainer mediaContainer, uploadContent interface{}) (interface{}, *structs.UploadedMedia, error) {
	// Handle local content

	utils.DebugLog("upload_media.go", "Grepping LOCAL resource")

	isAbso := path.IsAbs(desContainer.DES)

	if isAbso == false {
		// Try to convert a relative path to a absolute path
		desAbs, err := filepath.Abs(desContainer.DES)
		if err != nil {
			return nil, &structs.UploadedMedia{}, err
		}
		desContainer.DES = desAbs
	}

	// Check if a selected file are larger then 6MB
	localFile, err := os.Open(desContainer.DES)
	if err != nil {
		return nil, &structs.UploadedMedia{}, err
	}
	defer localFile.Close()

	fileInfo, err := localFile.Stat()
	if err != nil {
		return nil, &structs.UploadedMedia{}, err
	}

	var MaxFileSize int64 = 6000000

	if fileInfo.Size() > MaxFileSize {
		return nil, &structs.UploadedMedia{}, errors.New("file too large, Amino doesn't allow file size that are larger then 6MB")
	}

	// Check if file are 0 bytes

	if fileInfo.Size() == 0 {
		return nil, &structs.UploadedMedia{}, errors.New("0 byte or completely empty file are not allowed to be transfer to the API's server")
	}

	file, err := getLocalFileContent(desContainer.DES)
	if err != nil {
		return nil, &structs.UploadedMedia{}, err
	}
	uploadContent = file
	utils.DebugLog("upload_media.go", "Done grepping LOCAL resource")

	return uploadContent, nil, nil
}

func uploadRemoteFile(des string, uploadContent interface{}) (interface{}, *structs.UploadedMedia, error, chan bool) {
	// Handle remote content

	doneUploading := make(chan bool)

	utils.DebugLog("upload_media.go", "Grepping REMOTE resource")

	desRes, err := http.Get(des)
	if err != nil {
		return nil, &structs.UploadedMedia{}, err, doneUploading
	}
	if desRes.StatusCode > 299 {
		// Something fishy going on
		return nil, &structs.UploadedMedia{}, errors.New(fmt.Sprintf("error while trying to capture a remote resources, but ended up with a HTTP status code of: %d", desRes.StatusCode)), doneUploading
	}
	//defer desRes.Body.Close()

	var MaxFileSize = 6000000

	// Check if a selected file are larger then 6MB
	clHeader := desRes.Header.Get("Content-Length")
	clHeaderInt, err := strconv.Atoi(clHeader)
	if err != nil {
		return nil, &structs.UploadedMedia{}, err, doneUploading
	}
	if clHeaderInt > MaxFileSize {
		return nil, &structs.UploadedMedia{}, errors.New("file too large, Amino doesn't allow file size that are larger then 6MB"), doneUploading
	}

	// Check if file are 0 bytes
	if clHeaderInt == 0 {
		return nil, &structs.UploadedMedia{}, errors.New("0 byte or completely empty file are not allowed to be transfer to the API's server"), doneUploading
	}

	uploadContent = desRes.Body
	utils.DebugLog("upload_media.go", "Done grepping REMOTE resource")

	go func() {
		<-doneUploading
		defer desRes.Body.Close()
		close(doneUploading)
	}()

	return uploadContent, nil, nil, doneUploading
}
