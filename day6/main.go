package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coords struct {
	x int
	y int
}

func main() {

	file, err := os.Open("testInput.txt")
	var start_point coords

	var path [][]string

	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()

		line = strings.TrimSpace(line)

		line_split := strings.Split(line, "")

		path = append(path, line_split)
	}
	fmt.Println(path)

	for i, v := range path {
		for j, val := range v {
			if val == "^" {
				start_point = coords{i, j}
			}
		}
	}

	fmt.Println(start_point)

	part1(start_point, path)
}

func part1(start_point coords, path [][]string) {
	active_dir := 0 // 0 = up, 1 = right, 2 = down, 3 = left
	new_paths := 0

	direction := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} //up, right, down, left

	var visited_map map[coords]bool
	visited_map = make(map[coords]bool)

	fmt.Println("height", len(path))

	fmt.Println("width", len(path[0]))

	for {
		if (start_point.x <= 0 || start_point.x >= len(path[0])) || (start_point.y <= 0 || start_point.y >= len(path)) {
			fmt.Println("final coords", start_point)
			break
		}
		if _, ok := visited_map[start_point]; !ok {
			fmt.Println("Hi")
			visited_map[start_point] = true
			new_paths++
		}

		if path[start_point.x][start_point.y] == "#" {
			active_dir += 1
			if active_dir > 3 {
				active_dir = 0
			}

		}

		new_x := start_point.x + direction[active_dir][1]
		fmt.Println("new x", new_x)
		new_y := start_point.y + direction[active_dir][0]
		fmt.Println("new y", new_y)

		fmt.Println(path[new_x][new_y])

		fmt.Println(new_x, new_y)

		if (new_x <= 0 || new_x >= len(path[0])) || (new_y <= 0 || new_y >= len(path)) {
			fmt.Println("Hi")
			break
		}
		start_point.x = new_x
		start_point.y = new_y

	}

	fmt.Println(new_paths)
}
