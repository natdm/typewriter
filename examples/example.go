package stubs

import (
	"github.com/jinzhu/gorm"
	pkg "github.com/natdm/typewriter/examples/package"
	"github.com/ponzu-cms/ponzu/system/item"
)

// Data should all parse right.
// It's hard to get them to do that.
// @strict
type Data struct {
	MapStringInt  map[string]int            `json:"map_string_to_int" tw:"override_map_name,false"`  // I am a map of strings and ints
	MapStringInts map[string][]int          `json:"map_string_to_ints"`                              // I am a map of strings to a slice of ints
	MapStringMap  map[string]map[string]int `json:"map_string_to_maps" tw:"override_map_name2,true"` // I am a map of strings to maps
	MapIgnore     map[int]int               `json:"-"`
	Peeps         People                    `json:"peeps"`
	ExternalMap   pkg.DataType              `json:"external_embedded"`
	KebabCase     string                    `json:"kebab-case"`
}

type MyInvalidJsType struct {
	someProperty string `json:"some-property"` // wow, why did we do this? totally valid JS though
}

// Person ...
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// People is a map of strings to person
type People map[string]Person

// Nested defaults to the closest "Object" type in any language. Utilize the `tw` tag if needed.
type Nested struct {
	Person struct{} `json:"person"`
}

// Embedded will take all types from the embedded types and insert them in to the new type.
type Embedded struct {
	Person
}

// EmbeddedGormModelTest represents a model that has an embedded type in it.
type EmbeddedGormModelTest struct {
	gorm.Model
	Name        string `gorm:"column:name;index" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	EditEvent   bool   `gorm:"column:edit_event" json:"edit_event"`
	DelBid      bool   `gorm:"column:del_bid" json:"del_bid"`
	AddBid      bool   `gorm:"column:add_bid" json:"add_bid"`
	Billing     bool   `gorm:"column:billing" json:"billing"`
}

type ExternalEmbedded struct {
	item.Item
	Name string `json:"name"`
}

type Items []item.Item
