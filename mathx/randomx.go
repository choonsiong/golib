package mathx

import (
	"crypto/rand"
	"math/big"
)

// RandomIntRange returns a random int within [min, max), i.e. min is
// inclusive and max is exclusive. It uses crypto/rand so that the result
// is suitable for security-sensitive use.
func RandomIntRange(min, max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + min, nil
}
