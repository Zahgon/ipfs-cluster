package config

import (
	"encoding/json"
	"time"
)

// Saver implements common functionality useful for ComponentConfigs
type Saver struct {
	save    chan struct{}
	BaseDir string
}

// NotifySave signals the SaveCh() channel in a non-blocking fashion.
func (sv *Saver) NotifySave() { _ = "STUB: not implemented"; return }

// Non blocking, in case no one's listening

// SaveCh returns a channel which is signaled when a component wants
// to persist its configuration
func (sv *Saver) SaveCh() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// SetBaseDir is a setter for BaseDir and implements
// part of the ComponentConfig interface.
func (sv *Saver) SetBaseDir(dir string) {
	_ = "STUB: not implemented"

	// DefaultJSONMarshal produces pretty JSON with 2-space indentation
	return
}

func DefaultJSONMarshal(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// SetIfNotDefault sets dest to the value of src if src is not the default
// value of the type.
// dest must be a pointer.
func SetIfNotDefault(src interface{}, dest interface{}) { _ = "STUB: not implemented"; return }

// DurationOpt provides a datatype to use with ParseDurations
type DurationOpt struct {
	// The duration we need to parse
	Duration string
	// Where to store the result
	Dst *time.Duration
	// A variable name associated to it for helpful errors.
	Name string
}

// ParseDurations takes a time.Duration src and saves it to the given dst.
func ParseDurations(component string, args ...*DurationOpt) error {
	_ = "STUB: not implemented"
	return nil
}

// don't do anything. Let the destination field
// stay at its default.

type hiddenField struct{}

func (hf hiddenField) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (hf hiddenField) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"

	// DisplayJSON takes pointer to a JSON-friendly configuration struct and
	// returns the JSON-encoded representation of it filtering out any struct
	// fields marked with the tag `hidden:"true"`, but keeping fields marked
	// with `"json:omitempty"`.
	return nil
}

func DisplayJSON(cfg interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// create a new struct type with same fields
// but setting hidden fields as hidden.

// skip unexported

// remove omitempty from tag, ignore other tags except json

// Parse the original JSON into the new
// struct and re-convert it to JSON.

// Strings is a helper type that (un)marshals a single string to/from a single
// JSON string and a slice of strings to/from a JSON array of strings.
type Strings []string

// UnmarshalJSON conforms to the json.Unmarshaler interface.
func (o *Strings) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON conforms to the json.Marshaler interface.
func (o Strings) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

var _ json.Unmarshaler = (*Strings)(nil)
var _ json.Marshaler = (*Strings)(nil)
