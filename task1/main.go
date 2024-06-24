package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)

/*
Напишите 2 функции: 1 функция читает ввод с консоли. Ввод одного значения заканчивается нажатием клавиши enter. Вторая
функция пишет эти данные в файл. Свяжите эти функции каналом. Работа приложения должна завершиться нажатием клавиш
ctrl + c.
*/

var isTesting = false

func main() {
	inputChan := make(chan string)
	stopChan := make(chan struct{})
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	file, err := os.Create("task1.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	go readInput(inputChan, stopChan, os.Stdin)
	go writeToFile(inputChan, stopChan, file)

	<-sigChan
	close(stopChan)

	fmt.Println("\nФайл сохранен.")
}

func writeToFile(inputChan <-chan string, stopChan <-chan struct{}, writer io.Writer) {
	for {
		select {
		case <-stopChan:
			return
		case input, ok := <-inputChan:
			if !ok {
				return
			}
			if _, err := writer.Write([]byte(input + "\n")); err != nil {
				fmt.Println("Ошибка записи в файл: ", err)
				return
			}
		}
	}
}

func readInput(inputChan chan<- string, stopChan <-chan struct{}, reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for {
		select {
		case <-stopChan:
			close(inputChan)
			return
		default:
			if !isTesting {
				fmt.Print("Введите текст: ")
			}
			if scanner.Scan() {
				inputChan <- scanner.Text()
			}
		}
	}
}
