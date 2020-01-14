package schemamisc

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256Sum(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}
