package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	buf, err := ioutil.ReadFile("./day1/input")
	if err != nil {
		log.Fatal(err)
	}

	inputBodyLines := strings.Split(strings.Trim(string(buf), "\n"), "\n")
	inputBodyNumbers := []int{}
	for _, i := range inputBodyLines {
		num, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		inputBodyNumbers = append(inputBodyNumbers, num)
	}

	log.Println("Part 1 Answer:", part1(inputBodyNumbers))
	log.Println("Part 2 Answer:", part2(inputBodyNumbers))
}

func part1(lines []int) int {
	peak := 0
	for _, num := range lines {
		peak += num
	}

	return peak
}

func part2(lines []int) int {
	peak := 0
	result := 0
	frequenciesList := make(map[int]bool)

	for true {
		for _, num := range lines {
			peak += num

			if frequenciesList[peak] {
				return peak
			}
			frequenciesList[peak] = true
		}
	}

	return result
}
