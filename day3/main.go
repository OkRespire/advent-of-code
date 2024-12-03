package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file_name := "input.txt"
	file, err := os.Open(file_name)
	var mult_array [][]int
	if err != nil {
		panic("something went wrong")
	}
	scanner := bufio.NewReader(file)
	pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for {
		line, err := scanner.ReadString('\n')
		if err != nil {
			break
		}

		submatches := pattern.FindAllStringSubmatch(line, -1)

		for _, v := range submatches {
			fmt.Println(v[0])
			v1, err := strconv.Atoi(v[1])
			if err != nil {
				panic("something went wrong")
			}
			v2, err := strconv.Atoi(v[2])
			if err != nil {
				panic("something went wrong")
			}
			new_arr := []int{v1, v2}
			mult_array = append(mult_array, new_arr)
		}

	}

	sum := 0

	for _, v := range mult_array {
		sum += v[0] * v[1]
	}

	fmt.Println(sum)
}
