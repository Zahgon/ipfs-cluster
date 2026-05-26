package ipfsproxy

import (
	"strings"
)

// MultiError contains the results of multiple errors.
type multiError struct {
	err strings.Builder
}

func (e *multiError) add(err string) { _ = "STUB: not implemented"; return }

func (e *multiError) Error() string { _ = "STUB: not implemented"; return "" }
