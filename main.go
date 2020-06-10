// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/natdm/typewriter/parse"
	"github.com/natdm/typewriter/template"
	log "github.com/sirupsen/logrus"
)

func main() {
	inFlag := flag.String("dir", "./", "dir is to specify what folder to parse types from")
	fileFlag := flag.String("file", "", "file is to parse a single file. Will override a directory")
	langFlag := flag.String("lang", "", "determine the language. One of 'flow', 'ts")
	outFlag := flag.String("out", "", "file and path to save output to")
	vFlag := flag.Bool("v", false, "verbose logging")
	recursiveFlag := flag.Bool("r", true, "to recursively ascend all folders in dir")
	expandEmbeddedFlag := flag.Bool("e", false, "expand embedded structs inline")
	flag.Usage = usage
	flag.Parse()

	var lang template.Language
	switch *langFlag {
	case "flow":
		lang = template.Flow
	case "elm":
		lang = template.Elm
		if !*expandEmbeddedFlag {
			log.Fatalln(
				"You have to use -e flag with Elm, which does not support intersection types")
		}
	case "ts":
		lang = template.Typescript
	default:
		log.Fatalln("Please pick a proper language ['elm', 'flow', 'ts']")
	}

	var out io.Writer

	if *outFlag != "" {
		f, err := os.Create(*outFlag)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		out = f
	} else {
		out = os.Stdout
	}

	var (
		files []string
		types map[string]*template.PackageType
		err   error
	)

	if *fileFlag != "" {
		types, err = parse.Files([]string{*fileFlag}, *vFlag, *expandEmbeddedFlag)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		if err = parse.Directory(*inFlag, *recursiveFlag, &files, *vFlag); err != nil {
			log.Fatalln(err)
		}
		types, err = parse.Files(files, *vFlag, *expandEmbeddedFlag)
		if err != nil {
			log.Fatalln(err)
		}
	}
	ct, err := template.Draw(types, out, lang, *vFlag)
	if err != nil {
		log.Fatalln(err)
	}
	log.WithField("output_type_ct", ct).Info("Done")
}

func usage() {
	fmt.Print(`
	Typewriter
	Convert Go types to other languages
	Visit http://www.github.com/natdm/typewriter for more detailed example useage.

	Flags:
		-dir <dir>
			Parse a complete directory
			example: 	-dir= ../src/appname/models/
			default: 	./

		-file <gofile>
			Parse a single go file
			example: 	-file= ../src/appname/models/app.go
			overrides 	-dir and -recursive

		-out <path>
			Saves content to folder
			example: 	-out= ../src/appname/models/
						-out= ../src/appname/models/customname.js
			default: 	./models.

		-lang <lang>
			Language to parse to. One of ["flow", "elm", "ts"]
			example:	-lang flow
			default:	will not parse

		-r
			Transcends directories
			default:	true

		-v
			Verbose logging, detailing every skipped type, file, or field.
			default: 	false
`)
}
