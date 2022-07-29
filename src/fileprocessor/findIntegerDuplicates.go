package fileprocessor

//FindIntegerDuplicates function takes in a slice of integers and returns a list of integers containing duplicates.
// It will return an empty list if no duplicate is found
func FindIntegerDuplicates(arr []int) []int {
	itemMap := createDuplicateMap(arr)
	var duplicateList []int

	for key, val := range itemMap {
		if val > 1 {
			duplicateList = append(duplicateList, key)
		}
	}
	return duplicateList
}

func createDuplicateMap(arr []int) map[int]int {
	itemMap := make(map[int]int, 0)
	for i := 0; i < len(arr); i++ {
		_, duplicate := itemMap[arr[i]]
		if duplicate {
			itemMap[arr[i]] += 1
		} else {
			itemMap[arr[i]] = 1
		}
	}
	return itemMap
}
