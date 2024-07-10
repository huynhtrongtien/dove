package utilities

func SubSlice(arr []int, page, pageSize int) []int {
	lengh := len(arr)
	startIndex := (page - 1) * pageSize
	if startIndex >= lengh {
		return []int{}
	}

	endIndex := startIndex + pageSize
	if endIndex > lengh {
		endIndex = lengh
	}
	return arr[startIndex:endIndex]
}
