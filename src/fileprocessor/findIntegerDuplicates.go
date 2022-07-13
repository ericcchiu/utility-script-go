package fileprocessor

//FindDuplicates function takes in a slice of integers and returns a list of integers containing duplicates
func FindIntegerDuplicates(arr []int) []int {

	itemMap := make(map[int]bool, 0)
	var duplicateList []int
	for i := 0; i < len(arr); i++ {
		if itemMap[arr[i]] == true {
			duplicateList = append(duplicateList, arr[i])
		} else {
			itemMap[arr[i]] = true
		}
	}
	return duplicateList
}
