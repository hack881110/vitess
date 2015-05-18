// Copyright 2015, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/youtube/vitess/go/vt/mysqlctl/backupstorage"
	"github.com/youtube/vitess/go/vt/servenv"
)

func init() {
	servenv.OnRun(func() {
		backupstorage.RegisterFileBackupStorage()
	})
}
