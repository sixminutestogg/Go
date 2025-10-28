package main

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
)

func main() {

	fmt.Println("签名验证是为了保证接口安全和识别调用方身份，同时还需要满足以下几点：\n\n- 可变性：每次的签名必须是不一样的。\n- 时效性：每次请求的时效性，过期作废。\n- 唯一性：每次的签名是唯一的。\n- 完整性：能够对传入数据进行验证，防止篡改。")

}

func MD5Hash(input string) string {
	fmt.Println(" MD5Hash 返回 input 的十六进制 MD5 值")
	sum := md5.Sum([]byte(input))
	return hex.EncodeToString(sum[:])
}

func MD5WithSalt(input, slat string) string {
	fmt.Println(" MD5WithSalt 更常用：加盐")
	hash := md5.New()
	hash.Write([]byte(input))
	hash.Write([]byte(slat))
	return hex.EncodeToString(hash.Sum(nil))
}

func AESGCMEncrypt(plain, key []byte) (string, error) {
	fmt.Println("AESGCMEncrypt 使用 AES-GCM 加密，返回 base64(nonce|ciphertext)")
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("key length must be 16/24/32")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	noce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, noce)
	if err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nil, noce, plain, nil)

	out := append(noce, ciphertext...)

	return base64.StdEncoding.EncodeToString(out), nil
}

func AESGCMDecrypt(cipherB64 string, key []byte) (string, error) {
	fmt.Println("AESGCMDecrypt 解密 base64(nonce|ciphertext)")

	data, err := base64.StdEncoding.DecodeString(cipherB64)
	if err != nil {
		return "", err
	}
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	if len(data) < gcm.NonceSize() {
		return "", fmt.Errorf("ciphertext too short")
	}
	return "", nil
}

// 生成 RSA 私钥/公钥（PEM 编码）
//
// 使用公钥加密 / 私钥解密（适合短消息）
//
// 使用私钥签名（PKCS1v15 + SHA256）和用公钥验证签名
func GenerateRSAKey(bits int) (privPEM, pubPEM []byte, err error) {
	fmt.Println("GenerateRSAKey 生成 RSA 私钥，返回 PEM 编码的私钥和公钥")
	key, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	privDER := x509.MarshalPKCS1PrivateKey(key)
	privBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privDER}
	privPEM = pem.EncodeToMemory(privBlock)

	pubDER := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	pubBlock := &pem.Block{Type: "RSA PUBLIC KEY", Bytes: pubDER}
	pubPEM = pem.EncodeToMemory(pubBlock)
	return
}

func RSAPublicEncrypt(pubPEM []byte, plain []byte) (string, error) {

	fmt.Println("RSAPublicEncrypt 用公钥（PEM）加密")
	block, _ := pem.Decode(pubPEM)
	if block == nil {
		return "", fmt.Errorf("invalid public pem")
	}

	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)

	if err != nil {
		return "", err
	}
	ct, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, plain, nil)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(ct), nil

}

func RSAPrivateDecrypt(priPem []byte, cipherB64 string) ([]byte, error) {
	fmt.Println("RSAPrivateDecrypt 私钥解密")
	block, _ := pem.Decode(priPem)
	if block == nil {
		return nil, fmt.Errorf("invalid private pem")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	ct, err := base64.StdEncoding.DecodeString(cipherB64)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, ct, nil)
}

func RSASign(priPem []byte, data []byte) (string, error) {
	fmt.Println("RSASign 使用私钥签名（PKCS1v15 + SHA256），返回 base64(sig)")
	block, _ := pem.Decode(priPem)
	if block == nil {
		return "", fmt.Errorf("invalid private pem")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		return "", err
	}
	h := sha256.Sum256(data)
	sig, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, h[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}
