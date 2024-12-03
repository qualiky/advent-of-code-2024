package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dayOneProblemTwo() {

	readFile, error := os.Open("dayOneProblemOneList.txt")
	check(error)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var listOne []int
	var listTwo []int

	for _, line := range fileLines {
		currentList := strings.Split(line, "   ")
		valueOne, err := strconv.Atoi(currentList[0])
		check(err)
		valueTwo, err := strconv.Atoi(currentList[1])
		check(err)
		listOne = append(listOne, valueOne)
		listTwo = append(listTwo, valueTwo)
	}

	fmt.Println(len(listOne))
	fmt.Println(len(listTwo))

	if len(listOne) != len(listTwo) {
		fmt.Println("These lists are not the same length, terminating")
		panic("Error: incorrect input")
	}

	sum := 0

	for i := 0; i < len(listOne); i++ {
		occurrence := 0
		for j := 0; j < len(listTwo); j++ {
			if listOne[i] == listTwo[j] {
				occurrence++
			}
		}
		sum += listOne[i] * occurrence
	}

	fmt.Println(sum)

}
