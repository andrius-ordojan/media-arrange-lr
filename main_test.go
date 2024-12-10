package main

import (
	"os"
	"testing"
)

// TODO:
// how to do cleanup and setup for tests
// how to create temp dir using testing package

func TestShouldSkipWhenMediaExists(t *testing.T) {
	o, err := os.MkdirTemp(".", "pat")
	if err != nil {
		t.Error(err)
	}

	// TODO: how arrange files in a smart way? need to set some constants or structs for each file

	os.RemoveAll(o)
}

// reads subdirectories for media

func TestShouldMoveWhenMediaDoesNotExist(t *testing.T) {
}

func TestShouldMoveToSoocDirWhenProcessingJpgMedia(t *testing.T) {
}

func TestShouldMoveToPhotosDirWhenProcessingJpgOrRafMedia(t *testing.T) {
}

func TestShouldMoveToVideosDirWhenProcessingMovMedia(t *testing.T) {
}

func TestShouldSkipWhenMediaIsCopyButNameIsDifferent(t *testing.T) {
}

func TestShouldMoveWhenMediaExistsButIsNotPlacedCorrectly(t *testing.T) {
}

func TestShouldErrorWhenMetadataNotPresent(t *testing.T) {
}

// TODO: test dynamic hash chunk using mocking
