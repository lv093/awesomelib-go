package insertion

func InsertSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	for i := 1; i < len(arr); i ++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && key < arr[j] {
			arr[j + 1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}

	return arr
}