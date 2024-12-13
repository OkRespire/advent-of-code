package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache map[int]int = make(map[int]int)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var line_split []string

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()

		line_split = strings.Split(line, " ")
	}

	// for i := 0; i < 25; i++ {
	// 	line_split = blink(line_split)
	// }
	//
	// fmt.Println(len(line_split))

	var int_stones []int

	for _, v := range line_split {
		v_int, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		int_stones = append(int_stones, v_int)
	}
	sum := 0
	end := 75

	sum = blink2(int_stones, end)

	fmt.Println(sum)

}

func blink(stones []string) []string {
	var new_arr []string
	for _, v := range stones {
		curr_v, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		if curr_v == 0 {
			new_arr = append(new_arr, "1")

		} else if len(v)%2 == 0 {
			left_side_int, _ := strconv.Atoi(v[:len(v)/2])
			right_side_int, _ := strconv.Atoi(v[len(v)/2:])

			left_side := strconv.Itoa(left_side_int)
			right_side := strconv.Itoa(right_side_int)

			new_arr = append(new_arr, left_side)
			new_arr = append(new_arr, right_side)

		} else {
			curr_v = curr_v * 2024
			curr_v_str := strconv.Itoa(curr_v)
			new_arr = append(new_arr, curr_v_str)
		}

	}

	return new_arr

}

func blink2(stones []int, end int) int {
	stones_map := make(map[int]int)

	for _, v := range stones {
		stones_map[v]++
	}
	for i := 0; i < end; i++ {
		new_stones := make(map[int]int)
		for stone, count := range stones_map {

			if stone == 0 {
				new_stones[1] += count

			} else if len(strconv.Itoa(stone))%2 == 0 {
				mid := len(strconv.Itoa(stone)) / 2
				str_stone, _ := strconv.Atoi(strconv.Itoa(stone)[:mid])
				str_stone2, _ := strconv.Atoi(strconv.Itoa(stone)[mid:])
				new_stones[str_stone] += count
				new_stones[str_stone2] += count

			} else {
				new_stones[stone*2024] += count
			}
		}

		stones_map = new_stones

	}
	sum := 0

	for _, v := range stones_map {
		sum += v
	}
	return sum

}
