package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	x int
	y int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var arr [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		line = strings.TrimSpace(line)
		line_split := strings.Split(line, "")

		var temp_arr []int

		for _, v := range line_split {
			strToInt, err := strconv.Atoi(v)

			if err != nil {
				panic(err)
			}
			temp_arr = append(temp_arr, strToInt)
		}
		arr = append(arr, temp_arr)
	}
	fmt.Println(arr)
	// part1(arr)
	part2(arr)
}

func part1(arr [][]int) {
	sum := 0
	for i, v := range arr {
		for j := range v {
			sum += score(coords{j, i}, arr)
		}
	}
	fmt.Println(sum)
}

func part2(arr [][]int) {
	sum := 0
	for i, v := range arr {
		for j := range v {
			if arr[i][j] == 0 {
				sum += FindConsecutiveNums(arr, coords{j, i}, map[coords]int{})
			}
		}
	}
	fmt.Println(sum)
}

func FindConsecutiveNums(arr [][]int, point coords, visited map[coords]int) int {

	if v, ok := visited[point]; ok {
		return v
	}
	if arr[point.y][point.x] == 9 {
		visited[point] = 1
		return 1
	}
	directions := []coords{{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1}}

	ans := 0

	for _, v := range directions {
		new_point := coords{point.x + v.x, point.y + v.y}

		if outOfBounds(new_point, arr) {
			continue
		}

		if arr[new_point.y][new_point.x] == arr[point.y][point.x]+1 {
			ans += FindConsecutiveNums(arr, new_point, visited)
		}
	}

	visited[point] = ans

	return ans

}

// adapted solution from: https://www.youtube.com/watch?v=MmT5xgXnbcc
// Thank you William Y. Feng!!!
func score(point coords, arr [][]int) int {
	if arr[point.y][point.x] != 0 {
		return 0
	}
	ans := 0

	directions := []coords{{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1}}

	stack := []coords{point}

	visited := map[coords]bool{}

	for len(stack) > 0 {
		cur_point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if visited[cur_point] {
			continue
		}

		visited[cur_point] = true

		cur_val := arr[cur_point.y][cur_point.x]

		if cur_val == 9 {
			ans++
			continue
		}

		for _, v := range directions {
			new_point := coords{cur_point.x + v.x, cur_point.y + v.y}
			if outOfBounds(new_point, arr) {

				continue
			}
			neigh_val := arr[new_point.y][new_point.x]
			if neigh_val == cur_val+1 {
				stack = append(stack, new_point)
			}
		}
	}

	return ans
}

func outOfBounds(point coords, arr [][]int) bool {
	return point.x < 0 || point.x >= len(arr[0]) || point.y < 0 || point.y >= len(arr)
}
