package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	var fileName, fileExt string
	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	fileName = strings.TrimSuffix(filePth, filepath.Ext(filePth))
	fileExt = strings.TrimPrefix(filePth, (fileName + "."))

	// Еще вариант
	// fileName = filePth[:len(filePth)-len(filepath.Ext(filePth))]
	// fileExt = filepath.Ext(filePth)

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
