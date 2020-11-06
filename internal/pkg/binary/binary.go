package binary

import (
	"fmt"
	"strconv"
)

// StringToBinary take a string and turn it into a base 2 string
func StringToBinary(s string) (res string) {
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}

// ToHex ..
func ToHex() {}

// ToString take a string and turn it into a ascii string
func ToString(s string) (res string) {
	var arr [8]string
	j := 0
	for i, c := range s {
		if i%8 == 0 && i != 0 {
			j++
		}
		arr[j] += string(c)
	}

	for _, a := range arr {
		tmp, _ := strconv.ParseUint(a, 2, 64)
		res += string(tmp)
	}

	return res
}
