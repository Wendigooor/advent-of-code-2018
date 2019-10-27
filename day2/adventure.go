package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("./day2/input")
	if err != nil {
		log.Fatal(err)
	}

	inputBodyLines := strings.Split(strings.Trim(string(buf), "\n"), "\n")

	log.Println("Part 1 Answer:", part1(inputBodyLines))
	log.Println("Part 2 Answer:", part2(inputBodyLines))
}

func part1(lines []string) int {
	double := 0
	triple := 0
	alreadyFoundDouble := false
	alreadyFoundTriple := false
	frequenciesCharsList := make(map[string]int)

	for _, line := range lines {
		frequenciesCharsList = make(map[string]int)
		alreadyFoundDouble = false
		alreadyFoundTriple = false

		for _, char := range line {
			frequenciesCharsList[string(char)]++
		}

		for _, v := range frequenciesCharsList {
			if v == 2 && !alreadyFoundDouble {
				double++
				alreadyFoundDouble = true
			}
			if v == 3 && !alreadyFoundTriple {
				triple++
				alreadyFoundTriple = true
			}
		}
	}

	return double * triple
}

func part2(lines []string) string {
	result := ""
	found := false
	differentCharsCount := 0

	for _, line := range lines {
		if found {
			break
		}

		for _, nextLine := range lines {
			result = ""
			differentCharsCount = 0
			currentLineChars := strings.Split(line, "")
			nextLineChars := strings.Split(nextLine, "")

			for i := range nextLineChars {
				if currentLineChars[i] != nextLineChars[i] {
					differentCharsCount++
				} else {
					result = result + currentLineChars[i]
				}
			}

			if differentCharsCount == 1 {
				found = true
				break
			}
		}
	}

	return result
}
