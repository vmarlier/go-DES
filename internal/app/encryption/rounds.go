package encryption

import (
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

// subsitution of the right block to create a new 32 bits block
func subtitution(s []string) []string {
	// cut the block into 8 6bits blocks
	dividedBlocks := [][]string{s[:6], s[6:12], s[12:18], s[18:24], s[24:30], s[30:36], s[36:42], s[42:]}

	// declarations of the 8 subtitution matrixes
	substitutionMatrix1 := [16][4]int{[4]int{14, 0, 4, 15}, [4]int{4, 15, 1, 12}, [4]int{13, 7, 14, 8}, [4]int{1, 4, 8, 2},
		[4]int{2, 14, 13, 4}, [4]int{15, 2, 6, 9}, [4]int{11, 13, 2, 1}, [4]int{8, 1, 11, 7},
		[4]int{3, 10, 15, 5}, [4]int{10, 6, 12, 11}, [4]int{6, 12, 9, 3}, [4]int{12, 11, 7, 14},
		[4]int{5, 9, 3, 10}, [4]int{9, 5, 10, 0}, [4]int{0, 3, 5, 6}, [4]int{7, 8, 0, 13}}
	substitutionMatrix2 := [16][4]int{[4]int{15, 3, 0, 13}, [4]int{1, 13, 14, 8}, [4]int{8, 4, 7, 10}, [4]int{14, 7, 11, 1},
		[4]int{6, 15, 10, 3}, [4]int{11, 2, 4, 15}, [4]int{3, 8, 13, 4}, [4]int{4, 14, 1, 2},
		[4]int{9, 12, 5, 11}, [4]int{7, 0, 8, 6}, [4]int{2, 1, 12, 7}, [4]int{13, 10, 6, 12},
		[4]int{12, 6, 9, 0}, [4]int{0, 9, 3, 5}, [4]int{5, 11, 2, 14}, [4]int{10, 5, 15, 9}}
	substitutionMatrix3 := [16][4]int{[4]int{10, 13, 13, 1}, [4]int{0, 7, 6, 10}, [4]int{9, 0, 4, 13}, [4]int{14, 9, 9, 0},
		[4]int{6, 3, 8, 6}, [4]int{3, 4, 15, 9}, [4]int{15, 6, 3, 8}, [4]int{5, 10, 0, 7},
		[4]int{1, 2, 11, 4}, [4]int{13, 8, 1, 15}, [4]int{12, 5, 2, 14}, [4]int{7, 14, 12, 3},
		[4]int{11, 12, 5, 11}, [4]int{4, 11, 10, 5}, [4]int{2, 15, 14, 2}, [4]int{8, 1, 7, 12}}
	substitutionMatrix4 := [16][4]int{[4]int{7, 13, 10, 3}, [4]int{13, 8, 6, 15}, [4]int{14, 11, 9, 0}, [4]int{3, 5, 0, 6},
		[4]int{0, 6, 12, 10}, [4]int{6, 15, 11, 1}, [4]int{9, 0, 7, 13}, [4]int{10, 3, 13, 8},
		[4]int{1, 4, 15, 9}, [4]int{2, 7, 1, 4}, [4]int{8, 2, 3, 5}, [4]int{5, 12, 14, 11},
		[4]int{11, 1, 5, 12}, [4]int{12, 10, 2, 7}, [4]int{4, 14, 8, 2}, [4]int{15, 9, 4, 14}}
	substitutionMatrix5 := [16][4]int{[4]int{2, 14, 4, 11}, [4]int{12, 11, 2, 8}, [4]int{4, 2, 1, 12}, [4]int{1, 12, 11, 7},
		[4]int{7, 4, 10, 1}, [4]int{10, 7, 13, 14}, [4]int{11, 13, 7, 2}, [4]int{6, 1, 8, 13},
		[4]int{8, 5, 15, 6}, [4]int{5, 0, 9, 15}, [4]int{3, 15, 12, 0}, [4]int{15, 10, 5, 9},
		[4]int{13, 3, 6, 10}, [4]int{0, 9, 3, 4}, [4]int{14, 8, 0, 5}, [4]int{9, 6, 14, 3}}
	substitutionMatrix6 := [16][4]int{[4]int{12, 10, 9, 4}, [4]int{1, 15, 14, 3}, [4]int{10, 4, 15, 2}, [4]int{15, 2, 5, 12},
		[4]int{9, 7, 2, 9}, [4]int{2, 12, 8, 5}, [4]int{6, 9, 12, 15}, [4]int{8, 5, 3, 10},
		[4]int{0, 6, 7, 11}, [4]int{13, 1, 0, 14}, [4]int{3, 13, 4, 1}, [4]int{4, 14, 10, 7},
		[4]int{14, 0, 1, 6}, [4]int{7, 11, 13, 0}, [4]int{5, 3, 11, 8}, [4]int{11, 8, 6, 13}}
	substitutionMatrix7 := [16][4]int{[4]int{4, 13, 1, 6}, [4]int{11, 0, 4, 11}, [4]int{2, 11, 11, 13}, [4]int{14, 7, 13, 8},
		[4]int{15, 4, 12, 1}, [4]int{0, 9, 3, 4}, [4]int{8, 1, 7, 10}, [4]int{13, 10, 14, 7},
		[4]int{3, 14, 10, 9}, [4]int{12, 3, 15, 5}, [4]int{9, 5, 6, 0}, [4]int{7, 12, 8, 15},
		[4]int{5, 2, 0, 14}, [4]int{10, 15, 5, 2}, [4]int{6, 8, 9, 3}, [4]int{1, 6, 2, 12}}
	substitutionMatrix8 := [16][4]int{[4]int{13, 1, 7, 2}, [4]int{2, 15, 11, 1}, [4]int{8, 13, 4, 14}, [4]int{4, 8, 1, 7},
		[4]int{6, 10, 9, 4}, [4]int{15, 3, 12, 10}, [4]int{11, 7, 14, 8}, [4]int{1, 4, 2, 13},
		[4]int{10, 12, 0, 15}, [4]int{9, 5, 6, 12}, [4]int{3, 6, 10, 9}, [4]int{14, 11, 13, 0},
		[4]int{5, 0, 15, 3}, [4]int{0, 14, 3, 5}, [4]int{12, 9, 5, 6}, [4]int{7, 2, 8, 11}}

	var resultString string
	var substitutionResult int

	for i, dividedBlock := range dividedBlocks {
		// to determine the x you have to take the bits 2,3,4,5 on the concording dividedBlock and turn it into decimal
		xMatrix, _ := strconv.ParseUint(strings.Join(dividedBlock[1:5], ""), 2, 0)
		// to determine the y you have to take the bits 1,6 on the concording dividedBlock and turn it into decimal
		yMatrix, _ := strconv.ParseUint(strings.Join([]string{dividedBlock[0], dividedBlock[5]}, ""), 2, 0)

		// choose the substitution matrix matching the divided Block
		switch i {
		case 0:
			substitutionResult = substitutionMatrix1[xMatrix][yMatrix]
		case 1:
			substitutionResult = substitutionMatrix2[xMatrix][yMatrix]
		case 2:
			substitutionResult = substitutionMatrix3[xMatrix][yMatrix]
		case 3:
			substitutionResult = substitutionMatrix4[xMatrix][yMatrix]
		case 4:
			substitutionResult = substitutionMatrix5[xMatrix][yMatrix]
		case 5:
			substitutionResult = substitutionMatrix6[xMatrix][yMatrix]
		case 6:
			substitutionResult = substitutionMatrix7[xMatrix][yMatrix]
		case 7:
			substitutionResult = substitutionMatrix8[xMatrix][yMatrix]
		}

		// add a "0" bit if the result is inferior to 7 -> "111" -> "0111"
		if substitutionResult <= 1 {
			resultString += "000" + strconv.FormatUint(uint64(substitutionResult), 2)
		} else if substitutionResult <= 3 {
			resultString += "00" + strconv.FormatUint(uint64(substitutionResult), 2)
		} else if substitutionResult <= 7 {
			resultString += "0" + strconv.FormatUint(uint64(substitutionResult), 2)
		} else {
			resultString += strconv.FormatUint(uint64(substitutionResult), 2)
		}

	}

	return strings.Split(resultString, "")
}

func permutationP(s []string) []string {
	return []string{s[15], s[6], s[19], s[20], s[28], s[11], s[27], s[16],
		s[0], s[14], s[22], s[25], s[4], s[17], s[30], s[9],
		s[1], s[7], s[23], s[13], s[31], s[26], s[2], s[8],
		s[18], s[12], s[29], s[5], s[21], s[10], s[3], s[24]}
}

// Rounds is the most important part in the DES encryption
// The stages are Right block expension -> XOR Right block and Ki -> subtitution -> permutation -> XOR Right block and left block
func Rounds(binaryIP []string, keys []string, decrypt bool) (l16 []string, r16 []string) {
	// separate the message into 2 blocks
	leftBlock := binaryIP[:32]
	rightBlock := binaryIP[32:]

	for i := 0; i < 16; i++ {
		// r expansion
		rightBlockExpanded := expansion(rightBlock)

		// join the slices and turn them into decimal int
		rightBlockExpandedInt, _ := strconv.ParseUint(strings.Join(rightBlockExpanded, ""), 2, 0)
		var keyInt uint64

		if decrypt == false {
			keyInt, _ = strconv.ParseUint(keys[i], 2, 0)
		} else if decrypt == true {
			keyInt, _ = strconv.ParseUint(keys[15-i], 2, 0)
		}

		// XOR on the decimals int
		tmpUint := rightBlockExpandedInt ^ keyInt

		// convert back to base 2 string
		tmpString := strconv.FormatUint(tmpUint, 2)
		// Error can occur here cause in binary 0011 = 11, so we check the lenght and add 0 at the begining to have 32 bits
		for len(tmpString) < 48 {
			tmpString = "0" + tmpString
		}
		rightBlockExpanded = strings.Split(tmpString, "")

		// subtitution
		tmp := subtitution(rightBlockExpanded)

		// make the permutation with the substitution result
		tmp = permutationP(tmp)

		// XOR leftBlock and permutationPResult
		tmpUint, _ = strconv.ParseUint(strings.Join(tmp, ""), 2, 0)
		leftBlockInt, _ := strconv.ParseUint(strings.Join(leftBlock, ""), 2, 0)
		tmpUint = tmpUint ^ leftBlockInt
		tmpString = strconv.FormatUint(tmpUint, 2)

		// Error can occur here cause in binary 0011 = 11, so we check the lenght and add 0 at the begining to have 32 bits
		for len(tmpString) < 32 {
			tmpString = "0" + tmpString
		}
		tmp = strings.Split(tmpString, "")

		// defined xorResult as Ri+1 (rightBlock) and rightBlock as Li+1 (leftBlock)
		leftBlock = rightBlock
		rightBlock = tmp
	}

	// return the 2 blocs L16 and R16
	return leftBlock, rightBlock
}
