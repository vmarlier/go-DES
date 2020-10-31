package rounds

import (
	"fmt"
	"strconv"
	"strings"

	"usefull.pkg/des/internal/pkg/encryptionfunctions"
)

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
	re := encryptionfunctions.Expansion(r)
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
