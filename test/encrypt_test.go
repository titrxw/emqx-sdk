package gotest

import (
	encrypt "github.com/titrxw/emqx-sdk/src/Auth/Encrypt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	t.Run("testEncrypt", func(t *testing.T) {
		encrypt := new(encrypt.Sha256SaltEncrypt)
		encrypt.Encode("123456test", "123456")
		t.Skipped()
	})
}
