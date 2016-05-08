// Copyright (C) 2016 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package versioner

import (
	"io/ioutil"
	"os"
	"testing"
)

func setupTestDataDir() {
	os.RemoveAll("testdata")
	os.Mkdir("testdata", 0755)
	os.Chdir("testdata")
}

func cleanupTestDataDir() {
	os.Chdir("..")
	os.RemoveAll("testdata")
}

func TestDefaultVersioning_SuccessfullRemoval(t *testing.T) {
	setupTestDataDir()
	defer cleanupTestDataDir()

	if err := ioutil.WriteFile("file", []byte("data"), 0644); err != nil {
		t.Fatal(err)
	}

	// action
	err := NewDefault("folderId", "folderPath", map[string]string{}).Archive("file")
	if err != nil {
		t.Error(err)
	}

	fileInfos, err := ioutil.ReadDir(".")
	if err != nil {
		t.Fatalf("could not list dir: %s", err)
	}

	// verify
	if len(fileInfos) != 0 {
		t.Errorf("dir should be empty, but has %d entries", len(fileInfos))
	}
}

func TestDefaultVersioning_MissingFile(t *testing.T) {
	setupTestDataDir()
	defer cleanupTestDataDir()

	// action
	err := NewDefault("folderId", "folderPath", map[string]string{}).Archive("file")

	if err != nil {
		t.Errorf("should be nil, as there was no file, but %v", err)
	}
}
