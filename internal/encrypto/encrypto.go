package encrypto

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
)

// 对每个数据反复hmac
func HmacSha256Replace(values ...string) ([]byte, error) {
	if len(values) == 0 {
		return []byte{}, nil
	}
	buffer := bytes.NewBuffer(make([]byte, 0, 512))
	privateKey := []byte(values[0])

	for _, v := range values[1:] {
		if _, err := buffer.WriteString(v); err != nil {
			return nil, err
		}
		mac := hmac.New(sha256.New, privateKey)
		mac.Write(buffer.Bytes())
		privateKey = mac.Sum(nil)
		buffer.Reset()
	}

	return privateKey, nil
}

func Hmac(key, value []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(value)
	return mac.Sum(nil)
}
