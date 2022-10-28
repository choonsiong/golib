package mathx

import "testing"

func TestRandomIntRange(t *testing.T) {
	got := RandomIntRange(1, 10)

	inRange := got >= 1 && got < 10

	if !inRange {
		t.Errorf("RandomIntRange(1, 10) == %v; want value in [1, 10)", got)
	}
}
