package timezone

import (
	"testing"
)

func TestDescriptiveName(t *testing.T) {
	tests := []struct {
		name string
		code string
		want string
	}{
		{
			name: "valid timezone",
			code: "Asia/Tokyo",
			want: "Tokyo",
		},
		{
			name: "invalid timezone",
			code: "Asia/Petaling_Jaya",
			want: "",
		},
		{
			name: "empty timezone",
			code: "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DescriptiveName(tt.code)

			if got != tt.want {
				t.Errorf("DescriptiveName(%q) == %q; want %q", tt.code, got, tt.want)
			}
		})
	}
}

func TestHasTimezone(t *testing.T) {
	tests := []struct {
		name string
		tz   string
		want bool
	}{
		{"Asia/Tokyo", "Asia/Tokyo", true},
		{"Empty", "", false},
		{"Asia/London", "Asia/London", false},
		{"foobar", "foobar", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasTimezone(tt.tz)

			if got != tt.want {
				t.Errorf("HasTimezone(%s) == %v; want %v", tt.tz, got, tt.want)
			}
		})
	}
}
