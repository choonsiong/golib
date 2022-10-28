package mathx

import (
	"math/rand"
	"time"
)

// RandomIntRange returns a random int within min, max but not includes max.
func RandomIntRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
