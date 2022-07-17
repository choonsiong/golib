package finance

import "testing"

func TestSimpleInterest(t *testing.T) {
	tests := []struct {
		name       string
		principal  float64
		annualRate float64
		timePeriod float64
		want       float64
	}{
		{"positive interest", 200, 0.08, 4, 64},
		{"negative interest", 200, -0.08, 4, -64},
		{"zero interest", 200, 0, 4, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SimpleInterest(tt.principal, tt.annualRate, tt.timePeriod)
			if got != tt.want {
				t.Errorf("CalculateSimpleInterest(%v, %v, %v) == %v; want %v", tt.principal, tt.annualRate, tt.timePeriod, got, tt.want)
			}
		})
	}
}

func TestTotalAmountPlusInterest(t *testing.T) {
	tests := []struct {
		name       string
		principal  float64
		annualRate float64
		timePeriod float64
		want       float64
	}{
		{"positive interest", 200, 0.08, 4, 264},
		{"negative interest", 200, -0.08, 4, 136},
		{"zero interest", 200, 0, 4, 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TotalAmountPlusInterest(tt.principal, tt.annualRate, tt.timePeriod)
			if got != tt.want {
				t.Errorf("CalculateSimpleInterest(%v, %v, %v) == %v; want %v", tt.principal, tt.annualRate, tt.timePeriod, got, tt.want)
			}
		})
	}
}
