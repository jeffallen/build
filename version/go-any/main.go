// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go-any command downloads, installs, and runs a specific
// version of Go, based on its name.
//
// To get and use Go 1.9.1, for example:
//
//     $ go get golang.org/x/build/version/go-any
//     $ ln -s `go env GOPATH`/bin/go-any `go env GOPATH`/bin/go1.9.1
//     $ go1.9.1 download
//
// And then use the go1.9.1 command as if it were your normal go
// command. Make additional symlinks to the go-any command to get
// other versions of Go.
package main

import (
	"fmt"
	"os"
	"path"

	"golang.org/x/build/version"
)

func main() {
	exe := path.Base(os.Args[0])
	if exe == "go-any" {
		fmt.Fprintln(os.Stderr, "You need to make a symlink to this binary, for example: go-any -> go1.9.1.")
		return
	}
	version.Run(exe)
}
