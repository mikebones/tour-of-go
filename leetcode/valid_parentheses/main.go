package main

func main() {
	
	
}

func isValid (s string) bool {
	// '(', ')', '{', '}', '[' and ']'
	// assumes inputs of only valid types above
	m := make(map[rune]int)
	var counter int
	r := []rune(s)
	for i := 1 ; i < len(r) ; i++ {
		// stack 
		// if we go negative then false
		if r[i] == '(' || r[i] == '{' || r[i] == '[' {
			m[r[i]] = m[r[i]] + 1 
			counter++
		} else {
			m[r[i]] = m[r[i]] - 1
			counter--
			if m[r[i]] < 0 {
				return false
			}
		}
	}
	return counter==0
}
