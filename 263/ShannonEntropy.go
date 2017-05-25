package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func shannonEntropy(input string) float64 {
	charMap := make(map[rune]int)
	for _, c := range input {
		if _, ok := charMap[c]; !ok {
			charMap[c] = 1
		} else {
			charMap[c]++
		}
	}

	var sum float64
	length := float64(len(input))
	for _, cnt := range charMap {
		tmp := float64(cnt) / length
		sum += tmp * math.Log2(tmp)
	}
	return -1 * sum
}

func main() {
	input := bufio.NewReader(os.Stdin)
	var err error
	var line string
	for line, err = input.ReadString('\n'); err == nil; line, err = input.ReadString('\n') {
		line = line[:len(line)-1]
		fmt.Printf("%v : %.5f\n", line, shannonEntropy(line))
	}
}
