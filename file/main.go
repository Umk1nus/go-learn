package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "github.com/Umk1nus"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length = ", length)

	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("Текст внутри данных файла", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
