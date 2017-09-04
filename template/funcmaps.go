package template

import (
	"strings"
	"text/template"
)

// Language is a parsable language
// stringer -type=Language
type Language int

// languages
const (
	Typescript Language = iota
	Flow
	Elm
)

// custom types
const (
	EmptyInterface = "emptyIface"
	NestedStruct   = "struct"
	TimeStruct     = "Date"
)

var funcMap = template.FuncMap{
	"updateFlowType":  updateTypes(conversions[Flow]),
	"updateElmType":   updateTypes(conversions[Elm]),
	"updateTSType":    updateTypes(conversions[Typescript]),
	"flowComment":     langComment("//"),
	"elmComment":      langComment("--"),
	"tsComment":       langComment("//"),
	"flowTypeComment": typeComment("//"),
	"elmTypeComment":  typeComment("--"),
	"tsTypeComment":   typeComment("//"),
}

var conversions = map[Language][]string{
	Flow: []string{
		EmptyInterface, "any",
		NestedStruct, "Object",
		TimeStruct, "Date",
		"int64", "number",
		"int32", "number",
		"int16", "number",
		"int8", "number",
		"int", "number",
		"uint64", "number",
		"uint32", "number",
		"uint16", "number",
		"uint8", "number",
		"uint", "number",
		"byte", "number",
		"rune", "number",
		"float32", "number",
		"float64", "number",
		"complex64", "number",
		"complex128", "number",
		"bool", "boolean"},
	Typescript: []string{
		EmptyInterface, "any",
		NestedStruct, "Object",
		TimeStruct, "Date",
		"int64", "number",
		"int32", "number",
		"int16", "number",
		"int8", "number",
		"int", "number",
		"uint64", "number",
		"uint32", "number",
		"uint16", "number",
		"uint8", "number",
		"uint", "number",
		"byte", "number",
		"rune", "number",
		"float32", "number",
		"float64", "number",
		"complex64", "number",
		"complex128", "number",
		"bool", "boolean"},
	Elm: []string{
		"string", "String",
		EmptyInterface, "Maybe",
		NestedStruct, "Maybe",
		TimeStruct, "Date",
		"int64", "Int",
		"int32", "Int",
		"int16", "Int",
		"int8", "Int",
		"int", "Int",
		"uint64", "Int",
		"uint32", "Int",
		"uint16", "Int",
		"uint8", "Int",
		"uint", "Int",
		"byte", "Int",
		"rune", "Int",
		"float32", "Float",
		"float64", "Float",
		"complex64", "Int",
		"complex128", "Int",
		"bool", "Bool"},
}

// updateTypes takes a conversion slice and returns
// a function used as a string replacer
func updateTypes(t []string) func(string) string {
	return func(s string) string { return strings.NewReplacer(t...).Replace(s) }
}

func typeComment(prefix string) func(string) string {
	return func(c string) string {
		if c == "" {
			return c
		}
		out := ""
		sp := strings.Split(c, "\n")
		for i, v := range sp {
			if i != len(sp)-1 {
				out += prefix + " " + v + "\n"
			}
		}
		return out
	}
}

func langComment(prefix string) func(string) string {
	return func(c string) string {
		if c == "" {
			return c
		}
		return prefix + " " + c
	}
}
