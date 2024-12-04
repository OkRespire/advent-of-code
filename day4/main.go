package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	X = "X"
	M = "M"
	A = "A"
	S = "S"
)

type coords struct {
	x int
	y int
}
type Part1Params struct {
	IsX bool
	IsM bool
	IsA bool
	IsS bool
}

func main() {

	file, err := os.Open("testInput.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewReader(file)
	var cw_arr [][]string

	for {
		line, err := s.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		line = strings.TrimSpace(line)
		linesplit := strings.Split(line, "")

		cw_arr = append(cw_arr, linesplit)

	}

	params := Part1Params{IsX: false, IsM: false, IsA: false, IsS: false}

	// part1(cw_arr, params)
	part1Simple(cw_arr, params)

	for _, v := range cw_arr {
		fmt.Println(v)
		break
	}
}

func part1Simple(arr [][]string, params Part1Params) {
	str := ""

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[0][i])

		if params.IsA && params.IsM && params.IsX && params.IsS {
			fmt.Println(str)
			str = ""
			params = Part1Params{IsX: false, IsM: false, IsA: false, IsS: false}
		}

		if arr[0][i] == X {
			if params.IsX {
				params = Part1Params{IsX: true, IsM: false, IsA: false, IsS: false}
				fmt.Println("here")
				str = "X"
				continue
			}
			str += arr[0][i]
			params.IsX = true
		} else if arr[0][i] == M {
			if params.IsM {
				params = Part1Params{IsX: false, IsM: true, IsA: false, IsS: false}
				fmt.Println("here")
				str = "M"
				continue
			}

			str += arr[0][i]
			params.IsM = true
		} else if arr[0][i] == A {
			if params.IsA {
				params = Part1Params{IsX: false, IsM: false, IsA: true, IsS: false}
				fmt.Println("here")
				str = "A"

				continue
			}

			str += arr[0][i]

			params.IsA = true
		} else if arr[0][i] == S {
			if params.IsS {
				params = Part1Params{IsX: false, IsM: false, IsA: false, IsS: true}
				fmt.Println("here")
				str = "S"
				continue
			}

			str += arr[0][i]
			params.IsS = true
		}

	}
}

// func part1(array [][]rune, params Part1Params) {
//
// 	for i := 0; i < len(array); i++ {
// 		for j := 0; j < len(array[i]); j++ {
// 			newCoords := coords{i, j}
// 			if array[i][j] == X {
// 				params = Part1Params{IsX: true}
//
// 				// res := lookAtNeighbours(array, newCoords, params)
// 			}
// 			if array[i][j] == M {
//
// 				params = Part1Params{IsM: true}
// 				newCoords := coords{i, j}
//
// 				// res := lookAtNeighbours(array, newCoords, params)
// 			}
// 			if array[i][j] == A {
//
// 				params = Part1Params{IsA: true}
//
// 				// res := lookAtNeighbours(array, newCoords, params)
// 			}
// 			if array[i][j] == S {
//
// 				newCoords := coords{i, j}
// 				params := Part1Params{IsS: true}
// 				// res := lookAtNeighbours(array, newCoords, params)
// 			}
//
// 		}
// 	}
//
// }

// Look at neighbours, all directions
// func lookAtNeighbours(arr [][]rune, coords coords, params Part1Params) int {
//
// 	if params.IsA && params.IsM && params.IsS && params.IsX {
// 		return 1
// 	}
//
// 	//TODO: Write logic for if its at edges of x value
// 	//look at x neighbours
// 	if coords.x == 0 {
// 		// look at right neighbour
//
// 	} else if coords.x == len(arr) {
// 		// look at left neigbhour
// 	}
//
// }
