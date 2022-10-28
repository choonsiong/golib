package mathx

import "testing"

func TestCompareFloats(t *testing.T) {
	tests := []struct {
		name string
		f1   float64
		f2   float64
		e    float64
		want bool
	}{
		{"0.0 == 0.0", 0.0, 0.0, 1e-9, true},
		{"3.141 == 3.1412", 3.141, 3.1412, 0.01, true},
		{"0.0 == 1.0", 0.0, 1.0, 0.1, false},
		{"-3.141 == -3.1412", -3.141, -3.1412, 0.01, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompareFloats(tt.f1, tt.f2, tt.e)
			if got != tt.want {
				t.Errorf("CompareFloats(%v, %v, %v) == %v; want %v", tt.f1, tt.f2, tt.e, got, tt.want)
			}
		})
	}
}
