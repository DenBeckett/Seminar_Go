package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [100]int
	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	for i := range arr {
		value := rnd.Intn(20)
		arr[i] = value
	}

	slice := make([]int, 20)

	for _, v := range arr {
		slice[v]++
	}

	fmt.Println(slice)
}
