package main

import (
	"fmt"
	"strings"
)

func main() {

	sentence := "Hello World"
	sentence = strings.ToLower(sentence)
	// fmt.Println(sentence)
	letterCount := make(map[rune]int)
	totalLetters := 0

	// Подсчитываем количество каждой буквы
	for _, char := range sentence {
		if char != ' ' {
			letterCount[char]++
			totalLetters++
		}
	}

	// Выводим результат
	fmt.Println("Результат:")
	for char, count := range letterCount {
		percent := (float64(count) / float64(totalLetters)) * 100
		fmt.Printf("%c - %d %.2f %%\n", char, count, percent)
	}
}
