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

}

// Horizontal
func part1Simple(arr [][]string, params Part1Params) {
	str := ""
	count := 0

	for _, v := range arr {
		for j := 0; j < len(v); j++ {

			// fmt.Println(v[j])

			if params.IsA && params.IsM && params.IsX && params.IsS {
				fmt.Println(str)
				str = ""
				params = Part1Params{IsX: false, IsM: false, IsA: false, IsS: false}
				fmt.Println(count)
				count++
			}

			if v[j] == X {
				if params.IsX {
					params = Part1Params{IsX: true, IsM: false, IsA: false, IsS: false}
					// fmt.Println("here")
					str = "X"
					continue
				}
				str += v[j]
				params.IsX = true
			} else if v[j] == M {
				if params.IsM {
					params = Part1Params{IsX: false, IsM: true, IsA: false, IsS: false}
					// fmt.Println("here")
					str = "M"
					continue
				}

				str += v[j]
				params.IsM = true
			} else if v[j] == A {
				if params.IsA {

					params = Part1Params{IsX: false, IsM: false, IsA: true, IsS: false}
					// fmt.Println("here")
					str = "A"

					continue
				}

				str += v[j]

				params.IsA = true
			} else if v[j] == S {
				if params.IsS {
					params = Part1Params{IsX: false, IsM: false, IsA: false, IsS: true}
					// fmt.Println("here")
					str = "S"
					continue
				}

				str += v[j]
				params.IsS = true
			}

		}
	}

	countO := vertical(arr, params)

	fmt.Println(count)
	count = count + countO
	fmt.Println(count)
	fmt.Println(countO)
}

func vertical(arr [][]string, params Part1Params) int {
	str := ""
	count := 0

	for i := range arr {
		for j := 0; j < len(arr); j++ {

			// fmt.Println(arr[j][i])

			if params.IsA && params.IsM && params.IsX && params.IsS {
				fmt.Println(str)
				str = ""
				params = Part1Params{IsX: false, IsM: false, IsA: false, IsS: false}
				count++
			}

			if arr[j][i] == X {
				if params.IsX {
					params = Part1Params{IsX: true, IsM: false, IsA: false, IsS: false}
					// fmt.Println("here")
					str = "X"
					continue
				}
				str += arr[j][i]
				params.IsX = true
			} else if arr[j][i] == M {
				if params.IsM {
					params = Part1Params{IsX: false, IsM: true, IsA: false, IsS: false}
					// fmt.Println("here")
					str = "M"
					continue
				}

				str += arr[j][i]
				params.IsM = true
			} else if arr[j][i] == A {
				if params.IsA {
					params = Part1Params{IsX: false, IsM: false, IsA: true, IsS: false}
					// fmt.Println("here")
					str = "A"

					continue
				}

				str += arr[j][i]

				params.IsA = true
			} else if arr[j][i] == S {
				if params.IsS {
					params = Part1Params{IsX: false, IsM: false, IsA: false, IsS: true}
					// fmt.Println("here")
					str = "S"
					continue
				}

				str += arr[j][i]
				params.IsS = true
			}

		}
	}

	fmt.Println("Count ", count)

	return count
}

// func diagonals(arr [][]string, params Part1Params) int {
//
// 	str := ""
// 	count := 0
//
// 	for _, v := range arr {
// 		for j := 0; j < len(v); j++ {
//
// 			fmt.Println(v[j])
//
// 			if params.IsA && params.IsM && params.IsX && params.IsS {
// 				fmt.Println(str)
// 				str = ""
// 				params = Part1Params{IsX: false, IsM: false, IsA: false, IsS: false}
// 				count++
// 			}
//
// 			if v[j] == X {
// 				if params.IsX {
// 					params = Part1Params{IsX: true, IsM: false, IsA: false, IsS: false}
// 					fmt.Println("here")
// 					str = "X"
// 					continue
// 				}
// 				str += v[j]
// 				params.IsX = true
// 			} else if v[j] == M {
// 				if params.IsM {
// 					params = Part1Params{IsX: false, IsM: true, IsA: false, IsS: false}
// 					fmt.Println("here")
// 					str = "M"
// 					continue
// 				}
//
// 				str += v[j]
// 				params.IsM = true
// 			} else if v[j] == A {
// 				if params.IsA {
// 					params = Part1Params{IsX: false, IsM: false, IsA: true, IsS: false}
// 					fmt.Println("here")
// 					str = "A"
//
// 					continue
// 				}
//
// 				str += v[j]
//
// 				params.IsA = true
// 			} else if v[j] == S {
// 				if params.IsS {
// 					params = Part1Params{IsX: false, IsM: false, IsA: false, IsS: true}
// 					fmt.Println("here")
// 					str = "S"
// 					continue
// 				}
//
// 				str += v[j]
// 				params.IsS = true
// 			}
//
// 		}
//     return count
// }
//
