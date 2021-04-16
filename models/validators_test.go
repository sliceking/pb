package models

import "testing"

func TestIsEmailValid(t *testing.T) {
	testCase := []struct {
		email    string
		expected bool
	}{
		{"abc@google.com", true},
		{"1o2i3u1o23u12oiu@o2i3uo12u3o12iu3.com", false},
		{"mom@hotmail.com", true},
		{"avc@@google.com", false},
		{"no@google..com", false},
	}

	for _, tt := range testCase {
		got := IsEmailValid(tt.email)
		if got != tt.expected {
			t.Errorf("got %v, wanted %v", got, tt.expected)
		}
	}
}
