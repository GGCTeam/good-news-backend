package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

// MakeHash <function>
// is used to create hash from string
func MakeHash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}
