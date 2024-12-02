package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// f, err := os.Open("testinput.txt")
	f, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var arr1 []int
	var arr2 []int

	for scanner.Scan() {
		var text = scanner.Text()

		var split = strings.Split(text, "   ")
		// fmt.Printf("1:%s 2:%s\n", split[0], split[1])

		i, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		arr1 = append(arr1, i)

		j, err2 := strconv.Atoi(split[1])
		if err2 != nil {
			log.Fatal(err2)
		}
		arr2 = append(arr2, j)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// sort the array
	slices.Sort(arr1)
	slices.Sort(arr2)

	var final int64 = 0
	tmpptr := 0
	for i := 0; i < len(arr1) && tmpptr < len(arr1); i++ {
		count := 0
		tmpptr = 0
		for arr2[tmpptr] <= arr1[i] {
			if arr2[tmpptr] < arr1[i] {
				tmpptr++
				continue
			}
			tmpptr++
			count++
			if tmpptr >= len(arr2) {
				break
			}
		}
		final = final + int64(arr1[i]*count)

	}

	fmt.Println("final value:", int(final))
}
