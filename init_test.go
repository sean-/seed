package seed_test

import (
	"testing"

	"github.com/sean-/seed"
)

func TestInit(t *testing.T) {
	_, seeded, err := seed.Init()
	if !seeded {
		t.Fatalf("Seeding failed: %v", err)
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
