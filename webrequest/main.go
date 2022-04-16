package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://go.dev/"

func main() {

	response, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Тип респонса %T", response)

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)

	content := string(databyte)
	fmt.Println(content)

	file, err := os.Create("./index.html")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println(length)

	defer file.Close()

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
