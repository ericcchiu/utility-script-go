package main

import (
	"fmt"
	"utilitiy-scripts-go/src/fileprocessor"
)

func main() {
	fmt.Println(fileprocessor.FindIntegerDuplicates([]int{1, 2, 2, 3, 3, 4, 2, 0}))
	fmt.Println(fileprocessor.FindIntegerDuplicates([]int{1, 2, 3}))
	filePath := "./keys.txt"
	fmt.Println(fileprocessor.FindIntegerDuplicates(fileprocessor.ImportFileAsSlices(filePath)))
}
