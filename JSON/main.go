package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string   `json:"name"`
	LastName string   `json:"lastName"`
	Age      int      `json:"age"`
	Nickname string   `json:"nickname"`
	Stack    []string `json:"stack,omitempty"`
}

func main() {
	jsonCode()
	decodeJSON()
}

func jsonCode() {
	infoUser := []User{
		{"Ilya", "Shimaev", 21, "Umk1nus", []string{"js", "vue", "go", "html", "css"}},
		{"Daniil", "Jeludkov", 19, "Zybera", []string{"js", "php", "java", "html", "css"}},
		{"Anton", "Smirnov", 19, "Toxanski", []string{"js", "react", "php", "html", "css"}},
		{"Kirill", "Milytin", 20, "Ner7ul", nil},
	}

	finalJSON, err := json.MarshalIndent(infoUser, "", "\t")

	checkErrNil(err)
	fmt.Printf("%s\n", finalJSON)
}

func decodeJSON() {
	jsonDataFromWeb := []byte(`
	{
		"name": "Ilya",
		"lastName": "Shimaev",
		"age": 21,
		"nickname": "Umk1nus",
		"stack": ["js", "vue", "go", "html", "css"]
	}
`)

	var userItem User

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON valid")
		json.Unmarshal(jsonDataFromWeb, &userItem)
		fmt.Printf("%#v\n", userItem)
	} else {
		fmt.Println("JSON not valid")
	}
}

func checkErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
