package trigonometry

import (
	"math"
	"testing"
)

func TestDegreeToRadian(t *testing.T) {
	tests := []struct {
		name   string
		degree float64
		want   float64
	}{
		{"0 degree", 0.0, 0},
		{"90 degree", 90.0, math.Pi / 2},
		{"180 degree", 180.0, math.Pi},
		{"270 degree", 270.0, (math.Pi * 3) / 2},
		{"360 degree", 360.0, math.Pi * 2},
		{"-180 degree", -180.0, -math.Pi},
		{"720 degree", 720.0, math.Pi * 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DegreeToRadian(tt.degree)
			if got != tt.want {
				t.Errorf("DegreeToRadian(%v) == %v; want %v", tt.degree, got, tt.want)
			}
		})
	}
}

func TestRadianToDegree(t *testing.T) {
	tests := []struct {
		name   string
		radian float64
		want   float64
	}{
		{"0 radian", 0.0, 0},
		{"Pi/2", math.Pi / 2, 90.0},
		{"Pi", math.Pi, 180.0},
		{"3Pi/2", (math.Pi * 3) / 2, 270.0},
		{"2Pi", math.Pi * 2, 360.0},
		{"-Pi", -math.Pi, -180.0},
		{"4Pi", math.Pi * 4, 720.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RadianToDegree(tt.radian)
			if got != tt.want {
				t.Errorf("RadianToDegree(%v) == %v; want %v", tt.radian, got, tt.want)
			}
		})
	}
}

func TestDegreeToDMS(t *testing.T) {
	tests := []struct {
		name   string
		degree float64
		want   Angle
	}{
		{"0 degree", 0.0, Angle{0, 0, 0}},
		{"90 degree", 90.0, Angle{90, 0, 0}},
		{"90.05", 90.5, Angle{90, 30, 0}},
		{"90.55 degree", 90.55, Angle{90, 32, 59}},
		{"38.27 degree", 38.27, Angle{38, 16, 12}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DegreeToDMS(tt.degree)
			if got.Degree != tt.want.Degree {
				t.Errorf("DegreeToDMS(%v) == %v degree; want %v", tt.degree, got.Degree, tt.want.Degree)
			}
			if got.Minute != tt.want.Minute {
				t.Errorf("DegreeToDMS(%v) == %v minute; want %v", tt.degree, got.Minute, tt.want.Minute)
			}
			if got.Second != tt.want.Second {
				t.Errorf("DegreeToDMS(%v) == %v second; want %v", tt.degree, got.Second, tt.want.Second)
			}
		})
	}
}

func TestRadianToDMS(t *testing.T) {
	tests := []struct {
		name   string
		radian float64
		want   Angle
	}{
		{"0 radian", 0.0, Angle{0, 0, 0}},
		{"Pi radian", math.Pi, Angle{180, 0, 0}},
		{"Pi/2 radian", math.Pi / 2, Angle{90, 0, 0}},
		{"3Pi/2 radian", (math.Pi * 3) / 2, Angle{270, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RadianToDMS(tt.radian)
			if got.Degree != tt.want.Degree {
				t.Errorf("DegreeToDMS(%v) == %v degree; want %v", tt.radian, got.Degree, tt.want.Degree)
			}
			if got.Minute != tt.want.Minute {
				t.Errorf("DegreeToDMS(%v) == %v minute; want %v", tt.radian, got.Minute, tt.want.Minute)
			}
			if got.Second != tt.want.Second {
				t.Errorf("DegreeToDMS(%v) == %v second; want %v", tt.radian, got.Second, tt.want.Second)
			}
		})
	}
}

func TestAngle_DMSToDegree(t *testing.T) {
	tests := []struct {
		name  string
		angle Angle
		want  float64
	}{
		{"zero angle", Angle{0, 0, 0}, 0.0},
		{"straight angle", Angle{180, 0, 0}, 180.0},
		{"right angle", Angle{90, 0, 0}, 90.0},
		{"complete angle", Angle{360, 0, 0}, 360.0},
		{"minus angle", Angle{-180, 0, 0}, -180.0},
		{"angle with minute", Angle{90, 60, 0}, 91.0},
		{"angle with second", Angle{90, 0, 3600}, 91.0},
		{"angle with minute and second", Angle{0, 60, 3600}, 2.0},
		{"angle with minute and second", Angle{72, 45, 9}, 72.7525},
		{"angle with minute and second", Angle{55, 36, 18}, 55.6050},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.angle.DMSToDegree()
			if got != tt.want {
				if got > tt.want {
					if got-tt.want > 0.0001 {
						t.Errorf("Angle.ToDegree() == %v; want %v", got, tt.want)
					}
				} else {
					t.Errorf("Angle.ToDegree() == %v; want %v", got, tt.want)
				}
			}
		})
	}
}

