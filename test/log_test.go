package util

import (
	"log"
	"testing"
)

//fakeFile https://golang.org/pkg/io/#Writer
type fakeFile struct {
	payload string
}

func (ff *fakeFile) Write(p []byte) (n int, err error) {
	ff.payload = string(p[:])
	err = nil

	return len(p), err
}

//TestLog Testing log funcionality
func TestLog(t *testing.T) {
	ffile := &fakeFile{}

	logger := log.New(ffile, "", log.LstdFlags)

	logger.Println("TEST MSG")

	log.Println(ffile.payload)

	if ffile.payload != "TEST MSG" {
		t.Error("Something is not right.")
	}
}
