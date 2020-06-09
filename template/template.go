package template

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
)

// this file contains all the logic for creting types based on each language utilizing `fragments.go` within the template.

// Templater interface is able to write a template to a writer, based on a Language
type Templater interface {
	Template(w io.Writer, lang Language) error
}

type TypeSpec interface {
	Templater
	IsPointer() bool
}

var errNoType = errors.New("type not stored in package level type declaration")

// Header is the file header
func Header(w io.Writer, lang Language) error {
	tmpl, err := newTemplate(lang, header)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, nil)
}

// Raw is a template with raw input in it
func Raw(w io.Writer, raw string) error {
	tmpl, err := template.New("raw").Parse(raw)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, nil)
}

// TimeType is specifically for Go's time.Time
type TimeType struct {
	Name    string
	Comment string
	Tag     string
}

func (t *TimeType) Template(w io.Writer, lang Language) error {
	tmpl, err := newTemplate(lang, timeType)
	if err != nil {
		return err
	}

	return tmpl.Execute(w, t)
}

// PackageType is a package-level type. Any package type will
// be templated with a full type creation statement and possibly a comment
type PackageType struct {
	Name    string
	Comment string
	Type    Templater
	Tag     string
}

func (t *PackageType) Template(w io.Writer, lang Language) error {
	tmpl, err := newTemplate(lang, declaration)
	if err != nil {
		return err
	}
	if err = tmpl.Execute(w, t); err != nil {
		return err
	}
	if t.Type == nil {
		log.WithError(errNoType).WithField("name", t.Name).Error("error while writing package type")
		return errNoType
	}

	return t.Type.Template(w, lang)
}

// Basic is a basic type. Ints, strings, bools, etc.. or a custom type.
type Basic struct {
	Type    string
	Pointer bool
}

func (t *Basic) Template(w io.Writer, lang Language) error {
	tmpl, err := newTemplate(lang, basic)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, t)
}

func (t *Basic) IsPointer() bool {
	return t.Pointer
}

type Map struct {
	Key   Templater
	Value Templater
}

func (t *Map) Template(w io.Writer, lang Language) error {
	tmpl, err := newTemplate(lang, mapKey)
	if err != nil {
		return err
	}
	if err = tmpl.Execute(w, t); err != nil {
		return err
	}
	if err = t.Key.Template(w, lang); err != nil {
		return err
	}
	tmpl, err = newTemplate(lang, mapValue)
	if err != nil {
		return err
	}
	if err = tmpl.Execute(w, t); err != nil {
		return err
	}

	if err = t.Value.Template(w, lang); err != nil {
		return err
	}
	tmpl, err = newTemplate(lang, mapClose)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, t)
}

func (t *Map) IsPointer() bool {
	return false
}

// Array has a type
type Array struct {
	Type Templater
}

func (t *Array) Template(w io.Writer, lang Language) error {
	tmpl, err := newTemplate(lang, arrayOpen)
	if err != nil {
		return err
	}
	if err = tmpl.Execute(w, t); err != nil {
		return err
	}
	if err = t.Type.Template(w, lang); err != nil {
		return err
	}
	tmpl, err = newTemplate(lang, arrayClose)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, t)
}

func (t *Array) IsPointer() bool {
	return false // TODO: track pointers to arrays
}

// Struct only has fields
type Struct struct {
	Fields []Field

	// Strict is just for Flow types.
	Strict bool

	// Embedded are the embedded types for a struct
	Embedded []string
}

func (t *Struct) Template(w io.Writer, lang Language) error {
	tmpl, err := newTemplate(lang, structOpen)
	if err != nil {
		return err
	}
	if err = tmpl.Execute(w, t); err != nil {
		return err
	}
	for i, v := range t.Fields {
		if err = v.Template(w, lang); err != nil {
			return err
		}
		if i < len(t.Fields)-1 {
			tmpl, err = newTemplate(lang, fieldClose)
			if err != nil {
				return err
			}
			if err = tmpl.Execute(w, nil); err != nil {
				return err
			}
			tmpl, err = newTemplate(lang, comment)
			if err != nil {
				return err
			}
			if err = tmpl.Execute(w, v); err != nil {
				return err
			}
		} else {
			tmpl, err = newTemplate(lang, comment)
			if err != nil {
				return err
			}
			if err = tmpl.Execute(w, v); err != nil {
				return err
			}
			Raw(w, "\n")
		}
	}
	tmpl, err = newTemplate(lang, structClose)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, t)
}

// Field is a struct field
type Field struct {
	Name    string
	Type    TypeSpec
	Comment string
	Tag     string
}

func (t *Field) Template(w io.Writer, lang Language) error {
	jsonName := strings.Split(getTag("json", t.Tag), ",")[0]
	if jsonName != "" {
		t.Name = jsonName
	}

	// Golang allows any valid JSON property name to be provided in the JSON tag.
	// Some aren't valid JS identifiers, so we want to quote them.
	switch lang {
	case Typescript, Flow:
		if propertyShouldBeQuoted(t.Name) {
			t.Name = fmt.Sprintf(`"%s"`, t.Name)
		}
	default:
	}

	tmpl, err := newTemplate(lang, fieldName)
	basicType, isBasic := t.Type.(*Basic)
	if lang == Typescript && isBasic {
		// Special case for TS: top-level nullable type is written as
		// field?: T
		// but if that's the type parameter, it should become
		// T | undefined
		// So, we drop the Pointer flag for top-level types, since the field
		// already has "?" in it.
		t.Type = &Basic{
			Type:    basicType.Type,
			Pointer: false,
		}
	}

	if err != nil {
		return err
	}
	if err = tmpl.Execute(w, t); err != nil {
		return err
	}

	// If there is an override type on the struct field
	override := strings.Split(getTag("tw", t.Tag), ",")
	switch len(override) {
	case 2:
		ptr, err := strconv.ParseBool(string(override[1]))
		if err != nil {
			log.WithError(err).Errorf("error parsing bool for type %s", t.Name)
		}
		t.Type = &Basic{
			Type:    string(override[0]),
			Pointer: ptr,
		}
	case 1:
		if string(override[0]) != "" {
			t.Type = &Basic{
				Type:    string(override[0]),
				Pointer: false,
			}
		}
	}
	return t.Type.Template(w, lang)
}

func getTag(tag string, tags string) string {
	loc := strings.Index(tags, fmt.Sprintf("%s:\"", tag))
	if loc <= -1 {
		return ""
	}
	bs := []byte(tags)
	bs = bs[loc+len(tag)+2:]
	loc = strings.Index(string(bs), "\"")
	if loc == -1 {
		return ""
	}
	return string(bs[:loc])
}
