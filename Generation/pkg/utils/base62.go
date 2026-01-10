package utils

import (
	"errors"
	"strings"
)

const (
	alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base     = int64(62)
)

// Base62Encode converts an integer to a Base62 string
func Base62Encode(id int64) string {
	if id == 0 {
		return string(alphabet[0])
	}

	var chars []byte
	n := id
	// Since snowflake IDs are positive, we don't need to handle negative numbers specifically
	// but good to be safe with absolute value if reused
	if n < 0 {
		n = -n
	}

	for n > 0 {
		remainder := n % base
		chars = append(chars, alphabet[remainder])
		n = n / base
	}

	// Reverse the slice
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

// Base62Decode converts a Base62 string back to an integer
func Base62Decode(s string) (int64, error) {
	var id int64
	for _, char := range s {
		index := strings.IndexRune(alphabet, char)
		if index == -1 {
			return 0, errors.New("invalid character in base62 string")
		}
		id = id*base + int64(index)
	}
	return id, nil
}
