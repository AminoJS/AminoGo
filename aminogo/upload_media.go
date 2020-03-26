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
	"strings"
)

type mediaContainer struct {
	// The URL of the tagged source
	des string
	// Flag for IDing the source file whether it is a local file or a remote one
	isRemoteResource bool

	uploadContent interface{}
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
func UploadMedia(url string) (*mediaContainer, error) {

	SID := stores.Get("SID")
	if SID == nil {
		return &mediaContainer{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	if url == "" {
		return nil, errors.New("argument URL must not be empty")
	}

	mc := mediaContainer{
		des:              url,
		isRemoteResource: false,
	}
	return &mc, nil
}

// Upload a remote resource or a local binary file
func (mc *mediaContainer) Remote() (*structs.UploadedMedia, error) {

	if isValidUrl(mc.des) == false {
		return &structs.UploadedMedia{}, errors.New("invalid URL")
	}

	mc.isRemoteResource = true
	err, doneUploading := uploadRemoteFile(mc)
	if err != nil {
		return &structs.UploadedMedia{}, err
	}
	media, err := streamToServer(mc, doneUploading)
	if err != nil {
		return &structs.UploadedMedia{}, err
	}
	return media, nil
}

type PathInterface struct {
	BaseDirectory string
	FileName      string
}

// Upload a local resource or a local binary file
func (mc *mediaContainer) Local(pathArg *PathInterface) (*structs.UploadedMedia, error) {

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

func streamToServer(mc *mediaContainer, doneUploading chan bool) (media *structs.UploadedMedia, err error) {

	SID := stores.Get("SID")
	if SID == nil || SID == "" {
		return &structs.UploadedMedia{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}

	utils.DebugLog("upload_media.go", fmt.Sprintf("des: %s", mc.des))

	endpoint := routes.UploadMedia()

	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}
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

	return resMap, nil
}

func uploadLocalFile(mc *mediaContainer) error {
	// Handle local content

	utils.DebugLog("upload_media.go", "Grepping LOCAL resource")

	// Check if a selected file are larger then 6MB
	localFile, err := os.Open(mc.des)
	if err != nil {
		return err
	}
	defer localFile.Close()

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

func uploadRemoteFile(mc *mediaContainer) (error, chan bool) {
	doneUploading := make(chan bool)

	utils.DebugLog("upload_media.go", "Grepping REMOTE resource")

	desRes, err := http.Get(mc.des)
	if err != nil {
		return err, doneUploading
	}
	if desRes.StatusCode > 299 {
		// Something fishy going on
		return errors.New(fmt.Sprintf("error while trying to capture a remote resources, but ended up with a HTTP status code of: %d", desRes.StatusCode)), doneUploading
	}

	var MaxFileSize = 6000000

	// Check if a selected file are larger then 6MB
	clHeader := desRes.Header.Get("Content-Length")
	clHeaderInt, err := strconv.Atoi(clHeader)
	if err != nil {
		return err, doneUploading
	}
	if clHeaderInt > MaxFileSize {
		return errors.New("file too large, Amino doesn't allow file size that are larger then 6MB"), doneUploading
	}

	// Check if file are 0 bytes
	if clHeaderInt == 0 {
		return errors.New("0 byte or completely empty file are not allowed to be transfer to the API's server"), doneUploading
	}

	utils.DebugLog("upload_media.go", "Done grepping REMOTE resource")

	mc.uploadContent = desRes.Body

	go func() {
		<-doneUploading
		defer desRes.Body.Close()
		close(doneUploading)
	}()

	return nil, doneUploading
}
