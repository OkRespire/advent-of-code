package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("testInput.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var arr []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		arr = strings.Split(line, "")
	}

	fmt.Println(arr)

	// part1(arr)
	part2(arr)

}

func part1(arr []string) {

	new_arr := createSpace(arr)

	fmt.Println(new_arr)
	end_of_list := len(new_arr) - 1
	sum := 0

	for i, v := range new_arr {
		if end_of_list == i {
			break
		}

		for {
			if new_arr[end_of_list] == -1 {
				end_of_list--
				continue
			}
			break
		}
		if v == -1 {
			new_arr[i], new_arr[end_of_list] = new_arr[end_of_list], new_arr[i]
			end_of_list--
		} else {
			continue
		}
	}

	for i, v := range new_arr {
		if v == -1 {
			continue
		}
		sum += v * i
	}

	fmt.Println(sum)

}

func part2(arr []string) {
	compacted_arr := createSpace(arr)
	end_of_list := len(compacted_arr) - 1
	// var prev_val int

	map_of_vals := makeMap(compacted_arr)
	fmt.Println(map_of_vals)
	// sum := 0

	fmt.Println(compacted_arr)

	for i, v := range compacted_arr {
		for {
			if compacted_arr[end_of_list] == -1 {
				end_of_list--
				continue
			}
			break
		}
		if i == 0 {
			prev_val = v
		}
		if v == -1 {
			//look at every value to see if the previous value

			//if v == -1, and there are no more values that match the previous value, then have a unique value (and its files taking up the same space)

		}
	}

}

func createSpace(arr []string) []int {

	current_val := 0
	var new_arr []int

	//1st is id, 2nd is empty space
	for i, v := range arr {
		if i%2 == 0 {
			value_range, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			for j := 0; j < value_range; j++ {
				new_arr = append(new_arr, current_val)
			}
			current_val++
		} else {
			current_val, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			for j := 0; j < current_val; j++ {
				new_arr = append(new_arr, -1)
			}
		}
	}
	return new_arr
}

func makeMap(arr []int) map[int]int {
	m := make(map[int]int)
	for _, v := range arr {
		if v != -1 {
			not_minus_one_val := v
			continue
		}
		if v == -1 {
			m[not_minus_one_val]++
		}
	}
	return m
}

//Make a map of the consecutive -1 between two files
//make a temp end val for swaps
