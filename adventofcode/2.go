package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func findNumber(t string) int {
	if t == "one" {
		return 1
	} else if t == "two" {
		return 2
	} else if t == "three" {
		return 3
	} else if t == "four" {
		return 4
	} else if t == "five" {
		return 5
	} else if t == "six" {
		return 6
	} else if t == "seven" {
		return 7
	} else if t == "eight" {
		return 8
	} else if t == "nine" {
		return 9
	} else {
		return -1
	}
}
func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	var ans int = 0
	for scanner.Scan() {
		text := scanner.Text()
		start := -1
		end := 0
		for i := 0; i < len(text); i++ {
			var number int = -1
			if i+3 <= len(text) {
				number = findNumber(text[i : i+3])
			}
			if number == -1 && i+4 <= len(text) {
				number = findNumber(text[i : i+4])
			}
			if number == -1 && i+5 <= len(text) {
				number = findNumber(text[i : i+5])
			}
			if number != -1 {
				if start == -1 {
					start = number
					end = start
				} else {
					end = number
				}
				continue
			}
			if digit, err := strconv.Atoi(string(text[i])); err == nil {
				if start == -1 {
					start = digit
					end = start
				} else {
					end = digit
				}
			}
		}
		fmt.Println(start, end)
		ans = ans + start*10 + end
	}
	file.Close()
	fmt.Println(ans)
}
