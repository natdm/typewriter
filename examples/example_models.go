package stubs

import "time"

// Date is to assist things
// @ignore
type Date string

// Time ..
type Time Date

// Example represents most of what TW can do.
type Example struct {
	Embedded
	Basic    string           `json:"basic"`         // basic types
	Maps     map[string]Event `json:"maps"`          // map types
	Slices   []Event          `json:"slices_too"`    // slices
	Pointers *Event           `json:"event_pointer"` // pointers
}

// Event ..
type Event struct {
	Name string `json:"name"`
}

// Embedded is testing an embedded struct
type Embedded struct {
	created_at time.Time `json:"created_at" tw:"Date"` // manually overriding field with a name
}
