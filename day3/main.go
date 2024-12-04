package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file_name := "input.txt"
	file, err := os.Open(file_name)
	// var mult_array [][]int
	if err != nil {
		panic("open file failed :(")
	}
	scanner := bufio.NewReader(file)
	// pattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	// part1(scanner, pattern, mult_array)

	part2(scanner)

}

func part1(scanner *bufio.Reader, pattern *regexp.Regexp, mult_array [][]int) {

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

func part2(scanner *bufio.Reader) {
	sum := 0
	check := true

	do_or_dont := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d+,\d+\)`)
	digits := regexp.MustCompile(`\d+`)

	for {

		line, err := scanner.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}

		for _, match := range do_or_dont.FindAllString(line, -1) {
			fmt.Printf("Found match: %s\n", match)

			if match == "don't()" {
				check = false
				continue
			} else if match == "do()" {
				check = true
				continue
			}

			if check {
				matches := digits.FindAllString(match, -1)
				if len(matches) == 2 {
					num1, err1 := strconv.Atoi(matches[0])
					num2, err2 := strconv.Atoi(matches[1])
					if err1 != nil || err2 != nil {
						panic("Error parsing numbers in mul()")
					}
					product := num1 * num2
					fmt.Printf("Adding mul(%d,%d) = %d \n", num1, num2, product)
					sum += product
				}
			}
		}
	}

	fmt.Println(sum)

}
