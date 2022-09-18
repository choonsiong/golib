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

			if tt.want != got {
				t.Errorf("DescriptiveName(%q) == %q; want %q", tt.code, got, tt.want)
			}
		})
	}
}
