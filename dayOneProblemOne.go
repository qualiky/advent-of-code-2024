package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func dayOneProblemOne() {

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

	sort.Ints(listOne)
	sort.Ints(listTwo)

	for i := 0; i < len(listOne); i++ {
		if listOne[i] > listTwo[i] {
			sum += listOne[i] - listTwo[i]
		} else {
			sum += listTwo[i] - listOne[i]
		}
	}

	fmt.Println(sum)

}
