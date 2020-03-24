package aminogo

import (
	"errors"
	"os"
	"testing"
)

func TestUploadMediaBeforeLogin(t *testing.T) {
	_, err := UploadMedia("")
	if err == nil {
		t.Error("There should be a error since we have obtain a session token yet")
	}
}

func removeMockFiles(mockFilePath string) {
	os.Remove(mockFilePath)
}

func TestZeroByteFile(t *testing.T) {

	mockFileDes := "__test__.mock"
	_, err := os.Create(mockFileDes)
	if err != nil {
		t.Error(err)
	}

	err = Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		t.Error(err)
		removeMockFiles(mockFileDes)
		return
	}

	expectedError := errors.New("0 byte or completely empty file are not allowed to be transfer to the API's server")
	_, err = UploadMedia(mockFileDes)
	if err == nil {
		removeMockFiles(mockFileDes)
		t.Error("There should be a error since this test case is uploading a zero byte file, this action is be not allowed in this library")
	}
	if err.Error() != expectedError.Error() {
		removeMockFiles(mockFileDes)
		t.Errorf("Error message is difference from intended, \nGot:\n%v\nExpect:\n%v\n", err, expectedError)
	}

	removeMockFiles(mockFileDes)

}

func TestFileTooLarge(t *testing.T) {

	mockFileDes := "__test__large.mock"
	f, err := os.Create(mockFileDes)
	if err != nil {
		t.Error(err)
	}
	if err := f.Truncate(1e7); err != nil {
		t.Error(err)
	}

	err = Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = UploadMedia(mockFileDes)
	if err == nil {
		t.Error("There should be a error since this test case is uploading a 6MB+ file, this action is be not allowed in this library")
	}
	expectedError := errors.New("file too large, Amino doesn't allow file size that are larger then 6MB")
	if err.Error() != expectedError.Error() {
		removeMockFiles(mockFileDes)
		t.Errorf("Error message is difference from intended, \nGot:\n%v\nExpect:\n%v\n", err, expectedError)
	}
	removeMockFiles(mockFileDes)

}
