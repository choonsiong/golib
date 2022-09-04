package file

import (
	"errors"
	"testing"
)

func Test_convertToBinary(t *testing.T) {
	tests := []struct {
		name        string
		permissions string
		want        string
		wantErr     error
	}{
		{"-rwxrwxrwx", "-rwxrwxrwx", "111111111", nil},
		{"-abcrwxrwx", "-abcrwxrwx", "", ErrInvalidTriplet},
		{"-rwxabcrwx", "-rwxabcrwx", "", ErrInvalidTriplet},
		{"-rwxrwxabc", "-rwxrwxabc", "", ErrInvalidTriplet},
		{"empty", "", "", ErrInvalidTriplet},
		{"all invalid", "aaaaaaaaaa", "", ErrInvalidTriplet},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertToBinary(tt.permissions)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("convertToBinary(%v), want error %v; got nil", tt.permissions, tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("convertToBinary(%v), want error %v; got %v", tt.permissions, tt.wantErr, err)
				}
			}

			if got != tt.want {
				t.Errorf("convertToBinary(%v) == %q; want %q", tt.permissions, got, tt.want)
			}
		})
	}
}

func Test_tripletToBinary(t *testing.T) {
	tests := []struct {
		name    string
		triplet string
		want    string
		wantErr error
	}{
		{"rwx", "rwx", "111", nil},
		{"-wx", "-wx", "011", nil},
		{"--x", "--x", "001", nil},
		{"---", "---", "000", nil},
		{"r-x", "r-x", "101", nil},
		{"r--", "r--", "100", nil},
		{"rw-", "rw-", "110", nil},
		{"-w-", "-w-", "010", nil},
		{"empty", "", "", ErrInvalidTriplet},
		{"invalid triplet", "abc", "", ErrInvalidTriplet},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tripletToBinary(tt.triplet)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("tripletToBinary(%v), want error %v; got nil", tt.triplet, tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("tripletToBinary(%v), want error %v; got %v", tt.triplet, tt.wantErr, err)
				}
			}

			if got != tt.want {
				t.Errorf("tripletToBinary(%v) == %q; want %q", tt.triplet, got, tt.want)
			}
		})
	}
}
