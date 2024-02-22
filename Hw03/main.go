package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Item struct {
	Name string
	Link string
	Tags string
	Date time.Time
}

func main() {
	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		_ = keyboard.Close()
	}()
	var item []Item
	fmt.Println("Программа для добавления url в список")
	fmt.Println(`
	- Для просмотра списка URL нажмите "L"
	- Для добавления URL нажмите "A"
	- Для удаления URL нажмите "R"
	- Для выхода нажмите "Esc"
	`)

OuterLoop:
	for {
		// Подключаем отслеживание нажатия клавиш
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильный аргументы в формате url описание теги")
				continue OuterLoop
			}

			newItem := Item{
				Name: args[0],
				Link: args[1],
				Tags: strings.Join(args[2:], ", "),
				Date: time.Now(),
			}

			item = append(item, newItem)
			fmt.Println("Ссылка добавлена!")

		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>

			if len(item) > 1 {
				fmt.Println("Список ссылок:")
				for i, link := range item {
					fmt.Printf("Имя: <%s>\nСсылка: <%s>\nТеги: <%s>\nДата: <%s>\n", link.Name, link.Link, link.Tags, link.Date.Format(time.DateTime))
					i++
				}
			}
		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление url из списка хранения
			fmt.Println("Введите имя ссылки, которое нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			_ = text

			for i, link := range item {
				if link.Name == text {
					item = append(item[:i], item[i+1:]...)
					fmt.Println("Ссылка удалена!")
					break
				}
			}

		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
