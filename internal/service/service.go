package service

import (
	"crypto/rand"
	"hash/crc32"
	"math"
	"math/big"
	"time"
)

const symbols = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890_"

func GenerateLink(url string) (string, error) {
	crc32q := crc32.MakeTable(0xD5828281)
	Checksum := crc32.Checksum(append([]byte(url), []byte(time.Now().String())...), crc32q)
	rndGenerator, err := rand.Int(rand.Reader, big.NewInt(int64(float64(Checksum)+math.Pow(float64(len(symbols)), 10))))
	if err != nil {
		return "", err
	}
	n := rndGenerator.Int64()
	resultLink := make([]byte, 10)
	for i := 0; i < 10; i++ {
		if n >= 0 {
			resultLink[i] = symbols[n%int64(len(symbols))]
			n /= int64(len(symbols))
		} else {
			resultLink[i] = symbols[i]
		}
	}
	return string(resultLink), nil
}
