package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func p(s string) {
	fmt.Println(s)
}

func isVowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
		return true
	}
	return false
}

func isLineNice(s string) bool {
	var currChar rune
	var numVowels int = 0
	var doubleChar bool = false
	var sb strings.Builder
	var forbiddenStrings bool = false

	//p(s)
	key := 0
	for key, currChar = range s {
		// count the number of vowels
		if isVowel(currChar) {
			numVowels++
			//fmt.Println("numVowels: ", currChar)
		}
		// check for repeated characters
		if key > 0 {
			if currChar == rune(s[key-1]) {
				doubleChar = true
				//fmt.Println("doubleChar: ", currChar)
			}
			sb.WriteString(string(s[key-1]))
			sb.WriteString(string(currChar))

			//p(sb.String())
			if sb.String() == "ab" || sb.String() == "cd" || sb.String() == "pq" || sb.String() == "xy" {
				forbiddenStrings = true
				//fmt.Println("forbidden chars: ", sb.String())
			}
			sb.Reset()
		}
		//check to see if there is a 'ab', 'cd', 'pq', or 'xy'

	}
	return (numVowels >= 3) && doubleChar && !forbiddenStrings
}

func main() {
	var numNice int = 0
	//Open input.txt file
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if scanner.Err() != nil {
			break
		}
		if isLineNice(line) {
			numNice++
		}
	}
	fmt.Println("Number of nice strings: ", numNice)
}
