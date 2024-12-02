package main

import (
	"bufio"
	"fmt"
	"log"
	// "math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	valid := 0
	for scanner.Scan() {
		text := scanner.Text()
		var split = strings.Split(text, " ")

		var prevVal int
		validReport := true
		inc := true

		for index, element := range split {
			val, err := strconv.Atoi(element)
			if err != nil {
				log.Fatal(err)
			}
			if index == 0 {
				prevVal = val
				continue
			}

			differ := (float64(prevVal) - float64(val))

			if index == 1 {
				if differ > 0 {
					inc = false
				} else {
					inc = true
				}

			}

			if (inc && (differ >= 0 || differ < -3)) || (!inc && (differ <= 0 || differ > 3)) {
				validReport = false
				break
			}

			prevVal = val
		}

		if validReport {
			valid++
		}
	}

	fmt.Println("output:", valid)
}
