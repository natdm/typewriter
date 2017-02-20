# Typewriter

### Parse Go JSON-tagged types to other language types. Focused on front-end languages.


Currently supports JavaScript Flow, TypeScript, and (some) Elm.

For custom types, add the tag, `tw:"<CustomTypeName>,<PointerBool>"`

Please create an Issue for requests, and include examples of Go types to the requested language.

##### Currently under development. 

[See examples](./EXAMPLES.md)


Does not support:
Nested structs (changes to the closest form, eg 'Object' in flow)
Interfaces within structs

Usage:

```
$ go get github.com/natdm/typewriter
$ $GOPATH/bin/typewriter -dir ./your/models/directory -lang flow -v -out ./save/to/models.js
```

```bash
	Flags:
		-dir	Parse a complete directory 
			example: 	-dir= ../src/appname/models/
			default: 	./

		-file	Parse a single go file 
			example: 	-file= ../src/appname/models/app.go
			overrides 	-dir and -recursive

		-out	Saves content to folder
			example: 	-out= ../src/appname/models/
						-out= ../src/appname/models/customname.js
			default: 	./models. 

		-r		Transcends directories
			example:	-recursive= false
			default:	true

		-v		Verbose logging, detailing every skipped type, file, or field.
			default: 	false

		-lang 	Language to parse to. One of ["flow"]
			example:	-lang flow
			default:	will not parse
```

___
#### TODO:
* More tests