package main

import (
	"testing"
)

func TestParseIpv4(t *testing.T) {
	cases := []struct {
		input    string
		expected uint32
		err      error
	}{
		{"172.168.5.1", 2896692481, nil},
		{"192.168.1.1", 3232235777, nil},
		{"172 . 168.5.1", 2896692481, nil},
		{"0.0.0.0.0", 0, ErrInvalidIPv4Address},
		{"a.b.c.d", 0, ErrInvalidIPv4Address},
		{"1 72.168.5.1", 0, ErrInvalidIPv4Address},
		{"256.1.1.1", 0, ErrInvalidIPv4Address},
	}

	for _, tc := range cases {
		ip, err := parseIPv4(tc.input)
		if err != nil && tc.err == nil {
			t.Errorf("parseIPv4(%q) returned unexpected error: %v", tc.input, err)
		} else if err == nil && tc.err != nil {
			t.Errorf("parseIPv4(%q) expected error %v, but got nil", tc.input, tc.err)
		} else if ip != tc.expected {
			t.Errorf("parseIPv4(%q) = %d, expected %d", tc.input, ip, tc.expected)
		}
	}
}
