package main

func main() {
	
	
}

// func isValid (s string) bool {
// 	// '(', ')', '{', '}', '[' and ']'
// 	// assumes inputs of only valid types above
// 	m := make(map[rune]int)
// 	var counter int
// 	r := []rune(s)
// 	for i := 1 ; i < len(r) ; i++ {
// 		// stack 
// 		// if we go negative then false
// 		if r[i] == '(' || r[i] == '{' || r[i] == '[' {
// 			m[r[i]] = m[r[i]] + 1 
// 			counter++
// 		} else {
// 			m[r[i]] = m[r[i]] - 1
// 			counter--
// 			if m[r[i]] < 0 {
// 				return false
// 			}
// 		}
// 	}
// 	return counter==0
// }


func isValid (s string) bool {
	m := map[rune]rune  {
		'}' : '{',
		')' : '(',
		']' : '[',
	}
	stack := make([]rune, 0)
	for _, c := range s {
		// switch c == '{' || c == '(' || c == '[' {
		// 	case true
		// }
		switch c {
			case '(', '{', '[' :
				stack = append(stack, c)			
			case ')', '}', ']' : 
				if len(stack) == 0 || stack[len(stack) - 1] != m[c] {
					return false
				}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
// o (n) time because we only go through the entire list once
// o (n) space because the stack can be the entire size of the stack
