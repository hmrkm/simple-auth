package usecase

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreateHash(src string) string {
	sha256ByteArr := sha256.Sum256([]byte(src))
	return hex.EncodeToString(sha256ByteArr[:])
}
