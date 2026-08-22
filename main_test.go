package main

import "testing"

// The release gate runs `go test ./...` before tagging; a repo with no tests
// would pass vacuously — this pins that the demo exercises the gate.
func TestVersionDefaults(t *testing.T) {
	if version == "" || commit == "" {
		t.Fatal("stamp variables must never be empty strings")
	}
}
