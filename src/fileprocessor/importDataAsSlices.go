package fileprocessor

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ImportFileAsSlices(filePath string) []int {

	var dataArr []int

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		intVar, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		dataArr = append(dataArr, intVar)
	}

	fmt.Println(dataArr)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dataArr
}
