# Typewriter

### Parse Go JSON-tagged types to other language types. Focused on front-end languages.


Currently supports JavaScript Flow. In progress is Typescript and Elm.

Please create an Issue for other requests, and examples of Go types to the requested language.

##### Currently under development. If there's a need for something that isn't on the TODO list, please make an issue.

____
### Example:
Go types:
```go
// Maps should all parse right.
// It's hard to get them to do that.
type Maps struct {
	MapStringInt  map[string]int            `json:"map_string_to_int"`  // I am a map of strings and ints
	MapStringInts map[string][]int          `json:"map_string_to_ints"` // I am a map of strings to a slice of ints
	MapStringMap  map[string]map[string]int `json:"map_string_to_maps"` // I am a map of strings to maps
	MapIgnore     map[int]int               `json:"-"`
}
```

Converts to flow types:
```js
// Maps should all parse right.
// It's hard to get them to do that.
export type Maps = { 
	map_string_to_int: { [key: string]: number }, // I am a map of strings and ints
	map_string_to_ints: { [key: string]: Array<number> }, // I am a map of strings to a slice of ints
	map_string_to_maps: { [key: string]: { [key: string]: number } }// I am a map of strings to maps

}
```
---

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
* ~~Bring all types from embedded types to root types~~ *DONE*
* ~~Ignore in notes~~ *DONE* -- use @ignore
* Type name overrides in tags for multiple languages
* Flow
  * ~~Allow `strict` types~~ *DONE* -- use @strict in notes above struct