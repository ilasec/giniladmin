package utils

import (
	"crypto/rand"
	"encoding/base64"
	"strconv"
)

func Atoi(v string) int {
	i, e := strconv.Atoi(v)
	if e != nil {
		i = 0
	}
	return i
}

func RandStringV2(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
