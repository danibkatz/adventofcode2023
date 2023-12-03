// go main function
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var numberWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var mapWords = make(map[int]int)

// check if character is a number
func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

//function to identify whole word numbers

func findNumberWord(word string) {
	ind := 0
	for i, number := range numberWords {
		ind = strings.Index(word, number)
		if ind == -1 {
			continue
		}
		// fmt.Println(word, "  ", "Found", number, "at index", ind, "as", i+1)
		mapWords[ind] = i + 1
	}
}

func findNumber(word string) {
	numbersInLine := []int{}
	for i := 0; i < len(word); i++ {
		if isNumber(word[i]) {
			numbersInLine = append(numbersInLine, int(word[i])-48)
			// fmt.Println(int(word[i])-48, " ", word[i])
			mapWords[i] = int(word[i]) - 48
		}
	}
}

func sort_map(m map[int]int) {

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, ":", m[k])
	}
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

		fmt.Println("Working on line: ", line)
		// find words and numbers
		findNumberWord(line)
		findNumber(line)

		if len(mapWords) == 1 {
			fmt.Println("Adding ", mapWords[0]*10+mapWords[0])
			sum = sum + mapWords[0]*10 + mapWords[0]
			continue
		}
		sort_map(mapWords)
		maxKey := 0
		minKey := 1000
		for k := range mapWords {
			if k > maxKey {
				maxKey = k
			}
			if k < minKey {
				minKey = k
			}
		}
		fmt.Println("First word: ", mapWords[minKey], " Last word: ", mapWords[maxKey])
		fmt.Println("Adding ", mapWords[minKey]*10+mapWords[maxKey])
		sum = sum + mapWords[minKey]*10 + mapWords[maxKey]

		// clear the map every time
		for k := range mapWords {
			delete(mapWords, k)
		}

	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
