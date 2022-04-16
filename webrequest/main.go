package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/Umk1nus"

func main() {

	response, err := http.Get(url)

	checkNilErr(err)

	fmt.Printf("Тип респонса %T", response)

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)

	checkNilErr(err)
	content := string(databyte)
	fmt.Println(content)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
