package problems

import (
	"fmt"
	"strconv"
	"strings"
)

func PartOne2() int64 {
	lines := readFile("test_files/2.txt")
	incorrect := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		ran := strings.Split(parts[0], "-")
		min, err := strconv.ParseInt(ran[0], 10, 64)
		check(err)
		max, err := strconv.ParseInt(ran[1], 10, 64)
		check(err)

		rule := Rule{min, max, parts[1][:1]}
		fmt.Printf("%+v\n", rule)
		count := int64(countGlyphs(parts[2], rule.character))

		if count > max || count < min {
			incorrect += 1
		}

	}

	return int64(incorrect)
}

func PartTwo2() int64 {
	lines := readFile("test_files/2.txt")
	correct := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		ran := strings.Split(parts[0], "-")
		min, err := strconv.ParseInt(ran[0], 10, 64)
		check(err)
		max, err := strconv.ParseInt(ran[1], 10, 64)
		check(err)

		rule := Rule{min, max, parts[1][:1]}
		// fmt.Printf("%+v\n", rule)

		if validateRule(parts[2], rule) {
			correct += 1
		}

	}

	return int64(correct)
}

func countGlyphs(source string, glyph string) int {
	count := 0
	for _, c := range source {
		if string(c) == glyph {
			// fmt.Printf("%c %s\n", c, glyph)
			count += 1
		}
	}
	return count
}

func validateRule(source string, rule Rule) bool {
	// println(source[rule.maximum-1 : rule.maximum])
	firstCorrect := source[rule.minimum-1:rule.minimum] == rule.character
	secondCorrect := source[rule.maximum-1:rule.maximum] == rule.character
	return firstCorrect && !secondCorrect || !firstCorrect && secondCorrect
}

type Rule struct {
	minimum   int64
	maximum   int64
	character string
}
