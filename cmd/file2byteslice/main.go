// Copyright 2018 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// file2byteslice is a dead simple tool to embed a file to Go.
package main

import (
	"flag"
	"io"
	"os"

	"github.com/hajimehoshi/file2byteslice"
)

var (
	inputFilename  = flag.String("input", "", "input filename")
	outputFilename = flag.String("output", "", "output filename")
	packageName    = flag.String("package", "main", "package name")
	varName        = flag.String("var", "_", "variable name")
	compress       = flag.Bool("compress", false, "use gzip compression")
	buildTags      = flag.String("buildtags", "", "build tags")
)

func run() error {
	var out io.Writer
	if *outputFilename != "" {
		f, err := os.Create(*outputFilename)
		if err != nil {
			return err
		}
		defer f.Close()
		out = f
	} else {
		out = os.Stdout
	}

	var in io.Reader
	if *inputFilename != "" {
		f, err := os.Open(*inputFilename)
		if err != nil {
			return err
		}
		defer f.Close()
		in = f
	} else {
		in = os.Stdin
	}

	if err := file2byteslice.Write(out, in, *compress, *buildTags, *packageName, *varName); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		panic(err)
	}
}
