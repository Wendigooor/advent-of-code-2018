package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Claim struct {
	id         uint
	x, y, w, h uint
}

func main() {
	buf, err := ioutil.ReadFile("./day3/input")
	if err != nil {
		log.Fatal(err)
	}

	inputBodyLines := strings.Split(strings.Trim(string(buf), "\n"), "\n")

	var fabric [1000][1000]int
	var claims []Claim
	var claim Claim
	var count int

	for _, line := range inputBodyLines {
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &claim.id, &claim.x, &claim.y, &claim.w, &claim.h)
		claims = append(claims, claim)

		for i := claim.y; i < claim.y+claim.h; i++ {
			for j := claim.x; j < claim.x+claim.w; j++ {
				if fabric[i][j] == 0 {
					fabric[i][j] = 1
				} else if fabric[i][j] == 1 {
					count++
					fabric[i][j]++
				}
			}
		}
	}

	log.Println("Part 1 Answer:", count)

	for index, claim := range claims {
		var hasOverlap bool

		for i := claim.y; i < claim.y+claim.h; i++ {
			if hasOverlap {
				break
			}

			for j := claim.x; j < claim.x+claim.w; j++ {
				if fabric[i][j] != 1 {
					hasOverlap = true
					break
				}
			}
		}

		if !hasOverlap {
			log.Println("Part 2 Answer:", index+1)
		}
	}
}
