package main

import (
	"bufio"
	"fmt"
	"os"
	_ "unicode"
)

// Такой вариант тоже принимается, без очистки строки

// for _, l := range text {
// 	if unicode.IsSpace(l) {
// 		wordLine++
// 	}
//
// 	if unicode.IsLetter(l) {
// 		letterNumber++
// 	}
// }

// Можно и так в один проход, сделано на семинаре уже

// bytesNumber = len(text)
// var prev rune
// for _, s := range text {
// 	if unicode.IsLetter(s) {
// 		if unicode.IsSpace(prev) {
// 			wordLine++
// 		}
// 		symbolNumber++
// 	}
// 	prev = s
// }

// wordLine++

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	var wordLine, symbolNumber, bytesNumber int
	var prev rune    // предыдущий символ
	var clean string // очищенная от лишних пробелов строка
	for _, l := range text {
		if unicode.IsLetter(l) {
			clean += string(l)
		}
		if !(unicode.IsSpace(prev) && unicode.IsSpace(l)) {
			clean += string(l)
		}

		prev = l
		if unicode.IsLetter(l) {
			letterNumber++
		}
	}

	for _, l := range clean {
		if unicode.IsSpace(l) {
			wordLine++
		}
	}

	bytesNumber = len(text)

	fmt.Printf("Количество слов: %d", wordLine)
	fmt.Printf("Количество букв: %d", symbolNumber)
	fmt.Printf("Количество байт: %d", bytesNumber)
}
