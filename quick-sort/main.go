package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/dickeyy/sorting/lib"
)

func quickSort(arr []int, low, high int, visualize bool) {
	if low < high {
		if visualize {
			fmt.Printf("\nCurrent array segment [%d:%d]: %v\n", low, high, arr[low:high+1])
		}

		p := partition(arr, low, high, visualize)
		quickSort(arr, low, p-1, visualize)
		quickSort(arr, p+1, high, visualize)
	}
}

func partition(arr []int, low, high int, visualize bool) int {
	pivot := arr[high]
	i := low - 1

	if visualize {
		fmt.Printf("\nPartitioning with pivot: %d\n", pivot)
	}

	for j := low; j < high; j++ {
		if visualize {
			// Visual representation of current state
			fmt.Printf("\nStep: i=%d, j=%d\n", i, j)
			for k := low; k <= high; k++ {
				if k == high {
					fmt.Printf("[%d]p ", arr[k]) // pivot
				} else if k == i {
					fmt.Printf("[%d]i ", arr[k]) // i pointer
				} else if k == j {
					fmt.Printf("[%d]j ", arr[k]) // j pointer
				} else {
					fmt.Printf("%d ", arr[k])
				}
			}
			fmt.Println()
		}

		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]

			if visualize && i != j {
				fmt.Printf("Swapped %d and %d\n", arr[i], arr[j])
			}
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	if visualize {
		fmt.Printf("\nAfter placing pivot:\n")
		for k := low; k <= high; k++ {
			if k == i+1 {
				fmt.Printf("[%d]* ", arr[k]) // final pivot position
			} else {
				fmt.Printf("%d ", arr[k])
			}
		}
		fmt.Printf("\nPivot placed at index: %d\n", i+1)
		fmt.Println(strings.Repeat("-", 50))
	}

	return i + 1
}

func main() {
	// Generate random array
	size := 10_000
	arr := lib.MakeRandomArray(size, 1, 10_001)

	// fmt.Printf("Initial array: %v\n", arr)

	start := time.Now()
	quickSort(arr, 0, len(arr)-1, false)
	elapsed := time.Since(start)

	// fmt.Printf("\nFinal sorted array: %v\n", arr)
	fmt.Printf("Time elapsed: %s for %d items\n", elapsed, size)
}
