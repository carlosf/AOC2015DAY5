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

func strRemoveAt(s string, index, length int) string {
	return s[:index] + s[index+length:]
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

func isLineNice2(s string) bool {
	var currChar rune
	var doubleChar bool = false
	var cleanStr string = ""
	var stringtoCheck strings.Builder
	var stringDup bool = false

	//p(s)
	key := 0
	for key, currChar = range s {
		// check for repeated characters
		if key > 0 {
			cleanStr = strRemoveAt(s, key-1, 2)
			stringtoCheck.WriteString(string(s[key-1]))
			stringtoCheck.WriteString(string(currChar))
			//p(stringtoCheck.String())
			if strings.Contains(cleanStr, stringtoCheck.String()) {
				doubleChar = true
				p("found double: " + stringtoCheck.String())
			}
			stringtoCheck.Reset()
		}
		// contains at least one character which repeats with exactly one letter between them
		if key >= 2 && string(s[key-2]) == string(currChar) {
			stringDup = true
			p("repeat: " + string(s[key-2]) + string(s[key-1]) + string(currChar))
		}
		//check to see if there is a 'ab', 'cd', 'pq', or 'xy'

	}
	if doubleChar && stringDup {
		p(s)
	}
	return doubleChar && stringDup
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
		if isLineNice2(line) {
			numNice++
		}
	}
	fmt.Println("Number of nice strings: ", numNice)
}
