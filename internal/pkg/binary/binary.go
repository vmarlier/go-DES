package binary

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

// StringToBinary take a ASCII string return a base 2 string
func StringToBinary(s string) (res string) {
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

// ToHex take a ascii string and return a hexadecimal string
func ToHex(s string) string {
	return hex.EncodeToString([]byte(s))
}

// HexToBinary take a hexa string and return a binary string
func HexToBinary(s string) string {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	return StringToBinary(string(decoded))
}

// ToB64 take a ascii string and return a base64 string
func ToB64(s string) string {
	return ""
}

// B64ToBinary take a base64 string and return a binary string
func B64ToBinary(s string) string {
	return ""
}

// ToString take a string of 64 bits and return an ascii string
func ToString(s string) (res string) {
	arr := make([]string, 8)
	j := 0
	for i, c := range s {
		if i%8 == 0 && i != 0 {
			j++
		}
		arr[j] += string(c)
	}

	for _, a := range arr {
		tmp, _ := strconv.ParseUint(a, 2, 0)
		res += string(tmp)
	}

	return res
}
