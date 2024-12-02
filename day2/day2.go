package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func removeElement(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func processText(text string) bool {

	const maxErrorCount = 1
	var split = strings.Split(text, " ")
	var splitNumber []int
	var duplicateSplitNumber []int

	for _, element := range split {
		val, err := strconv.Atoi(element)
		if err != nil {
			log.Fatal(err)
		}
		splitNumber = append(splitNumber, val)
		duplicateSplitNumber = append(duplicateSplitNumber, val)
	}

	var prevVal int
	validReport := true
	inc := true
	deleteIndex := 0

	for index := 0; index < len(splitNumber); index++ {
		val := splitNumber[index]
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
			if deleteIndex < len(duplicateSplitNumber) {
				splitNumber = removeElement(duplicateSplitNumber, deleteIndex)
				deleteIndex++
				index = -1
				continue
			}
			validReport = false
			break
		}

		prevVal = val
	}
	return validReport
}

func main() {
	// f, err := os.Open("testinput.txt")
	f, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	valid := 0
	invalid := 0

	for scanner.Scan() {
		text := scanner.Text()

		validReport := processText(text)

		if validReport {
			valid++
		} else {
			invalid++
		}

	}

	fmt.Println("output:", valid)
}
