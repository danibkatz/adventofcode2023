// go main function
package main

import (
	"bufio"
	"fmt"
	"os"
)

// check if character is a number
func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

func main() {
	sum := 0
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
		// Process each line of the file here

		//declare slice
		numbersInLine := []int{}

		for i := 0; i < len(line); i++ {
			if isNumber(line[i]) {
				numbersInLine = append(numbersInLine, int(line[i])-48)
				fmt.Println(int(line[i])-48, " ", line[i])
			}
		}
		fmt.Println(numbersInLine)
		if len(numbersInLine) == 0 {
			fmt.Println("No numbers in line")
		} else if len(numbersInLine) == 1 {
			sum = sum + numbersInLine[0]*10 + numbersInLine[0]
		} else {
			sum = sum + numbersInLine[0]*10 + numbersInLine[len(numbersInLine)-1]
		}
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
