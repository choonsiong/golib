package auth

import "testing"

func TestPassword_CheckPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		auth     *Password
		want     bool
	}{
		{"minimum password length ok", "abcd", &Password{4, false, false, false, false}, true},
		{"minimum password length not ok", "abc", &Password{4, false, false, false, false}, false},
		{"want uppercase ok", "aBcd", &Password{4, true, false, false, false}, true},
		{"want uppercase not ok", "abcd", &Password{4, true, false, false, false}, false},
		{"want lowercase ok", "ABcD", &Password{4, false, true, false, false}, true},
		{"want lowercase not ok", "1234ABC", &Password{4, false, true, false, false}, false},
		{"want mixed cases ok", "1234ABCdef", &Password{4, true, true, false, false}, true},
		{"want mixed cases not ok", "1234567", &Password{4, true, true, false, false}, false},
		{"want number ok", "1234ABCdef", &Password{4, false, false, true, false}, true},
		{"want number not ok", "ABCdef", &Password{4, false, false, true, false}, false},
		{"want special character ok", "1234ABCdef@!#", &Password{4, false, false, false, true}, true},
		{"want special character not ok", "1234ABCdef", &Password{4, false, false, false, true}, false},
		{"want all ok", "1234ABCdef@!#", &Password{4, true, true, true, true}, true},
		{"want all not ok", "1234@!#", &Password{4, true, true, true, true}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.auth.CheckPassword(tt.password)
			if got != tt.want {
				t.Errorf("CheckPassword(%s) == %t; want %t", tt.password, got, tt.want)
			}
		})
	}
}
