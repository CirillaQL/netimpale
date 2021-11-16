package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"netimpale/utils/log"
)

var LOG = log.LOG

// GenerateRSAKey 生成
func GenerateRSAKey() (privateKey, publicKey []byte) {
	//生成私钥文件
	pKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		LOG.Errorf("Generate PrivateKey failed, Error: %v", err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(pKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privateKey = pem.EncodeToMemory(block)
	//公钥部分
	pbKey := &pKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(pbKey)
	if err != nil {
		LOG.Errorf("Generate PublicKey failed, Error: %v", err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	publicKey = pem.EncodeToMemory(block)
	return
}

// RsaSignWithSha256 使用私钥进行签名
func RsaSignWithSha256(data []byte, keyBytes []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		LOG.Error("Private Key is Error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		LOG.Errorf("ParsePKCS8PrivateKey Error: %v", err)
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		LOG.Errorf("Error from signing: %v", err)
		panic(err)
	}

	return signature
}

// RsaVerySignWithSha256 使用公钥对信息进行验证，确定是由正确的私钥进行签名
func RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		LOG.Error("Private Key is Error")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}
