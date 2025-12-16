package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/celsuss/advent-of-code-2025/utils"
)

func getDirAndRot(line string) (int, int) {
	// Get direction
	var dir int = 0
	if string(line[0]) == "R" {
		dir = 1
	} else {
		dir = -1
	}

	// Get length
	length, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}
	return dir, length
}

// Day one solution
func dayOne(input []string) int {
	var dial int = 50
	var pass int = 0
	for _, line := range input {
		dir, length := getDirAndRot(line)
		dial = dial + (length * dir)

		for dial >= 100 {
			dial -= 100
		}
		for dial < 0 {
			dial += 100
		}
		if dial == 0 {
			pass += 1
		}
	}
	return pass
}

func main() {
	// Day one part one test
	input, err := utils.ReadInput("../data/day-one-test.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(input, "\n")
	res := dayOne(lines)

	// Day one part one
	input, err = utils.ReadInput("../data/day-one.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines = strings.Split(input, "\n")
	res = dayOne(lines)

	fmt.Println(res)
}
