package stubs

// Maps should all parse right.
// It's hard to get them to do that.
// @strict
type Maps struct {
	MapStringInt  map[string]int            `json:"map_string_to_int"`  // I am a map of strings and ints
	MapStringInts map[string][]int          `json:"map_string_to_ints"` // I am a map of strings to a slice of ints
	MapStringMap  map[string]map[string]int `json:"map_string_to_maps"` // I am a map of strings to maps
	MapIgnore     map[int]int               `json:"-"`
}

type Names []string

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type People map[string]Person

type Nested struct {
	Person struct{} `json:"person"`
}

type Embedded struct {
	Person
}

type Something struct {
	SomeMap map[string][]Embedded `json:"some_map"`
}

// OutgoingSocketMessage sends an action to the client that is easy
// for a redux store to parse
type OutgoingSocketMessage struct {
	Type    string      `json:"type"`    // action type
	Payload interface{} `json:"payload"` // action payload
	Key     string      `json:"key"`     // event key
}
