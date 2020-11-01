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

// subsitution of the right block to create a new 32 bits block
func subtitution(s []string) {
	// We cut the block into 8 6bits blocks
	dividedBlocks := [][]string{s[:6], s[6:12], s[12:18], s[18:24], s[24:30], s[30:36], s[36:42], s[42:]}

	fmt.Println("divided Right Block: ")
	fmt.Println(len(dividedBlocks))

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
	substiitutionMatrix5 := [16][4]int{[4]int{2, 14, 4, 11}, [4]int{12, 11, 2, 8}, [4]int{4, 2, 1, 12}, [4]int{1, 12, 11, 7},
		[4]int{7, 4, 10, 1}, [4]int{10, 7, 13, 14}, [4]int{11, 13, 7, 2}, [4]int{6, 1, 8, 13},
		[4]int{8, 5, 15, 6}, [4]int{5, 0, 9, 15}, [4]int{3, 15, 12, 0}, [4]int{15, 10, 5, 9},
		[4]int{13, 3, 6, 10}, [4]int{0, 9, 3, 4}, [4]int{14, 8, 0, 5}, [4]int{9, 6, 14, 3}}
        12	1	10	15	9	2	6	8	0	13	3	4	14	7	5	11
        10	15	4	2	7	12	9	5	6	1	13	14	0	11	3	8
        9	14	15	5	2	8	12	3	7	0	4	10	1	13	11	6
        4	3	2	12	9	5	15	10	11	14	1	7	6	0	8	13
	substitutionMatrix6 := [16][4]int{[4]int{}, [4]int{}, [4]int{}, [4]int{},
		[4]int{}, [4]int{}, [4]int{}, [4]int{},
		[4]int{}, [4]int{}, [4]int{}, [4]int{},
		[4]int{}, [4]int{}, [4]int{}, [4]int{}}
        0	4	11	2	14	15	0	8	13	3	12	9	7	5	10	6	1
        13	0	11	7	4	9	1	10	14	3	5	12	2	15	8	6
        1	4	11	13	12	3	7	14	10	15	6	8	0	5	9	2
        6	11	13	8	1	4	10	7	9	5	0	15	14	2	3	12
	substitutionMatrix7 := [16][4]int{[4]int{}, [4]int{}, [4]int{}, [4]int{},
		[4]int{}, [4]int{}, [4]int{}, [4]int{},
		[4]int{}, [4]int{}, [4]int{}, [4]int{},
		[4]int{}, [4]int{}, [4]int{}, [4]int{}}
        13	2	8	4	6	15	11	1	10	9	3	14	5	0	12	7
        1	15	13	8	10	3	7	4	12	5	6	11	0	14	9	2
        7	11	4	1	9	12	14	2	0	6	10	13	15	3	5	8
        2	1	14	7	4	10	8	13	15	12	9	0	3	5	6	11
	substitutionMatrix8 := [16][4]int{[4]int{}, [4]int{}, [4]int{}, [4]int{},
		[4]int{}, [4]int{}, [4]int{}, [4]int{},
		[4]int{}, [4]int{}, [4]int{}, [4]int{},
		[4]int{}, [4]int{}, [4]int{}, [4]int{}}

	fmt.Println("substitution Matrix 1")
	fmt.Println(substitutionMatrix1)

	var resultString string
	var substitutionResult int

	for i, dividedBlock := range dividedBlocks {
		// to determine the x you have to take the bits 2,3,4,5 on the concording dividedBlock and turn it into decimal
		xMatrix, _ := strconv.ParseUint(strings.Join(dividedBlocks[i][1:5], ""), 2, 0)
		// to determine the y you have to take the bits 1,6 on the concording dividedBlock and turn it into decimal
		yMatrix, _ := strconv.ParseUint(strings.Join([]string{dividedBlocks[i][0], dividedBlocks[i][5]}, ""), 2, 0)

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

		if substitutionResult <= 7 {
			resultString += "0" + strconv.FormatUint(uint64(substitutionResult), 2)
		} else {
			resultString += strconv.FormatUint(uint64(substitutionResult), 2)
		}

	}

	fmt.Println(resultString)
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
	subtitution(rightBlockExpanded)

	return []string{}
}
