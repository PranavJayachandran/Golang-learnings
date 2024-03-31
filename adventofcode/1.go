// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"unicode"
// )

// func main() {
// 	file, _ := os.Open("input.txt")

// 	scanner := bufio.NewScanner(file)

// 	var ans int = 0
// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		start := -1
// 		end := 0
// 		for _, char := range text {
// 			if unicode.IsDigit(char) {
// 				if start == -1 {
// 					start, _ = strconv.Atoi(string(char))
// 					end = start
// 				} else {
// 					end, _ = strconv.Atoi(string(char))
// 				}
// 			}
// 		}
// 		fmt.Println(start, end)
// 		ans = ans + start*10 + end
// 	}
// 	file.Close()
// 	fmt.Println(ans)
// }
