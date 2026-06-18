package mathx

import "testing"

func TestRandomIntRange(t *testing.T) {
	got, err := RandomIntRange(1, 10)
	if err != nil {
		t.Fatalf("RandomIntRange(1, 10) returned unexpected error: %v", err)
	}

	inRange := got >= 1 && got < 10

	if !inRange {
		t.Errorf("RandomIntRange(1, 10) == %v; want value in [1, 10)", got)
	}
}
