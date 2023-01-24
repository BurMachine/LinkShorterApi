package service

import (
	"github.com/speps/go-hashids"
)

const symbols = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890_"

func GenerateLink(url string) (string, error) {
	hd := hashids.NewData()
	hd.Salt = url
	hd.MinLength = 10
	hd.Alphabet = symbols
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{45, 434, 1313, 99})
	e = e[:10]
	println(e)
	return e, nil
}
