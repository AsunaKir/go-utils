package sm4

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/tjfoc/gmsm/sm4"
)

const (
	sm4Secret = "" // sm4密钥
)

// EncryptSm4 加密
func EncryptSm4(tmp []byte) (string, error) {
	key, err := hex.DecodeString(sm4Secret)
	if err != nil {
		return "", err
	}
	str, err := sm4.Sm4Ecb(key, tmp, true)
	if err != nil {
		return "", err
	}
	return Base64Encode(str), nil
}

// DecryptSm4 解密
func DecryptSm4(tmp []byte) ([]byte, error) {
	msg, err := Base64Decode(string(tmp))
	if err != nil {
		return nil, err
	}
	key, err := hex.DecodeString(sm4Secret)
	if err != nil {
		return nil, err
	}
	str, err := sm4.Sm4Ecb(key, msg, false)
	if err != nil {
		return nil, err
	}
	return str, nil
}

// Base64Encode base64 编码
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode base64 解码
func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
