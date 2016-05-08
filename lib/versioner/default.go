// Copyright (C) 2016 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package versioner

import (
	"github.com/syncthing/syncthing/lib/osutil"
	"os"
)

func init() {
	// Register the constructor for this type of versioner with the name "default"
	Factories["default"] = NewDefault
}

type Default struct{}

func NewDefault(folderID, folderPath string, params map[string]string) Versioner {
	s := Default{}
	l.Debugf("instantiated %#v", s)
	return s
}

// Archive deletes the named file away to a version archive. If this function
// returns nil, the named file does not exist any more (has been archived).
func (v Default) Archive(filePath string) error {
	_, err := osutil.Lstat(filePath)
	if os.IsNotExist(err) {
		l.Debugln("not archiving nonexistent file", filePath)
		return nil
	} else if err != nil {
		return err
	}
	return osutil.Remove(filePath)
}
