package main

import (
	"unicode"
)


func main() {
	// convert to lower
	// remove alpha numeric
	// unicode.IsLetter

}

func isPalindrome(s string) bool {
	sliceOfRunes := []rune(s)
	// for idx, val := range sliceOfRunes {
	// for i := len(sliceOfRunes) -1 ; i >= 0 ; i ++ {
	// 	// if unicode.IsLetter(val) {
	// 	// 	unicode.ToLower(val)
	// 	// }
	// 	// if unicode.IsLetter(sliceOfRunes[i]) {
	// 	// 	if unicode.IsLett(sliceOfRunes(indexToCheck)) {
	// 	// 		if unicode.ToLower(sliceOfRunes[i]) == 
	// 	// 	} else {
	// 	// 		indexToCheck--
	// 	// 	}
			
	// 	// }
	// 	if unicode.IsLetter(sliceOfRunes[i]) {
	// 		mapOfLetters[]
	// 	}
	// }
	n := (len(sliceOfRunes)-1)
	i := 0
	if len(sliceOfRunes) == 1 {
		return true
	}
	//regExp := "[A-Za-z0-9]"
	var palindrome bool = true
	for palindrome == true {
		if unicode.IsLetter(sliceOfRunes[i]) && unicode.IsLetter(sliceOfRunes[n]) {
			if unicode.ToLower(sliceOfRunes[i]) != unicode.ToLower(sliceOfRunes[n]) {
				palindrome = false
				return palindrome
			} else if unicode.ToLower(sliceOfRunes[i]) == unicode.ToLower(sliceOfRunes[n]) {
				if i == n && (unicode.ToLower(sliceOfRunes[i]) == unicode.ToLower(sliceOfRunes[n])) {
					return palindrome
				}
				if i + 1 == n - 1 && (unicode.ToLower(sliceOfRunes[i]) == unicode.ToLower(sliceOfRunes[n])){
					palindrome = true
					return palindrome
				}
				i++
				n--
			}
		} else if i == n && !(unicode.ToLower(sliceOfRunes[i]) == unicode.ToLower(sliceOfRunes[n])) {
			palindrome = false
			return palindrome
		} else if !(unicode.IsLetter(sliceOfRunes[i])) {
			i++
		} else if !(unicode.IsLetter(sliceOfRunes[n])) {
			n--
		} else if i + 1 == n - 1 && !(unicode.ToLower(sliceOfRunes[i]) == unicode.ToLower(sliceOfRunes[n])){
			palindrome = false
			return palindrome
		}
	}
	return palindrome
}

/*
if i character is letter and n character is leter and they equal
palindrom true
if i character is not letter 
i ++ 
*/

import "fmt"

func isPalindrome(s string) bool {
	sliceOfRunes := []rune(s)
	// for idx, val := range sliceOfRunes {
	// for i := len(sliceOfRunes) -1 ; i >= 0 ; i ++ {
	// 	// if unicode.IsLetter(val) {
	// 	// 	unicode.ToLower(val)
	// 	// }
	// 	// if unicode.IsLetter(sliceOfRunes[i]) {
	// 	// 	if unicode.IsLett(sliceOfRunes(indexToCheck)) {
	// 	// 		if unicode.ToLower(sliceOfRunes[i]) == 
	// 	// 	} else {
	// 	// 		indexToCheck--
	// 	// 	}
			
	// 	// }
	// 	if unicode.IsLetter(sliceOfRunes[i]) {
	// 		mapOfLetters[]
	// 	}
	// }
	n := (len(sliceOfRunes)-1)
	i := 0
	//regExp := "[A-Za-z0-9]"
	var palindrome bool = true
	for palindrome == true {
        fmt.Println(len(sliceOfRunes))
        fmt.Println(n)
        fmt.Println(i)
		if unicode.IsLetter(sliceOfRunes[i]) && unicode.IsLetter(sliceOfRunes[n]) {
            fmt.Println("is letter")
			if unicode.ToLower(sliceOfRunes[i]) != unicode.ToLower(sliceOfRunes[n]) {
				palindrome = false
				return palindrome
			} else if unicode.ToLower(sliceOfRunes[i]) == unicode.ToLower(sliceOfRunes[n]) {
                fmt.Println("is same")
				i++
				n--
			}
		} else if i == n && (unicode.ToLower(sliceOfRunes[i]) == unicode.ToLower(sliceOfRunes[n])) {
			return palindrome
		} else if i == n && !(unicode.ToLower(sliceOfRunes[i]) == unicode.ToLower(sliceOfRunes[n])) {
			palindrome = false
			return palindrome
		} else if !(unicode.IsLetter(sliceOfRunes[i])) {
			i++
		} else if !(unicode.IsLetter(sliceOfRunes[n])) {
			n--
		} else if i + 1 == n - 1 && (unicode.ToLower(sliceOfRunes[i]) == unicode.ToLower(sliceOfRunes[n])){
			palindrome = true
			return palindrome
		} else if i + 1 == n - 1 && !(unicode.ToLower(sliceOfRunes[i]) == unicode.ToLower(sliceOfRunes[n])){
			palindrome = false
			return palindrome
		}
	}
	return palindrome
}
