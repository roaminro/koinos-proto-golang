package canonical

import "errors"

var (
	// ErrNoCanonical is returned when the object has no valid canonical serialization
	ErrNoCanonical = errors.New("no canonical serialization")
)
