package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Фильтруем в новый слайс
func filter(arr1 []int) []int {
	filtered := make([]int, 0, len(arr1))
	for _, v := range arr1 {
		if v%2 != 0 {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// Фильтруем оригинальный слайс
func filter1(arr1 []int) []int {
	var n int
	for _, v := range arr1 {
		if v%2 != 0 {
			arr1[n] = v
			n++
		}
	}

	return arr1[:n]
}

func insert(pos, value int, arr1 []int) []int {
	return append(arr1[:pos], append([]int{value}, arr1[pos:]...)...)
}

func insert1(pos, value int, arr1 []int) []int {
	arr1 = append(arr1, 0)

	copy(arr1[pos+1:], arr1[pos:])
	arr1[pos] = value
	return arr1
}

func main() {
	// Этот код нужен для ввода массива из стандартного ввода
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите элементы массива через пробел: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разбиваем строку на массив строк
	strValues := strings.Split(input, " ")

	// Преобразуем строки в числа и заполняем массив
	var arr []int
	for _, str := range strValues {
		val, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		arr = append(arr, val)
	}

	var pos int
	fmt.Println("Введите позицию для вставки")
	if _, err := fmt.Fscan(os.Stdin, &pos); err != nil {
		log.Fatal("Неправильное значение", err)
	}

	if pos < 0 || pos > len(arr)-1 {
		log.Fatal("Позиция вставки должна входить в диапазон индексов введенного слайса")
	}

	var value int
	fmt.Println("Введите значение для вставки")
	if _, err := fmt.Fscan(os.Stdin, &value); err != nil {
		log.Fatal("Неправильное значение")
	}
	if pos < 0 || pos > len(arr)-1 {
		log.Fatal("Позиция вставки должна входить в диапазон индексов введенного слайса")
	}

	// Скопируйте слайс arr в слайс arr1.
	// Это можно сделать через функцию Append или через функцию copy

	arr1 := make([]int, len(arr))
	copy(arr1, arr)
	arr1 = insert1(pos, value, arr1)

	fmt.Println(arr1)
	arr1 = filter(arr1)

	fmt.Println(arr1)
}
