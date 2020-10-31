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

func subtitution(s []string) {
}

// Round is the most important part in the DES encryption
// The stages are Right block expension -> XOR Right block and Ki -> subtitution -> permutation -> XOR Right block and left block
func Round(binaryIP []string, key string, i int) []string {
	fmt.Printf("Round n°%v\n", i)

	// separate the message into 2 blocks
	leftBlock := binaryIP[:32]
	rightBlock := binaryIP[32:]
	fmt.Println("separation:")
	fmt.Print(leftBlock)
	fmt.Println(rightBlock)
	fmt.Println(len(leftBlock))
	fmt.Println(len(rightBlock))

	// r expansion
	rightBlockExpanded := expansion(rightBlock)
	fmt.Println("expansion:")
	fmt.Print("re :")
	fmt.Println(rightBlockExpanded)
	fmt.Println(len(rightBlockExpanded))

	fmt.Println("key:")
	fmt.Println(key)

	// join the slices and turn them into decimal int
	rightBlockExpandedInt, _ := strconv.ParseUint(strings.Join(rightBlockExpanded, ""), 2, 0)
	keyInt, _ := strconv.ParseUint(key, 2, 0)

	// XOR on the decimals int
	rightBlockXorKey := rightBlockExpandedInt ^ keyInt
	fmt.Println("XOR:")
	fmt.Println(rightBlockXorKey)

	// convert back to base 2 string
	rightBlockExpanded = strings.Split(strconv.FormatUint(rightBlockXorKey, 2), "")
	fmt.Println("Back to binary:")
	fmt.Println(rightBlockExpanded)

	// subtitution

	return []string{}
}
