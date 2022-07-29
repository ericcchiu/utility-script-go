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

	//defer file.Close()
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		intVar, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		dataArr = append(dataArr, intVar)
	}

	fmt.Println(dataArr)
	return dataArr
}
