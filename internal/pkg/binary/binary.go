package binary

import "fmt"

// StringToBinary take a string and transform it into a base 2 string
func StringToBinary(s string) (res string) {
	for _, c := range s {
		res = fmt.Sprintf("%s%.8b", res, c)
	}
	return res
}
