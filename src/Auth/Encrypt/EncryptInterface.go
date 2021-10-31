package encrypt

type EncryptInterface interface {
	Encode(source string, salt string) string
}
