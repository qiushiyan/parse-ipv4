package main

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidIPv4Address = errors.New("invalid IPv4 address")

func parseIPv4(s string) (uint32, error) {
	var result uint32
	parts := strings.Split(s, ".")
	if len(parts) != 4 {
		return 0, ErrInvalidIPv4Address
	}

	for i, part := range parts {
		part = strings.TrimSpace(part)
		// allow space between number and dots
		if i == 0 && len(part) > 1 && part[0] == '0' {
			return 0, ErrInvalidIPv4Address
		}
		value, err := strconv.Atoi(part)
		if err != nil {
			return 0, ErrInvalidIPv4Address
		}
		if value < 0 || value > 255 {
			return 0, ErrInvalidIPv4Address
		}
		result = (result << 8) | uint32(value)
	}

	return result, nil
}
