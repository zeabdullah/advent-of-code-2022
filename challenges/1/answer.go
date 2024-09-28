package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("/Users/zabdullah/Desktop/projects/personal/my-advent-of-code/challenges/1/input.txt")
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)

	sum := 0
	for {
		currentLine, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}

		// the input will contain digits 1-9 (no zeros) within its lines
		var numPair = []int{0, 0}

		for _, char := range currentLine {
			// when reading the current line, check if the current char is a number
			if unicode.IsDigit(char) {
				num, err := strconv.Atoi(string(char))
				check(err)

				if numPair[0] == 0 && numPair[1] == 0 {
					// if it's the first digit that we run through, this means our numPair are both zeros still,
					// hence, set first and last number to it.
					numPair[0] = num
					numPair[1] = num
				} else {
					// otherwise, we are at our non-first digit,
					// so, set the second number in the pair to it
					numPair[1] = num
				}
			}
		}

		sum += (numPair[0] * 10) + numPair[1]

		if err == io.EOF {
			break
		}
	}

	fmt.Println(sum)

}

// Using AoC's input, your answer should = 54940
