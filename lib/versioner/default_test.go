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

func setupTestDataDir(testdir string) {
	os.RemoveAll(testdir)
	os.Mkdir(testdir, 0755)
}

func TestDefaultVersioningOnSuccessfullRemoval(t *testing.T) {
	testdir := "testdata23/"
	setupTestDataDir(testdir)
	defer os.RemoveAll(testdir)

	if err := ioutil.WriteFile(testdir+"file", []byte("data"), 0644); err != nil {
		t.Fatal(err)
	}

	// action
	err := NewDefault("folderId", "folderPath", map[string]string{}).Archive(testdir + "file")
	if err != nil {
		t.Error(err)
	}

	fileInfos, err := ioutil.ReadDir(testdir)
	if err != nil {
		t.Fatalf("could not list dir: %s", err)
	}

	// verify
	if len(fileInfos) != 0 {
		t.Errorf("dir should be empty, but has %d entries", len(fileInfos))
	}
}

func TestDefaultVersioningOnMissingFile(t *testing.T) {
	testdir := "testdata23/"
	setupTestDataDir(testdir)
	defer os.RemoveAll(testdir)

	// action
	err := NewDefault("folderId", "folderPath", map[string]string{}).Archive(testdir + "file")

	if err != nil {
		t.Errorf("should be nil, as there was no file, but %v", err)
	}
}

func TestDefaultVersioningRename(t *testing.T) {
	testdir := "testdata23/"
	setupTestDataDir(testdir)
	defer os.RemoveAll(testdir)

	if err := ioutil.WriteFile(testdir+"from", []byte("data"), 0644); err != nil {
		t.Fatal(err)
	}

	// action
	versioner := NewDefault("folderId", "folderPath", map[string]string{})
	err := versioner.Rename(testdir+"from", testdir+"to")

	if err != nil {
		t.Errorf("rename failed, but %v", err)
	}

	if _, err := os.Lstat(testdir + "to"); err != nil {
		t.Errorf("rename failed, but %v", err)
	}
}
