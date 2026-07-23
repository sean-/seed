package seed_test

import (
	"testing"

	"github.com/sean-/seed"
)

func TestInit(t *testing.T) {
	secure, err := seed.Init()
	if !secure {
		t.Fatalf("Failed to securely seed: %v", err)
	}
}

func TestMustInit(t *testing.T) {
	seed.MustInit()

	if !seed.Seeded() {
		t.Fatalf("MustInit() failed to seed")
	}

	if !seed.Secure() {
		t.Fatalf("MustInit() failed to securely seed")
	}
}

func TestInitConcurrent(t *testing.T) {
	// Should not race or panic when many goroutines call Init.
	const n = 32
	errc := make(chan error, n)
	for i := 0; i < n; i++ {
		go func() {
			_, err := seed.Init()
			errc <- err
		}()
	}
	for i := 0; i < n; i++ {
		if err := <-errc; err != nil {
			t.Fatalf("Init: %v", err)
		}
	}
	if !seed.Seeded() {
		t.Fatal("expected Seeded after concurrent Init")
	}
}
