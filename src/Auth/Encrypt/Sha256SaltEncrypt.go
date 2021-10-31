package encrypt

import (
	"crypto/sha256"
	"encoding/hex"
)

type Sha256SaltEncrypt struct {
	EncryptInterface
}

func (encrypt *Sha256SaltEncrypt) Encode(source string, salt string) string {
	source = source + salt
	bytes := sha256.Sum256([]byte(source))
	return hex.EncodeToString(bytes[:])
}
