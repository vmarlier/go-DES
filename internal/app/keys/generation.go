package keys

import (
	"strings"

	"go-des/internal/pkg/binary"
)

// Pc1 function is the implementation of the PC-1 permutation in the keys generating process
func pc1(s []string) ([]string, []string) {

	return []string{
			s[56], s[48], s[40], s[32], s[24], s[16], s[8],
			s[0], s[57], s[49], s[41], s[33], s[25], s[17],
			s[9], s[1], s[58], s[50], s[42], s[34], s[26],
			s[18], s[10], s[2], s[59], s[51], s[43], s[35]},
		[]string{
			s[62], s[54], s[46], s[38], s[30], s[22], s[14],
			s[6], s[61], s[53], s[45], s[37], s[29], s[21],
			s[13], s[5], s[60], s[52], s[44], s[36], s[28],
			s[20], s[12], s[4], s[27], s[19], s[11], s[3]}
}

// LeftShift function provide a left shift of i case in a string slice
func leftShift(s []string, i int) []string {

	// Choose if the left shift is 1 or 2
	if i == 1 {

		// isolate the s[0]
		first := s[0]
		// loop on the block to shift
		for x := range s {
			if x == len(s)-1 {
				s[x] = first
			} else {
				s[x] = s[x+1]
			}
		}

	} else if i == 2 {

		// isolate the s[0] and s[1]
		firsts := []string{s[0], s[1]}
		// loop on the block to shift
		for x := range s {
			if x == len(s)-2 {
				s[x] = firsts[0]
			} else if x == len(s)-1 {
				s[x] = firsts[1]
			} else {
				s[x] = s[x+2]
			}
		}
	}

	return s
}

// Pc2 function is the implementation of the PC-2 permutation in the keys generating process
func pc2(s []string) []string {

	return []string{
		s[13], s[16], s[10], s[23], s[0], s[4], s[2], s[27],
		s[14], s[5], s[20], s[9], s[22], s[18], s[11], s[3],
		s[25], s[7], s[15], s[6], s[26], s[19], s[12], s[1],
		s[40], s[51], s[30], s[36], s[46], s[54], s[29], s[39],
		s[50], s[44], s[32], s[47], s[43], s[48], s[38], s[55],
		s[33], s[52], s[45], s[41], s[49], s[35], s[28], s[31]}
}

// GenerateKeys will take a ascii key (8 characters) and return 16 keys of 48 bits
func GenerateKeys(str string) (keys []string) {
	// turn the ascii key into binary key (64 bits)
	initialKey := binary.StringToBinary(str)

	// split the 64 bits key into a slice of string
	initialKeySlice := strings.Split(initialKey, "")

	// make the PC-1 permutation and get the c and d blocks (28 bits each)
	leftBlock, rightBlock := pc1(initialKeySlice)

	// create the leftShift Index
	lsIndex := []int{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

	for i := 0; i < 16; i++ {
		// leftShift on c and d blocks
		leftBlock = leftShift(leftBlock, lsIndex[i])
		rightBlock = leftShift(rightBlock, lsIndex[i])
		// concatenate the 2 blocks
		concatenateKey := append(leftBlock, rightBlock...)

		// make the PC-2 permutation
		keys = append(keys, strings.Join(pc2(concatenateKey), ""))
	}

	return keys
}
