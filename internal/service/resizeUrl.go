package service

import (
	"crypto/sha256"
	"encoding/hex"
)

func ResizeUrl(url string) string {
	hesh := sha256.Sum256([]byte(url))
	return hex.EncodeToString(hesh[:])[:8]
}
