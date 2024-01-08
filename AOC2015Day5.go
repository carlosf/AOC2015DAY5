package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func p(s string) {
	fmt.Println(s)
}

func isLineNice(s string) bool {

	return true
}

func main() {

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
			p("Nice")
		} else {
			p("Naughty")
		}
	}

}
