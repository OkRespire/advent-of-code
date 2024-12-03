package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	textFile := "input.txt"
	var l1, l2 []int
	file, err := os.Open(textFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		s := string(line)
		split := strings.Split(s, "   ")
		for i, v := range split {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println(err)
				return
			}
			if i%2 == 0 {
				l1 = append(l1, vInt)
			} else {
				l2 = append(l2, vInt)
			}
		}
	}

	// sort.Ints(l1) for task1
	// sort.Ints(l2)

	fmt.Println("l1:", l1)
	fmt.Println("l2:", l2)
	sum := 0

	for i := 0; i < len(l1); i++ {
		if i < len(l2) {
			float := math.Abs(float64(l1[i] - l2[i]))
			sum += int(float)
			fmt.Println("Absolute difference at index", i, "is:", float)
		} else {
			break
		}
	}

	fmt.Println(sum)

	a := similar(l1, l2)
	fmt.Println(" ")
	fmt.Println(a)
}

func similar(l1, l2 []int) int {
	myMap := make(map[int]int)

	for _, v := range l1 {
		myMap[v] = 0
	}

	for _, v := range l2 {
		if _, ok := myMap[v]; ok {
			myMap[v] += 1
		}
	}
	sum := 0

	for k, v := range myMap {
		sum += k * v
	}

	return sum
}
