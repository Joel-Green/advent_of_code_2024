package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getValidValues(str string) [][]int {
	var finalValues [][]int
	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	final := r.FindAllString(str, -1)
	for _, element := range final {
		val := strings.Split(element[4:len(element)-1], ",")

		x, _ := strconv.Atoi(val[0])
		y, _ := strconv.Atoi(val[1])

		finalValues = append(finalValues, []int{x, y})
	}

	return (finalValues)
}

func readInput(filepath string) (string, error) {
	dat, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func mulAndAdd(val [][]int) int {
	sum := 0
	for _, element := range val {
		sum = sum + (element[0] * element[1])
	}
	return sum
}

func processAndReturnData(filepath string) int {
	s, err := readInput(filepath)
	if err != nil {
		panic(err)
	}
	val := getValidValues(s)
	return mulAndAdd(val)
}

func main() {
	x := processAndReturnData("./input.txt")
	fmt.Println(x)

}
