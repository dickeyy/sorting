package main

import "fmt"

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

// partition partitions the array around a pivot value
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{4, 2, 1, 3, 5, 6, 7, 8, 9, 10}
	quickSort(arr, 0, len(arr)-1)

	for _, v := range arr {
		fmt.Println(v)
	}
}
