package converter

import (
	"encoding/hex"
	"strconv"
)

//Konverterer int64 til hex
func IntToHex(n int64) []byte {
	return []byte(strconv.FormatInt(n, 16))
}

func StringToHex(s string) []byte {
	src := []byte(s)

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	return dst
}

func Add_0a(s string) string {
	return s + "0a"
}

func Remove_0a(s string) string {
	if last := len(s) - 1; last >= 0 && s[last] == '\n' {
		s = s[:last]
	}
	return s
}
