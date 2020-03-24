package aminogo

import (
	"os"
	"testing"
)

func TestZeroByteFile(t *testing.T) {

	_, err := UploadMedia("")
	if err == nil {
		t.Error("There should be a error since we have obtain a session token yet")
	}

	mockFileDes := "__test__.mock"
	mockFile, err = os.Create(mockFileDes)
	if err != nil {
		t.Error(err)
	}

	err = Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = UploadMedia(mockFileDes)
	if err == nil {
		t.Error("There should be a error since this test case is uploading a zero byte file, this action is be not allowed in this library")
	}

}
