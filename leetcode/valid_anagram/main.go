package anagram

func main (s, t string) bool {
	// because we have to use all the same letters the length must be the same
	if len(s) != len(t) {
		return false
	}
	srune := []rune(s)
	trune := []rune(t)
	for i :=0 ; i < len(s) ; i++ {
	}
	return true
}

// if t is an anagram of s
// otherwise false
