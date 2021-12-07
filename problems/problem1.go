package problems

import (
	"strconv"
)

func PartOne1() int64 {
	lines := readFile("test_files/1.txt")
	var numbers []int64
	for _, line := range lines {
		num, err := strconv.ParseInt(line, 10, 64)
		check(err)
		numbers = append(numbers, num)
	}
	goal := 2020
	var notUsed []int64
	//copy(notUsed, numbers)
	notUsed = append(numbers[:1], numbers[2:]...)
	for _, x := range numbers {
		for _, y := range notUsed {
			if x+y == int64(goal) {
				return x * y
			}
		}
		notUsed = append(notUsed[:1], notUsed[2:]...)
	}
	return 0
}

func PartTwo1() int64 {
	lines := readFile("test_files/1.txt")
	var numbers []int64
	for _, line := range lines {
		num, err := strconv.ParseInt(line, 10, 64)
		check(err)
		numbers = append(numbers, num)
	}
	goal := 2020

	for i := 0; i < len(numbers)-2; i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == int64(goal) {
					return numbers[i] * numbers[j] * numbers[k]
				}
			}
		}
	}
	return 0
}
