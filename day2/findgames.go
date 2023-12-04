// go main function
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 12 red cubes, 13 green cubes, and 14 blue cubes
var arrayBag = []int{12, 13, 14}

func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}
func main() {
	sum := 0
	var num int

	maxRed := 1
	maxGreen := 1
	maxBlue := 1
	// count := 0
	// read file input
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the file contents
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)

		//split line into map
		parts := strings.Split(line, ":")
		games := strings.Split(parts[1], ";")
		// fmt.Println(gameNum)
		// fmt.Println(games)
		for _, game := range games {
			// fmt.Println(game)
			pieces := strings.Split(game, ",")
			// fmt.Println(pieces)
			// fmt.Println(pieces[0][1] - 48)
			for _, piece := range pieces {
				if isNumber(piece[2]) {
					num = int(piece[1]-48)*10 + int(piece[2]-48)
				} else {
					num = int(piece[1] - 48)
				}
				fmt.Println("num: ", num)
				if strings.Index(piece, "red") != -1 {
					// fmt.Println("red")
					if num > maxRed {
						maxRed = num
					}

				}
				if strings.Index(piece, "green") != -1 {
					// fmt.Println("green")
					if num > maxGreen {
						maxGreen = num
					}
				}
				if strings.Index(piece, "blue") != -1 {
					// fmt.Println("blue")
					if num > maxBlue {
						maxBlue = num
					}
				}
			}

		}

		fmt.Println("MaxRed: ", maxRed, " MaxGreen: ", maxGreen, " MaxBlue:", maxBlue)
		minPower := maxRed * maxGreen * maxBlue
		sum += minPower
		fmt.Println("which ", sum)
		// sum += game_num

		//reset
		maxRed = 1
		maxGreen = 1
		maxBlue = 1

		// Continue with the rest of your code
	}
	fmt.Println(sum)
}
