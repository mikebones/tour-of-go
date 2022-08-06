package main

import (
	"bufio"
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
	scanEmails("test.txt")
}


func findEmails(path string) {
	_, err := os.Open(path)
	check(err)
	regex := regexp.MustCompile("[A-z0-9._]+@[A-z0-9.]+") // must compile handles err
	entireFile, err := os.ReadFile(path)
	check(err)
	fmt.Printf("%q\n", regex.FindAll(entireFile, -1))
}
// scanner to do new lines?

func scanEmails(path string) {
	//open file
	//remember to close file
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	// compile regex
	regexEmail := regexp.MustCompile("[A-z0-9._]+@[A-z0-9._]")
	scanner := bufio.NewScanner(file)
//	scanner.Split(bufio.ScanRunes)
	go regexEmailScan(*scanner, *regexEmail)
	for scanner.Scan() {
		if regexEmail.FindString(scanner.Text()) != "" {
			fmt.Printf("%s\n", scanner.Text())
		}	
	}	
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}

func regexEmailScan(scanner bufio.Scanner, regexEmail regexp.Regexp) {
	for scanner.Scan() {
		if regexEmail.FindString(scanner.Text()) != "" {
			fmt.Printf("%s\n", scanner.Text())
		}	
	}

}
// it could be faster to load the entire file into ram and run regex on it
// if the file is huge we could run it line by line but this could be an io issue with disk
// if we try to break this up into smaller chunks then do we have an issue of not catching our regex
