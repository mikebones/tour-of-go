package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"regexp"
)
	

func findEmail(path string) {
	_, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		_, err = os.Create(path)
	}
	check(err)
	regex := regexp.MustCompile("[A-z0-9._]+@[A-z0-9.]+") // must compile handles err
	// read := make([]byte, 10)
	// byteRead, err := file.Read(read)
	
	// byteRead, err := file.Read(read)
	// for !errors.Is(err, io.EOF) {
	// 	byteRead.
	// 	fmt.Printf("%q\n", regex.FindAll(read[:byteRead], -1))
	
	// }
	// reader = bufio.NewReader(file)
	// io.Reader.Read()
	// regex.FindAll()
	entireFile, err := os.ReadFile(path)
	check(err)
	fmt.Printf("%q\n", regex.FindAll(entireFile, -1))
}

// most calls will be checked for errors
// helper function streamline error checks
func check(e error) {
	if e != nil {
		panic(e)
	}
} 

func main() {
	findEmail("test.txt")
}