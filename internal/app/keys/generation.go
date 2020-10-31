package keys

import (
	"strings"

	"usefull.pkg/des/internal/pkg/binary"
	"usefull.pkg/des/internal/pkg/encryptionfunctions"
	"usefull.pkg/des/internal/pkg/permutations"
)

// GenerateKeys will take a ascii key (8 characters) and return 16 keys of 48 bits
func GenerateKeys(str string) (keys []string) {
	// turn the ascii key into binary key (64 bits)
	ki := binary.StringToBinary(str)

	// split the 64 bits key into a slice of string
	kr := strings.Split(ki, "")

	// make the PC-1 permutation and get the c and d blocks (28 bits each)
	c, d := permutations.Pc1(kr)

	// create the leftShift Index
	lsIndex := []int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

	for i := 0; i < 16; i++ {
		// leftShift on c and d blocks
		c = encryptionfunctions.LeftShift(c, lsIndex[i])
		d = encryptionfunctions.LeftShift(d, lsIndex[i])
		// concatenate the 2 blocks
		cd := append(c, d...)

		// make the PC-2 permutation
		keys = append(keys, strings.Join(permutations.Pc2(cd), ""))
	}

	return keys
}
