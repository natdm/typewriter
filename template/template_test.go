package template

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TemplateTestSuite struct {
	suite.Suite
}

func TestTemplateTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateTestSuite))
}

func (s *TemplateTestSuite) TestFlowTemplatePackageTypes() {
	p := &PackageType{
		Name:    "Maps",
		Comment: "... Comment\n",
		Type: &Struct{
			Fields: []Field{
				{
					Name:        "MapStringMap",
					Type:        &Map{Key: &Basic{"string", true}, Value: &Map{Key: &Basic{"string", true}, Value: &Basic{"string", true}}},
					LineComment: "I am a map of strings and ints",
					Tag:         `json:"map_string_map"`,
				}, {
					Name:        "MapStringInts",
					Type:        &Map{Key: &Basic{"string", false}, Value: &Array{Type: &Basic{"number", false}}},
					LineComment: "I am a map of strings to a slice of ints",
					Tag:         `json:"map_string_ints"`,
				},
			},
		},
	}

	buf := new(bytes.Buffer)
	s.Require().NoError(p.Template(buf, Flow))
	expected := `
// ... Comment
export type Maps = {
	map_string_map: { [key: ?string]: { [key: ?string]: ?string } }, // I am a map of strings and ints
	map_string_ints: { [key: string]: number[] }, // I am a map of strings to a slice of ints
}`
	s.EqualValues(expected, buf.String())
}

func (s *TemplateTestSuite) TestFlowTemplateArrayOfMaps() {
	p := &PackageType{
		Name:    "Array",
		Comment: "... Comment\n",
		Type: &Array{
			Type: &Map{
				Key:   &Basic{"int", false},
				Value: &Basic{"string", true},
			},
		},
	}

	buf := new(bytes.Buffer)
	s.Require().NoError(p.Template(buf, Flow))
	expected := `
// ... Comment
export type Array = Array<{ [key: number]: ?string }>`
	s.Equal(expected, buf.String())
}

func (s *TemplateTestSuite) TestFlowTemplateArrayOfCustom() {
	p := &PackageType{
		Name:    "CustomTypeArray",
		Comment: "... Comment\n",
		Type: &Array{
			Type: &Basic{
				Type:    "CustomType",
				Pointer: false,
			},
		},
	}

	buf := new(bytes.Buffer)
	s.Require().NoError(p.Template(buf, Flow))
	expected := `
// ... Comment
export type CustomTypeArray = CustomType[]`
	s.Equal(expected, buf.String())
}

func (s *TemplateTestSuite) TestFlowTemplateArrayOfInts() {
	p := &PackageType{
		Name:    "Array",
		Comment: "... Comment\n",
		Type: &Array{
			Type: &Basic{
				Type:    "int",
				Pointer: false,
			},
		},
	}

	buf := new(bytes.Buffer)
	s.Require().NoError(p.Template(buf, Flow))
	expected := `
// ... Comment
export type Array = number[]`
	s.Equal(expected, buf.String())
}

func (s *TemplateTestSuite) TestFlowTemplateMapOfStringsInts() {
	p := &PackageType{
		Name:    "MapOfStringInts",
		Comment: "... Comment\n",
		Type: &Map{
			Key:   &Basic{"string", false},
			Value: &Basic{"int", false},
		},
	}

	buf := new(bytes.Buffer)
	s.Require().NoError(p.Template(buf, Flow))
	expected := `
// ... Comment
export type MapOfStringInts = { [key: string]: number }`
	s.Equal(expected, buf.String())
}

func (s *TemplateTestSuite) TestFlowInt() {
	p := &PackageType{
		Name:    "AliasToInt",
		Comment: "... Comment\n",
		Type:    &Basic{"int", false},
	}

	buf := new(bytes.Buffer)
	s.Require().NoError(p.Template(buf, Flow))
	expected := `
// ... Comment
export type AliasToInt = number`
	s.Equal(expected, buf.String())
}

func (s *TemplateTestSuite) TestTime() {
	p := &PackageType{
		Name:    "TimeToDate",
		Comment: "... Comment\n",
		Type:    &TimeType{"TestTime", "", ""},
	}

	buf := new(bytes.Buffer)
	s.Require().NoError(p.Template(buf, Flow))
	expected := `
// ... Comment
export type TimeToDate = Date`
	s.Equal(expected, buf.String())
}
