package encryption

import (
	"fmt"
	"strconv"
	"strings"
)

// Expansion is the implementation of the expansive permutation in the DES algorithm
// the function take a slice of 32 strings and return a slice of 48 after the expansion
func expansion(s []string) []string {

	return []string{s[31], s[0], s[1], s[2], s[3], s[4], s[3], s[4],
		s[5], s[6], s[7], s[8], s[7], s[8], s[9], s[10],
		s[11], s[12], s[11], s[12], s[13], s[14], s[15], s[16],
		s[15], s[16], s[17], s[18], s[19], s[20], s[19], s[20],
		s[21], s[22], s[23], s[24], s[23], s[24], s[25], s[26],
		s[27], s[28], s[27], s[28], s[29], s[30], s[31], s[0]}
}

// Round is the most important part in the DES encryption
// The stages are Right block expension -> XOR Right block and Ki -> subtitution -> permutation -> XOR Right block and left block
func Round(ti []string, k string, i int) []string {
	fmt.Printf("Round n°%v\n", i)

	// separate the message into 2 blocks
	l := ti[:32]
	r := ti[32:]
	fmt.Println("separation:")
	fmt.Print(l)
	fmt.Println(r)
	fmt.Println(len(l))
	fmt.Println(len(r))

	// r expension
	re := expansion(r)
	fmt.Println("expansion:")
	fmt.Print("re :")
	fmt.Println(re)
	fmt.Println(len(re))

	fmt.Println("key:")
	fmt.Println(k)

	// join the slices and turn them into decimal int
	ri, _ := strconv.ParseUint(strings.Join(re, ""), 2, 0)
	ki, _ := strconv.ParseUint(k, 2, 0)

	// XOR on the decimals int
	rx := ri ^ ki
	fmt.Println("XOR:")
	fmt.Println(rx)

	// Convert back to base 2 string
	re = strings.Split(strconv.FormatUint(rx, 2), "")
	fmt.Println("Back to binary:")
	fmt.Println(re)

	// Fonction de substitution

	return []string{}
}
