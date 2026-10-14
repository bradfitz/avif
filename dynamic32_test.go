package avif

import (
	"strconv"
	"testing"
)

// TestNoDynamicOn32Bit checks that the build tags disable dynamic mode on
// every 32-bit architecture, where the native bindings aren't verified to
// work. They crash on 386.
func TestNoDynamicOn32Bit(t *testing.T) {
	if strconv.IntSize != 32 {
		t.Skip("not a 32-bit build")
	}
	if dynamicBuild {
		t.Fatal("dynamic mode is built in on a 32-bit architecture")
	}
}
