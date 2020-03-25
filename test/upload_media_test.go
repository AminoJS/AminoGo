package test

import (
	"errors"
	"github.com/AminoJS/AminoGo/aminogo"
	"os"
	"testing"
)

func TestUploadMediaBeforeLogin(t *testing.T) {
	_, err := aminogo.UploadMedia("")
	if err == nil {
		t.Error("There should be an error since we have obtain a session token yet")
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

	err = aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		t.Error(err)
		removeMockFiles(mockFileDes)
		return
	}

	expectedError := errors.New("0 byte or completely empty file are not allowed to be transfer to the API's server")
	_, err = aminogo.UploadMedia(mockFileDes)
	if err == nil {
		removeMockFiles(mockFileDes)
		t.Error("There should be an error since this test case is uploading a zero byte file, this action is be not allowed in this library")
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

	err = aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = aminogo.UploadMedia(mockFileDes)
	if err == nil {
		t.Error("There should be an error since this test case is uploading a 6MB+ file, this action is be not allowed in this library")
	}
	expectedError := errors.New("file too large, Amino doesn't allow file size that are larger then 6MB")
	if err.Error() != expectedError.Error() {
		removeMockFiles(mockFileDes)
		t.Errorf("Error message is difference from intended, \nGot:\n%v\nExpect:\n%v\n", err, expectedError)
	}
	removeMockFiles(mockFileDes)

}

func TestUploadLocalMissingLocalResource(t *testing.T) {
	picture := "./missing.jpg"

	err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = aminogo.UploadMedia(picture)
	if err == nil {
		t.Error("There should be an error since this test case is uploading a missing none existing local file")
	}
}

func TestUploadLocalMissingRemoteResource(t *testing.T) {
	picture := "http://pm1.narvii.com/7502/17fe54011759e3ced794abb6e569028620faa81ar1-400-400v2_00.oof"

	err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		t.Error(err)
		return
	}

	expectedError := errors.New("error while trying to capture a remote resources, but ended up with a HTTP status code of: 404")
	_, err = aminogo.UploadMedia(picture)
	if err == nil {
		t.Error("There should be an error since this test case is uploading a missing none existing remote file")
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Error message is difference from intended, \nGot:\n%v\nExpect:\n%v\n", err, expectedError)
	}
}

func TestUploadRemoteResource(t *testing.T) {
	picture := "http://pm1.narvii.com/7502/17fe54011759e3ced794abb6e569028620faa81ar1-400-400v2_00.jpg"

	err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = aminogo.UploadMedia(picture)
	if err != nil {
		t.Error(err)
	}
}

func TestUploadLocalResource(t *testing.T) {
	picture := "image.jpg"

	err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD"))
	if err != nil {
		t.Error(err)
		return
	}

	_, err = aminogo.UploadMedia(picture)
	if err != nil {
		t.Error(err)
	}
}
