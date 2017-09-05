package template

import "regexp"

// This file contains utilities for validating ECMAScript tokens prior to emitting them

var ecmaReservedWords = map[string]struct{}{
	"break":        {},
	"case":         {},
	"catch":        {},
	"class":        {},
	"const":        {},
	"continue":     {},
	"debugger":     {},
	"default":      {},
	"delete":       {},
	"do":           {},
	"else":         {},
	"export":       {},
	"extends":      {},
	"finally":      {},
	"for":          {},
	"function":     {},
	"if":           {},
	"import":       {},
	"in":           {},
	"instanceof":   {},
	"new":          {},
	"return":       {},
	"super":        {},
	"switch":       {},
	"this":         {},
	"throw":        {},
	"try":          {},
	"typeof":       {},
	"var":          {},
	"void":         {},
	"while":        {},
	"with":         {},
	"yield":        {},
	"enum":         {},
	"implements":   {},
	"interface":    {},
	"let":          {},
	"package":      {},
	"private":      {},
	"protected":    {},
	"public":       {},
	"static":       {},
	"abstract":     {},
	"boolean":      {},
	"byte":         {},
	"char":         {},
	"double":       {},
	"final":        {},
	"float":        {},
	"goto":         {},
	"int":          {},
	"long":         {},
	"native":       {},
	"short":        {},
	"synchronized": {},
	"throws":       {},
	"transient":    {},
	"volatile":     {},
	"await":        {},
}

// Greedy regex that permits common valid identifiers like `$apply` or `_`
// This quotes more than is strictly necessary, because unicode "letters" in identifier names
// are valid. See: https://stackoverflow.com/questions/2008279/validate-a-javascript-function-name
var validIdentifier = regexp.MustCompile(`^[$\w]+$`)

/**
 * propertyShouldBeQuoted takes a name and returns true if it ought to be quoted as part of a
 * javascript or Typescript/Flow typedef property name. Those conditions are:
 * A - Property is a reserved word or future reserved word in ECMAScript as of Sep 2017
 * B - Property contains a non-word character (note, quotes some valid identifiers)
 */
func propertyShouldBeQuoted(name string) bool {
	_, isReservedWord := ecmaReservedWords[name]
	return isReservedWord || !validIdentifier.MatchString(name)
}
