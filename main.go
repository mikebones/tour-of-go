package main

import (
	"fmt"
	"math"
	"strconv"
)

var c, python, java bool

func main() {
	var sedan, truck, van = true, false, true
	fmt.Println(sedan, truck, van, java)
	fmt.Println(math.Pi)
	fmt.Println(add(6,7))
	fmt.Println(addNakedReturn(6,7))
	fmt.Println(swap("world!", "Hello"))
	split(155)
	splitReverse(155)
	strinngtomaprunes("stringtorunes")
}

// func addArray(x int, y []int) int {
// 	for _, y := range y {
// 		fmt.Println(y)
// 	}
// 	return x + y
// }

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {	
	return y, x
}

func addNakedReturn(x,y int) (sum int) {
	sum = x + y
	return
}

func split(x int) (split []rune) {
	stringconv := strconv.Itoa(x)
	chars := []rune(stringconv) // a slice of runes
	for i := 0 ; i < len(chars) ; i++ {
		char := string(chars[i])
		fmt.Println(char)
	}
	return chars
}

func strinngtomaprunes(breakthis string) (runes []rune) {
	fmt.Println("entering strinngtomaprunes")
	runes = []rune(breakthis)
	for i := 0; i < len(runes); i++ {
		fmt.Println(string(runes[i]))
		fmt.Printf("Index is %v", i)
	}
	return
}


// https://www.tutorialkart.com/golang-tutorial/golang-convert-string-into-array-of-characters/
func splitReverse (x int) (splitReverse []rune) {
	stringconv := strconv.Itoa(x) //returns the string
	chars := []rune(stringconv) // cast the string to a slice of runs
	// fmt.Println(len(chars)-1)
	// fmt.Print(string(chars[2]))
	for i := len(chars)-1  ; i >= 0 ; i-- { // len is number of elements // variables declared within for only have scope of for statement
		char := string(chars[i])
		fmt.Println(char)
	}
	var i int = 0
	for  i < len(chars) {
		
	}
	return splitReverse
}

func isvowel (char []rune) (isvowel bool) {
	var vowels = [...]string{"a", "e", "i", "o", "u"}
	fmt.Printf("vowels: %v\n", vowels)
	return
}


/*
pre -- vs post --
prefix value is operand on then returned
postfix value is return then operand
*/

/*
binary shift
>> (multiplied by 2)
<< (divided by two)
https://go.dev/ref/spec
https://stackoverflow.com/questions/5801008/go-and-operators
*/

/*
:= new variable instatiation and assignment with implicit type, cannot be used outside function : use var, cannot be used with constants
= variable assignment for variables that already exist
*/

/*
for initialization statement, executed first ; conditional statement, evaluated prior to everay iteration ; post statement, executed after each iteration
// loop stops once condition is FALSE
// initiatilization and post are optional
*/

/* top level key words
var, const, func
*/
