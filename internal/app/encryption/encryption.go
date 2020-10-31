package encryption

import (
	"fmt"
	"strings"

	"usefull.pkg/des/internal/pkg/binary"
)

// Ip function is the implementation of the IP permutation in the DES encryption
func ip(s []string) []string {

	return []string{s[57], s[49], s[41], s[33], s[25], s[17], s[9], s[1],
		s[59], s[51], s[43], s[35], s[27], s[19], s[11], s[3],
		s[61], s[53], s[45], s[37], s[29], s[21], s[13], s[5],
		s[63], s[55], s[47], s[39], s[31], s[23], s[15], s[7],
		s[56], s[48], s[40], s[32], s[24], s[16], s[8], s[0],
		s[58], s[50], s[42], s[34], s[26], s[18], s[10], s[3],
		s[60], s[52], s[44], s[36], s[28], s[20], s[12], s[4],
		s[62], s[54], s[46], s[38], s[30], s[22], s[14], s[7]}
}

// Ipl1 function is the implementation of the IP^-1 permutation in the DES encryption
func ipl1(s []string) []string {

	return []string{s[39], s[7], s[47], s[15], s[55], s[23], s[63], s[31],
		s[38], s[6], s[46], s[14], s[54], s[22], s[62], s[30],
		s[37], s[5], s[45], s[13], s[53], s[21], s[61], s[29],
		s[36], s[4], s[44], s[12], s[52], s[20], s[60], s[28],
		s[35], s[3], s[43], s[11], s[51], s[19], s[59], s[27],
		s[34], s[2], s[42], s[10], s[50], s[18], s[58], s[26],
		s[33], s[1], s[41], s[9], s[49], s[17], s[57], s[25],
		s[32], s[0], s[40], s[8], s[48], s[16], s[56], s[24]}
}

// EncryptMessage ...
func EncryptMessage(s string) {
	// convert the message into binary
	b := binary.StringToBinary(s)
	// split the binary string into a slice of strings
	m := strings.Split(b, "")
	fmt.Println("binary: ")
	fmt.Println(m)

	// make the IP inital permutation
	fmt.Println("permutation:")
	ti := ip(m)
	fmt.Println(ti)

	// generate the keys for the rounds
	keys := getKeys()

	// make the rounds
	for i, k := range keys {
		fmt.Println(rounds.Round(ti, k, i))
		break
	}

}
