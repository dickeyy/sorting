package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/dickeyy/sorting/lib"
	"golang.org/x/exp/rand"
)

func binarySearch(arr []int, target int, visualize bool) int {
	low := 0
	high := len(arr) - 1
	step := 1

	if visualize {
		fmt.Printf("\nSearching for: %d\n", target)
		fmt.Printf("Initial array: %v\n", arr)
	}

	for low <= high {
		mid := (low + high) / 2

		if visualize {
			fmt.Printf("\nStep %d:\n", step)
			fmt.Printf("Low: %d, High: %d, Mid: %d\n", low, high, mid)
			fmt.Printf("Current element: %d\n", arr[mid])

			// Visual representation of the search space
			for i := 0; i < len(arr); i++ {
				if i == mid {
					fmt.Printf("[%d]", arr[i])
				} else if i >= low && i <= high {
					fmt.Printf(" %d ", arr[i])
				} else {
					fmt.Printf(" . ")
				}
			}
			fmt.Println()
		}

		if arr[mid] == target {
			if visualize {
				fmt.Printf("\nFound %d at index %d!\n", target, mid)
			}
			return mid
		} else if arr[mid] < target {
			if visualize {
				fmt.Printf("Target %d is larger, searching right half\n", target)
			}
			low = mid + 1
		} else {
			if visualize {
				fmt.Printf("Target %d is smaller, searching left half\n", target)
			}
			high = mid - 1
		}

		step++
	}

	if visualize {
		fmt.Printf("\nTarget %d not found in array\n", target)
	}
	return -1
}

func main() {
	// Seed the random number generator
	rand.Seed(uint64(time.Now().UnixNano()))

	// Generate random sorted array
	size := 10_000

	min := 1
	max := rand.Int()

	items := lib.MakeRandomArray(size, min, max)
	sort.Ints(items)

	// Pick random target from the array
	target := items[rand.Intn(len(items))]

	fmt.Printf("Array Length: %d\n", size)
	fmt.Printf("Target: %d\n", target)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	// time
	start := time.Now()

	res := binarySearch(items, target, false)
	fmt.Printf("Found at index: %d\n", res)

	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s for %d items\n", elapsed, size)
}
