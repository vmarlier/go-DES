package encryptionfunctions

// LeftShift function provide a left shift of i case in a string slice
func LeftShift(s []string, i int) []string {

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
