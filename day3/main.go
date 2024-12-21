package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findNextIndex(idxArr []int, curIdx int) (int, int) {
	fmt.Println("INPUT:", idxArr, curIdx)
	for idx, ele := range idxArr {
		if ele > curIdx {
			return ele, idx
		}
	}
	return idxArr[len(idxArr)-1], len(idxArr) - 1
}

func removeDisabledSections(str string) string {
	do_r, _ := regexp.Compile("do\\(\\)")
	dont_r, _ := regexp.Compile("don't\\(\\)")
	ptr, maxlength := 0, len(str)

	_do_idx := do_r.FindAllStringIndex(str, -1)
	_dont_idx := dont_r.FindAllStringIndex(str, -1)

	var do_idx []int
	var dont_idx []int

	for _, e := range _do_idx {
		do_idx = append(do_idx, e[0])
	}

	for _, e := range _dont_idx {
		dont_idx = append(dont_idx, e[0])
	}

	enabledIndex := 0
	disabledIndex := 0

	new_str := ""

	fmt.Println("DO", do_idx)
	fmt.Println("DONT", dont_idx)

	var temp int
	for ptr < maxlength {

		for disabledIndex <= enabledIndex {

			if len(dont_idx) == 0 {
				fmt.Println("EXIT code 1 ????", enabledIndex, " ", disabledIndex)
				return new_str + str[enabledIndex:]
			}
			disabledIndex, temp = findNextIndex(dont_idx, disabledIndex)
			dont_idx = dont_idx[temp:]

			if len(dont_idx) == 1 {
				if disabledIndex <= enabledIndex {
					fmt.Println("EXIT code 2 ????", enabledIndex, " ", disabledIndex)
					return new_str + str[enabledIndex:]
				}
				break
			}
		}

		new_str = new_str + str[enabledIndex:disabledIndex]
		ptr = disabledIndex + 1

		for enabledIndex < disabledIndex {
			enabledIndex, temp = findNextIndex(do_idx, enabledIndex)
			do_idx = do_idx[temp:]
			if len(do_idx) == 1 {
				if enabledIndex <= disabledIndex {
					fmt.Println("EXIT code 4 ????", enabledIndex, " ", disabledIndex)
					return new_str
				}
				break
			}
		}
	}

	fmt.Println("EXIT code 5 ????")
	return new_str
}

func getValidValues(str string) [][]int {
	var finalValues [][]int

	fmt.Println("OLDSTRING:", str)
	new_str := removeDisabledSections(str)
	fmt.Println("NEWSTRING:", new_str)

	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	final := r.FindAllString(new_str, -1)
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
