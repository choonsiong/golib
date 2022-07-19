package compare

import "testing"

func TestFloat64Maps(t *testing.T) {
	tests := []struct {
		name string
		map1 map[float64]float64
		map2 map[float64]float64
		want bool
	}{
		{"identical maps", map[float64]float64{1.0: 1.0, 2.0: 2.0}, map[float64]float64{1.0: 1.0, 2.0: 2.0}, true},
		{"different maps", map[float64]float64{1.0: 1.0, 2.0: 2.0}, map[float64]float64{1.0: 2.0, 2.0: 3.0}, false},
		{"maps with different length", map[float64]float64{1.0: 1.0}, map[float64]float64{1.0: 2.0, 2.0: 3.0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float64Maps(tt.map1, tt.map2)
			if got != tt.want {
				t.Errorf("Float64Maps(%v, %v) == %v; want %v", tt.map1, tt.map2, got, tt.want)
			}
		})
	}
}
