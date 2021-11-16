package encrypt

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey, publicKey := GenerateRSAKey()
	fmt.Println(string(privateKey))
	fmt.Println(string(publicKey))
}

func TestRsaSignWithSha256(t *testing.T) {
	privateKey, _ := GenerateRSAKey()
	data := []byte("Test")
	result := RsaSignWithSha256(data, privateKey)
	fmt.Println(hex.EncodeToString(result))
}
