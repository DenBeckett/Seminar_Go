package main

import (
	"fmt"
	"strings"
)

// На основе шаблона напишите программу, подсчитывающую сколько раз буква встречается в предложении, 
// а также частоту встречаемости в процентах. То есть в предложении hello world результатом будет:
// h - 1 0.1
// e - 1 0.1
// l - 3 0.33
// o - 2 0.2
// w - 1 0.1
// r - 1 0.1
// d - 1 0.1

func main() {
	var input string
	input = "hello world"
	input = strings.ReplaceAll(input, " ", "")
	frequency := countCharacters(input)
	fmt.Println("Character frequency:")
	for char, count := range frequency {
		fmt.Printf("%c: %d %0.1f\n", char, count, float32(count)/float32(len(frequency)))
	}
}

func countCharacters(str string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range str {
		frequency[char]++
	}

	return frequency
}