package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	random := rand.Intn(100) + 1
	fmt.Println("Выбираю число от 1 до 100")
	fmt.Println("Число выбрано")
	reader := bufio.NewReader(os.Stdin)
	succeed := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("Количество ваших попыток: ", 10-guesses)
		fmt.Println("Напишите ваше число")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if random == guess {
			fmt.Println("Числа совпали")
			succeed = true
			break
		} else {
			fmt.Println("Числа не совпали")
		}
	}
	if !succeed {
		fmt.Println("Попытки закончились, число:", random)
	}
}
