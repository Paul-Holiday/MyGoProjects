package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Reader interface {
	Read() string
}

type Writer interface {
	Write(text string)
}

type ReadWriter interface {
	Reader
	Writer
}

type StrIO struct {
	buffer string
}

func (io *StrIO) Read() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите текст: ")

	text, _ := reader.ReadString('\n')
	io.buffer = strings.TrimSpace(text)
	return io.buffer
}

func (io *StrIO) Write(text string) {
	fmt.Printf("Вывод: %s\n", text)
}

func (io *StrIO) GetBuffer() string {
	return io.buffer
}

func main() {
	io := &StrIO{}

	input := io.Read()

	io.Write("Вы ввели: " + input)

	fmt.Println("В буфере: " + io.GetBuffer())

}
