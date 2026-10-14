//go:build (!linux && !darwin && !windows) || android || nodynamic || arm || 386 || mips || mipsle

package avif

import (
	"bytes"
	"testing"
)

func TestWASMFallback(t *testing.T) {
	if dynamic {
		t.Fatal("fallback builds must use the WASM decoder")
	}
	if err := Dynamic(); err == nil {
		t.Fatal("Dynamic returned nil with dynamic loading disabled")
	}
	img, err := Decode(bytes.NewReader(testAvif8))
	if err != nil {
		t.Fatal(err)
	}
	if got := img.Bounds().Size(); got.X != 512 || got.Y != 512 {
		t.Fatalf("dimensions = %v; want 512x512", got)
	}
}