func TestAngle_DMSToRadian(t *testing.T) {
	tests := []struct {
		name  string
		angle Angle
		want  float64
	}{
		{"zero angle", Angle{0, 0, 0}, 0.0},
		{"straight angle", Angle{180, 0, 0}, math.Pi},
		{"right angle", Angle{90, 0, 0}, math.Pi / 2},
		{"complete angle", Angle{360, 0, 0}, math.Pi * 2},
		{"minus angle", Angle{-180, 0, 0}, -math.Pi},
		{"angle with minute", Angle{90, 60, 0}, 1.5882},
		{"angle with second", Angle{90, 0, 3600}, 1.5882},
		{"angle with minute and second", Angle{0, 60, 3600}, 0.0349},
		{"angle with minute and second", Angle{72, 45, 9}, 1.2697},
		{"angle with minute and second", Angle{55, 36, 18}, 0.9704},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.angle.DMSToRadian()
			if got != tt.want {
				if got > tt.want {
					if got-tt.want > 0.0001 {
						t.Errorf("Angle.ToRadian() == %v; want %v", got, tt.want)
					}
				} else {
					t.Errorf("Angle.ToRadian() == %v; want %v", got, tt.want)
				}
			}
		})
	}
}

func TestAngle_IsAcuteAngle(t *testing.T) {
	tests := []struct {
		name  string
		angle Angle
		want  bool
	}{
		{"70 degree", Angle{70, 0, 0}, true},
		{"90 degree", Angle{90, 0, 0}, false},
		{"90 degree 1 minute", Angle{90, 1, 0}, false},
		{"180 degree", Angle{180, 0, 0}, false},
		{"0 degree", Angle{0, 0, 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.angle.IsAcuteAngle()
			if got != tt.want {
				t.Errorf("Angle.IsAcuteAngle() == %v; want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_IsObtuseAngle(t *testing.T) {
	tests := []struct {
		name  string
		angle Angle
		want  bool
	}{
		{"70 degree", Angle{70, 0, 0}, false},
		{"90 degree", Angle{90, 0, 0}, false},
		{"90 degree 1 minute", Angle{90, 1, 0}, true},
		{"120 degree", Angle{120, 0, 0}, true},
		{"180 degree", Angle{180, 0, 0}, false},
		{"0 degree", Angle{0, 0, 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.angle.IsObtuseAngle()
			if got != tt.want {
				t.Errorf("Angle.IsObtuseAngle() == %v; want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_IsRightAngle(t *testing.T) {
	tests := []struct {
		name  string
		angle Angle
		want  bool
	}{
		{"90 degree", Angle{90, 0, 0}, true},
		{"90 degree 1 minute", Angle{90, 1, 0}, false},
		{"180 degree", Angle{180, 0, 0}, false},
		{"0 degree", Angle{0, 0, 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.angle.IsRightAngle()
			if got != tt.want {
				t.Errorf("Angle.IsRightAngle() == %v; want %v", got, tt.want)
			}
		})
	}
}

func TestAngle_IsCompleteAngle(t *testing.T) {
	tests := []struct {
		name  string
		angle Angle
		want  bool
	}{
		{"90 degree", Angle{90, 0, 0}, false},
		{"120 degree", Angle{120, 0, 0}, false},
		{"180 degree", Angle{180, 0, 0}, false},
		{"0 degree", Angle{0, 0, 0}, false},
		{"360 degree", Angle{360, 0, 0}, true},
		{"720 degree", Angle{720, 0, 0}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.angle.IsCompleteAngle()
			if got != tt.want {
				t.Errorf("Angle.IsCompleteAngle() == %v; want %v", got, tt.want)
			}
		})
	}
}
