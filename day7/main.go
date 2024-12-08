package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

func main() {
	file, err := os.Open("input.txt")

	var arr_nums [][]int

	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	for s.Scan() {
		line := s.Text()

		line = strings.TrimSpace(line)
		line = strings.Replace(line, ":", "", -1)

		line_split := strings.Split(line, " ")

		var temp_arr []int

		for _, v := range line_split {
			strToInt, err := strconv.Atoi(v)

			if err != nil {
				panic(err)
			}
			temp_arr = append(temp_arr, strToInt)

		}

		arr_nums = append(arr_nums, temp_arr)
	}
	part1(arr_nums)
	part2(arr_nums)
}

func part1(arr [][]int) {
	sum := 0
	ops := []string{"+", "*"}

	for _, v := range arr {
		opsCombinations := getCartesianProduct(ops, len(v)-1)
		ans := v[0]
		v = v[1:]

		for _, ops := range opsCombinations {
			temp_ans := 0

			for i, op := range ops {
				if op == "+" {
					temp_ans += v[i]
				} else if op == "*" {
					if temp_ans == 0 {
						temp_ans = 1
					}
					temp_ans *= v[i]
				}
			}
			if temp_ans == ans {
				fmt.Printf("Actual Answer %d \nTemp Answer %d\n", ans, temp_ans)
				sum += ans
				break
			}

		}

	}
	fmt.Println(sum)

}

func part2(arr [][]int) {
	sum := 0
	ops := []string{"+", "*", "|"}

	for _, v := range arr {
		opsCombinations := getCartesianProduct(ops, len(v)-1)
		ans := v[0]
		v = v[1:]

		for _, ops := range opsCombinations {
			temp_ans := 0

			for i, op := range ops {
				if op == "+" {
					temp_ans += v[i]
				} else if op == "*" {
					if temp_ans == 0 {
						temp_ans = 1
					}
					temp_ans *= v[i]
				} else if op == "|" {
					temp_ansStr := strconv.Itoa(temp_ans)
					vStr := strconv.Itoa(v[i])
					temp_ans, _ = strconv.Atoi(temp_ansStr + vStr)

				}
			}
			if temp_ans == ans {
				fmt.Printf("Actual Answer %d \nTemp Answer %d\n", ans, temp_ans)
				sum += ans
				break
			}

		}

	}
	fmt.Println(sum)
}

func getCartesianProduct(ops []string, r int) [][]string {

	opsToInt := map[int]string{0: "+", 1: "*", 2: "|"}
	dim := make([]int, r)

	//lengths of all the combinations for the operators
	for i := 0; i < r; i++ {
		dim[i] = len(ops)
	}
	var combos [][]string

	cg := combin.NewCartesianGenerator(dim)
	for cg.Next() {
		combo := cg.Product(nil)
		var strCombos []string

		for _, v := range combo {
			value, _ := opsToInt[v]
			strCombos = append(strCombos, value)
		}
		combos = append(combos, strCombos)

	}

	return combos
}
