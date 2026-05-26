// Package pinsvc contains type definitions for the Pinning Services API
package pinsvc

import (
	"net/url"
	"time"

	types "github.com/ipfs-cluster/ipfs-cluster/api"
)

func init() {
	// initialize trackerStatusString
	stringStatus = make(map[string]Status)
	for k, v := range statusString {
		stringStatus[v] = k
	}
}

// APIError is returned by the API as a body when an error
// occurs. It implements the error interface.
type APIError struct {
	Details APIErrorDetails `json:"error"`
}

// APIErrorDetails contains details about the APIError.
type APIErrorDetails struct {
	Reason  string `json:"reason"`
	Details string `json:"details,omitempty"`
}

func (apiErr APIError) Error() string { _ = "STUB: not implemented"; return "" }

// PinName is a string limited to 255 chars when serializing JSON.
type PinName string

// MarshalJSON converts the string to JSON.
func (pname PinName) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON reads the JSON string and errors if over 256 chars.
func (pname *PinName) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// "a_string" 255 + 2 for quotes
	return nil
}

// Pin contains basic information about a Pin and pinning options.
type Pin struct {
	Cid     types.Cid         `json:"cid"`
	Name    PinName           `json:"name,omitempty"`
	Origins []types.Multiaddr `json:"origins,omitempty"`
	Meta    map[string]string `json:"meta,omitempty"`
}

// Defined returns if the pinis empty (Cid not set).
func (p Pin) Defined() bool { _ = "STUB: not implemented"; return false }

// MatchesName returns in a pin status matches a name option with a given
// match strategy.
func (p Pin) MatchesName(nameOpt string, strategy MatchingStrategy) bool {
	_ = "STUB: not implemented"
	return false
}

// MatchesMeta returns true if the pin status metadata matches the given.  The
// metadata should have all the keys in the given metaOpts and the values
// should, be the same (metadata map includes metaOpts).
func (p Pin) MatchesMeta(metaOpts map[string]string) bool { _ = "STUB: not implemented"; return false }

// Status represents a pin status, which defines the current state of the pin
// in the system.
type Status int

// Values for the Status type.
const (
	StatusUndefined Status = 0
	StatusQueued           = 1 << iota
	StatusPinned
	StatusPinning
	StatusFailed
)

var statusString = map[Status]string{
	StatusUndefined: "undefined",
	StatusQueued:    "queued",
	StatusPinned:    "pinned",
	StatusPinning:   "pinning",
	StatusFailed:    "failed",
}

// values autofilled in init()
var stringStatus map[string]Status

// String converts a Status into a readable string.
// If the given Status is a filter (with several
// bits set), it will return a comma-separated list.
func (st Status) String() string {
	_ = "STUB: not implemented"

	// simple and known composite values
	return ""
}

// other filters

// Match returns true if the tracker status matches the given filter.
func (st Status) Match(filter Status) bool { _ = "STUB: not implemented"; return false }

// MarshalJSON uses the string representation of Status for JSON
// encoding.
func (st Status) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON sets a tracker status from its JSON representation.
func (st *Status) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// StatusFromString parses a string and returns the matching
// Status value. The string can be a comma-separated list
// representing a Status filter. Unknown status names are
// ignored.
func StatusFromString(str string) Status { _ = "STUB: not implemented"; return *new(Status) }

// MatchingStrategy defines a type of match for filtering pin lists.
type MatchingStrategy int

// Values for MatchingStrategy.
const (
	MatchingStrategyUndefined MatchingStrategy = iota
	MatchingStrategyExact
	MatchingStrategyIexact
	MatchingStrategyPartial
	MatchingStrategyIpartial
)

// MatchingStrategyFromString converts a string to its MatchingStrategy value.
func MatchingStrategyFromString(str string) MatchingStrategy {
	_ = "STUB: not implemented"
	return *new(MatchingStrategy)
}

// PinStatus provides information about a Pin stored by the Pinning API.
type PinStatus struct {
	RequestID string            `json:"requestid"`
	Status    Status            `json:"status"`
	Created   time.Time         `json:"created"`
	Pin       Pin               `json:"pin"`
	Delegates []types.Multiaddr `json:"delegates"`
	Info      map[string]string `json:"info,omitempty"`
}

// PinList is the result of a call to List pins
type PinList struct {
	Count   uint64      `json:"count"`
	Results []PinStatus `json:"results"`
}

// ListOptions represents possible options given to the List endpoint.
type ListOptions struct {
	Cids             []types.Cid
	Name             string
	MatchingStrategy MatchingStrategy
	Status           Status
	Before           time.Time
	After            time.Time
	Limit            uint64
	Meta             map[string]string
}

// FromQuery parses ListOptions from url.Values.
func (lo *ListOptions) FromQuery(q url.Values) error { _ = "STUB: not implemented"; return nil }

// default

// FIXME: This is a bit lazy, as "invalidxx,pinned" would result in a
// valid "pinned" filter.

// implicit default
