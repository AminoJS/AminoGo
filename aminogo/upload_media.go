package aminogo

import (
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/routes"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/AminoJS/AminoGo/structs"
	"github.com/AminoJS/AminoGo/utils"
	"github.com/imroc/req"
	"math/rand"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// A container that store a upload/pre-upload media data
type MediaContainer struct {
	// The URL of the tagged source
	des string
	// Flag for IDing the source file whether it is a local file or a remote one
	isRemoteResource bool

	uploadContent interface{}

	FinalDes string

	// Amino data

	Captions interface{}

	referenceKey string
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
	/* #nosec G304 */
	// Counter measure #1
	// For more reading, please reference SECURITY.md
	file, fsErr := os.Open(filePath)
	if fsErr != nil {
		return nil, fsErr
	}
	return file, nil
}

// Create new MediaContainer
func UploadMedia(url string) (*MediaContainer, error) {

	SID := stores.Get("SID")
	if SID == nil {
		return &MediaContainer{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	if url == "" {
		return nil, errors.New("argument URL must not be empty")
	}

	mc := MediaContainer{
		des:              url,
		isRemoteResource: false,
	}
	return &mc, nil
}

// Upload a remote resource or a local binary file
//func (mc *MediaContainer) Remote() (*structs.UploadedMedia, error) {
//
//    if isValidUrl(mc.des) == false {
//        return &structs.UploadedMedia{}, errors.New("invalid URL")
//    }
//
//    mc.isRemoteResource = true
//    err, doneUploading := uploadRemoteFile(mc)
//    if err != nil {
//        return &structs.UploadedMedia{}, err
//    }
//    media, err := streamToServer(mc, doneUploading)
//    if err != nil {
//        return &structs.UploadedMedia{}, err
//    }
//    return media, nil
//}

type PathInterface struct {
	BaseDirectory string
	FileName      string
}

// Upload a local resource or a local binary file
func (mc *MediaContainer) Local(pathArg *PathInterface) (*structs.UploadedMedia, error) {

	// Format the a valid path

	fileLocation := path.Join(pathArg.BaseDirectory, pathArg.FileName)

	isAbso := path.IsAbs(fileLocation)

	if isAbso == false {
		// Try to convert a relative path to a absolute path
		desAbs, err := filepath.Abs(fileLocation)
		if err != nil {
			return &structs.UploadedMedia{}, err
		}
		fileLocation = desAbs
	}

	// Security check: ageist Gosec G304 attack

	baseDir := filepath.Dir(fileLocation)

	if strings.Contains(baseDir, pathArg.BaseDirectory) == false {

		errText := fmt.Sprintf(`
Possible G304 attack!

Suspects:

	The base directory path is different from the one you has initial given.
	
	Initial Given Base Director:
		%s
	
	Operate Directory:
		%s

References:
	Official Gosec GitHub page: https://github.com/securego/gosec#available-rules`, pathArg.BaseDirectory, baseDir)

		return &structs.UploadedMedia{}, errors.New(errText)
	}

	utils.DebugLog("upload_media.go", fmt.Sprintf("%s %s", pathArg.BaseDirectory, baseDir))

	if isValidUrl(mc.des) == true {
		return &structs.UploadedMedia{}, errors.New("media des is a valid URL, please consider using the Remote() method instead")
	}

	mc.isRemoteResource = false
	mc.des = fileLocation
	err := uploadLocalFile(mc)
	if err != nil {
		return &structs.UploadedMedia{}, err
	}
	media, err := streamToServer(mc, nil)
	if err != nil {
		return &structs.UploadedMedia{}, err
	}
	return media, nil
}

// Generate a random reference key for PostBlog function to use
func (mc *MediaContainer) GenerateReferenceKey() string {
	refKey := randStringRunes(4)
	mc.referenceKey = refKey
	return refKey
}

func streamToServer(mc *MediaContainer, doneUploading chan bool) (media *structs.UploadedMedia, err error) {

	SID := stores.Get("SID")
	if SID == nil || SID == "" {
		return &structs.UploadedMedia{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	utils.DebugLog("upload_media.go", fmt.Sprintf("des: %s", mc.des))

	endpoint := routes.UploadMedia()

	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}

	req.SetTimeout(5 * time.Minute)
	res, err := req.Post(endpoint, header, mc.uploadContent)
	if err != nil {
		return &structs.UploadedMedia{}, err
	}

	// Telling the remote upload function it's done uploading, it's time to close out that precious HTTP connection
	if mc.isRemoteResource && doneUploading != nil {
		doneUploading <- true
	}

	err = utils.ThrowHttpErrorIfFail(res.Response())
	if err != nil {
		return &structs.UploadedMedia{}, err
	}

	resMap := &structs.UploadedMedia{}
	err = res.ToJSON(&resMap)
	if err != nil {
		return &structs.UploadedMedia{}, err
	}

	mc.FinalDes = resMap.MediaValue

	return resMap, nil
}

func uploadLocalFile(mc *MediaContainer) error {
	// Handle local content

	utils.DebugLog("upload_media.go", "Grepping LOCAL resource")

	// Check if a selected file are larger then 6MB
	localFile, err := os.Open(mc.des)
	if err != nil {
		return err
	}
	err = localFile.Close()
	if err != nil {
		return err
	}

	fileInfo, err := localFile.Stat()
	if err != nil {
		return err
	}

	var MaxFileSize int64 = 6000000

	if fileInfo.Size() > MaxFileSize {
		return errors.New("file too large, Amino doesn't allow file size that are larger then 6MB")
	}

	// Check if file are 0 bytes

	if fileInfo.Size() == 0 {
		return errors.New("0 byte or completely empty file are not allowed to be transfer to the API's server")
	}

	file, err := getLocalFileContent(mc.des)
	if err != nil {
		return err
	}
	mc.uploadContent = file
	utils.DebugLog("upload_media.go", "Done grepping LOCAL resource")
	return nil
}

//func uploadRemoteFile(mc *MediaContainer) (error, chan bool) {
//	doneUploading := make(chan bool)
//
//	utils.DebugLog("upload_media.go", "Grepping REMOTE resource")
//
//	req.SetTimeout(5 * time.Minute)
//	desRes, err := req.Get(mc.des)
//	if err != nil {
//		return err, doneUploading
//	}
//	if desRes.Response().StatusCode > 299 {
//		// Something fishy going on
//		return errors.New(fmt.Sprintf("error while trying to capture a remote resources, but ended up with a HTTP status code of: %d", desRes.Response().StatusCode)), doneUploading
//	}
//
//	var MaxFileSize = 6000000
//
//	// Check if a selected file are larger then 6MB
//	clHeader := desRes.Response().Header.Get("Content-Length")
//	clHeaderInt, err := strconv.Atoi(clHeader)
//	if err != nil {
//		return err, doneUploading
//	}
//	if clHeaderInt > MaxFileSize {
//		return errors.New("file too large, Amino doesn't allow file size that are larger then 6MB"), doneUploading
//	}
//
//	// Check if file are 0 bytes
//	if clHeaderInt == 0 {
//		return errors.New("0 byte or completely empty file are not allowed to be transfer to the API's server"), doneUploading
//	}
//
//	utils.DebugLog("upload_media.go", "Done grepping REMOTE resource")
//
//	mc.uploadContent = desRes.Request().Body
//
//	go func() {
//		<-doneUploading
//		defer desRes.Response().Body.Close()
//		close(doneUploading)
//	}()
//
//	return nil, doneUploading
//}
