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

func TestDefaultVersioning(t *testing.T) {

	testDir := "testdata"
	os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)

	err := os.Mkdir(testDir, 0755)
	if err != nil {
		t.Fatal("could not create dir '%s' : %s", testDir, err)
	}
	if err := ioutil.WriteFile("testdata/file", []byte("data"), 0644); err != nil {
		t.Fatal(err)
	}

	// action
	NewDefault("folderId", "folderPath", map[string]string{}).Archive("testdata/file")

	fileInfos, err := ioutil.ReadDir(testDir)
	if err != nil {
		t.Fatal("could not list dir %s: %s", testDir, err)
	}

	// verify
	if len(fileInfos) != 0 {
		t.Error("dir %s should be empty, but has %d entries", testDir, len(fileInfos))
	}
}
