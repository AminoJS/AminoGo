package test

import (
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/aminogo"
	"os"
	"strings"
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

	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
		removeMockFiles(mockFileDes)
	}

	expectedError := errors.New("0 byte or completely empty file are not allowed to be transfer to the API's server")
	mc, err := aminogo.UploadMedia(mockFileDes)
	if err != nil {
		t.Error(err)
	}
	_, err = mc.Local(&aminogo.PathInterface{
		BaseDirectory: os.Getenv("PWD"),
		FileName:      fmt.Sprintf("./test/%s", mockFileDes),
	})
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

	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
		removeMockFiles(mockFileDes)
	}

	mc, err := aminogo.UploadMedia(mockFileDes)
	if err != nil {
		t.Error(err)
	}

	_, err = mc.Local(&aminogo.PathInterface{
		BaseDirectory: os.Getenv("PWD"),
		FileName:      fmt.Sprintf("./test/%s", mockFileDes),
	})
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
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	mc, err := aminogo.UploadMedia(picture)
	if err != nil {
		t.Error(err)
	}

	_, err = mc.Local(&aminogo.PathInterface{
		BaseDirectory: os.Getenv("PWD"),
		FileName:      fmt.Sprintf("./test/%s", picture),
	})
	if err == nil {
		t.Error("There should be an error since this test case is uploading a missing none existing local file")
	}
}

func TestUploadLocalMissingRemoteResource(t *testing.T) {
	picture := "http://pm1.narvii.com/7502/17fe54011759e3ced794abb6e569028620faa81ar1-400-400v2_00.oof"
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	expectedError := errors.New("error while trying to capture a remote resources, but ended up with a HTTP status code of: 404")
	mc, err := aminogo.UploadMedia(picture)
	if err != nil {
		t.Error(err)
	}

	_, err = mc.Remote()
	if err == nil {
		t.Error("There should be an error since this test case is uploading a missing none existing remote file")
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Error message is difference from intended, \nGot:\n%v\nExpect:\n%v\n", err, expectedError)
	}
}

func TestG304Attack(t *testing.T) {
	picture := "./attack.jpg"
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	mc, err := aminogo.UploadMedia(picture)
	if err != nil {
		t.Error(err)
	}

	_, err = mc.Local(&aminogo.PathInterface{
		BaseDirectory: os.Getenv("PWD"),
		FileName:      fmt.Sprintf("./../../test/%s", picture),
	})
	if err == nil {
		t.Error("There should be an error since this the FileName is out-of-range from the original indented base directory")
	}
	if strings.Contains(err.Error(), "Possible G304 attack!") == false {
		t.Errorf("Error are difference from intented, Got: %v", err)
	}
}

func TestUploadRemoteResource(t *testing.T) {
	picture := "http://pm1.narvii.com/7502/17fe54011759e3ced794abb6e569028620faa81ar1-400-400v2_00.jpg"
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	mc, err := aminogo.UploadMedia(picture)
	if err != nil {
		t.Error(err)
	}

	_, err = mc.Remote()
	if err != nil {
		t.Error(err)
	}

}

func TestUploadLocalResource(t *testing.T) {
	picture := "image.jpg"
	if err := aminogo.Login(os.Getenv("AMINO_USERNAME"), os.Getenv("AMINO_PASSWORD")); err != nil {
		t.Error(err)
	}

	mc, err := aminogo.UploadMedia(picture)
	if err != nil {
		t.Error(err)
	}

	_, err = mc.Local(&aminogo.PathInterface{
		BaseDirectory: os.Getenv("PWD"),
		FileName:      fmt.Sprintf("./test/%s", picture),
	})
	if err != nil {
		t.Error(err)
	}

}
