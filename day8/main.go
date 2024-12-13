package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type coords struct {
	x int
	y int
}

func main() {
	file, err := os.Open("testInput.txt")
	var arr [][]string

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()
		line = strings.TrimSpace(line)
		line_split := strings.Split(line, "")

		fmt.Println(line_split)
		arr = append(arr, line_split)

	}

}

func part1(arr [][]string) {

	var map_of_antinodes map[string][]coords

	for _, v := range arr {
		for _, n := range v {
			if isAlphaNumeric(n) {

			}

		}

	}

}

func isAlphaNumeric(a string) bool {

	return unicode.IsLetter(rune(a[0])) || unicode.IsNumber(rune(a[0]))
}
