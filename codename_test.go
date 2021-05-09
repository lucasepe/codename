package codename

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

// Make sure the generated names exist
func TestGenerator(t *testing.T) {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 10; i++ {
		name := Generate(rng, 3, false)
		t.Logf("%v", name)
	}
}

func TestReturnsSameForSameSeed(t *testing.T) {
	var seed int64 = 1001
	rng := rand.New(rand.NewSource(seed))

	n1 := Generate(rng, 4, false)

	rng = rand.New(rand.NewSource(seed))
	n2 := Generate(rng, 4, false)

	t.Logf("%v -- %v", n1, n2)
	if n1 != n2 {
		t.Error("Should return same")
	}
}

func TestCryptoSeed(t *testing.T) {
	seed, err := NewCryptoSeed()
	if err != nil {
		t.Fatal(err)
	}

	rng := rand.New(rand.NewSource(seed))

	id := Generate(rng, 4, false)

	parts := strings.Split(id, "-")
	if got, want := len(parts[len(parts)-1]), 4; got != want {
		t.Fatalf("token length does not match got [%d], want [%d]", got, want)
	}

}
