package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewReader(file)
	flag := false

	m := make(map[int][]int)

	var csv_list [][]int
	var temp_arr []int

	for {
		line, err := s.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if line == "\n" {
			flag = true
			continue

		}

		if flag {
			if line == "\n" {
				continue
			}
			line = strings.TrimSpace(line)
			split_line := strings.Split(line, ",")
			for i := range split_line {
				int_line, _ := strconv.Atoi(split_line[i])
				temp_arr = append(temp_arr, int_line)
			}

			csv_list = append(csv_list, temp_arr)
			temp_arr = nil
		} else {
			//adds the definitions into a map
			line = strings.TrimSpace(line)
			line_split := strings.Split(line, "|")

			key, err := strconv.Atoi(line_split[0])
			if err != nil {
				panic(err)
			}

			val, err := strconv.Atoi(line_split[1])
			if err != nil {
				panic(err)
			}

			if value, ok := m[key]; ok {
				value = append(value, val)
				m[key] = value
			} else {
				m[key] = []int{val}
			}

		}

	}

	// fmt.Println(m)

	// for _, v := range csv_list {
	// 	fmt.Println(v)
	// }

	part1(m, csv_list)
	part2(m, csv_list)

}

func part1(m map[int][]int, arr [][]int) {
	key := arr[0][0]
	sum := 0
	new_arr, ok := m[key]

	var mid_val int
	flag := true

	for _, v := range arr {
		for j, k := range v {
			if j == 0 {
				key = k
				continue
			}

			mid_val = middleVal(v)

			new_arr, ok = m[key]

			if !ok && flag && j == len(v)-1 {
				flag = true
			}

			if slices.Contains(new_arr, k) {
				key = k
				flag = true

			} else {
				flag = false
				break
			}

		}

		if flag {
			// fmt.Println(mid_val)
			sum += mid_val
		}
	}

	fmt.Println(sum)
}

func part2(m map[int][]int, arr [][]int) {
	sum := 0
	for _, v := range arr {

		if isSorted(v, m) {
			continue
		}

		temp_array := bubbleSort(v, m)

		// fmt.Println(temp_array)
		mid_val := middleVal(temp_array)

		sum += mid_val

	}

	fmt.Println(sum)

}

func middleVal(arr []int) int {
	if len(arr)%2 == 0 {
		return arr[len(arr)/2-1]
	}
	return arr[len(arr)/2]
}

func bubbleSort(v []int, m map[int][]int) []int {
	sorted := slices.Clone(v)

	for i := 0; i < len(v)-1; i++ {
		for j := 0; j < len(v)-1; j++ {

			if !slices.Contains(m[sorted[j]], sorted[j+1]) {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}

		}

	}

	return sorted
}

func isSorted(arr []int, m map[int][]int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if !slices.Contains(m[arr[i]], arr[i+1]) {
			return false
		}
	}
	return true
}
