package main

import (
	"bufio"
	"fmt"
	"os"
)

const factor = 100

func getValue(lines []string) int {
	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] {
			continue
		}
		x, y := i, i-1
		works := true
		for j := 0; j < min(i, len(lines)-i)-1; j++ {
			x++
			y--
			if lines[x] != lines[y] {
				works = false
				break
			}
		}
		if works {
			return i
		}
	}

	return 0

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	var lines []string
	var vLines []string

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			sum += factor * getValue(lines)
			lines = lines[:0]
			sum += getValue(vLines)
			vLines = vLines[:0]
			continue
		}
		if len(vLines) == 0 {
			for _, v := range text {
				vLines = append(vLines, string(v))
			}
		} else {
			for k, v := range text {
				vLines[k] = vLines[k] + string(v)
			}

		}
		lines = append(lines, text)
	}

	sum += factor * getValue(lines)
	sum += getValue(vLines)
	fmt.Println(sum)
}
