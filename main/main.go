// https://www.codewars.com/kata/5946a0a64a2c5b596500019a/solutions/go

package main

import (
	"fmt"
	"math"
)

func SplitAndAdd(numbers []int, n int) []int {
	for i := 0; i < n; i++ {
		var size = len(numbers)                                                        // Current size of the numbers array
		var mid = size / 2                                                             // The middle index of the array
		var left = numbers[0:mid]                                                      // The left part of the array
		var right = numbers[mid:]                                                      // The right part of the array
		var leftSize = len(left)                                                       // The size of the left part of the array
		var rightSize = len(right)                                                     // The size of the right part of the array
		var leftPointer = leftSize - 1                                                 // The pointer to the left array. Starting from the end of the values in the array.
		var rightPointer = rightSize - 1                                               // The pointer to the right array. Starting from the end of the values in the array.
		var resultPointer = int(math.Max(float64(leftPointer), float64(rightPointer))) // The pointer to the result array. Starting from the end position.
		var result = make([]int, int(math.Ceil(float64(size)/2.0)))                    // The result array

		for leftPointer >= 0 || rightPointer >= 0 {
			// Assume the value of the left and right position is 0. This is to cater the length of both arrays are not the same.
			var leftValue, rightValue = 0, 0

			if leftPointer >= 0 {
				// There is still item to check on the left array. Grab the value.
				leftValue = left[leftPointer]
			}

			if rightPointer >= 0 {
				// There is still item to check on the right array. Grab the value.
				rightValue = right[rightPointer]
			}

			// Add the left and right value and put it into the result array
			result[resultPointer] = leftValue + rightValue

			// Move all pointers to the left 1 position
			resultPointer--
			leftPointer--
			rightPointer--
		}

		// Mutate the original numbers to be the result and use this for further calculation
		numbers = result
	}

	return numbers
}

func main() {
	fmt.Println(SplitAndAdd([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(SplitAndAdd([]int{1, 2, 3, 4, 5}, 3))
}
