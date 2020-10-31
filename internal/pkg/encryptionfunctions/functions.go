package encryptionfunctions

// Expansion is the implementation of the expansive permutation in the DES algorithm
// the function take a slice of 32 strings and return a slice of 48 after the expansion
func Expansion(s []string) []string {

	return []string{s[31], s[0], s[1], s[2], s[3], s[4], s[3], s[4],
		s[5], s[6], s[7], s[8], s[7], s[8], s[9], s[10],
		s[11], s[12], s[11], s[12], s[13], s[14], s[15], s[16],
		s[15], s[16], s[17], s[18], s[19], s[20], s[19], s[20],
		s[21], s[22], s[23], s[24], s[23], s[24], s[25], s[26],
		s[27], s[28], s[27], s[28], s[29], s[30], s[31], s[0]}
}
