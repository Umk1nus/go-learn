package main

import "fmt"

func main() {
	notes := [7]string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"g",
		"h",
	}
	for _, note := range notes {
		fmt.Println(note)
	}
}
