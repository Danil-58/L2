package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpackString(s string) string {
	var result strings.Builder
	var repeatCount int

	if len(s) == 0 {
		return ""
	}

	if unicode.IsDigit(rune(s[0])) {
		return "Некорректная строка"
	}

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			repeatCount, _ = strconv.Atoi(string(s[i]))
			prevChar := s[i-1]
			result.WriteString(strings.Repeat(string(prevChar), repeatCount))
		} else {
			result.WriteByte(s[i])
		}
	}
	return result.String()
}

func main() {
	var s string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&s)
	fmt.Println(unpackString(s))
}
