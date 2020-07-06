package main

import (
	"fmt"
)

func main() {
	array := []float64{11, 1, 3, 2, 4, 6, 8, 7, 5, 9, 10}
	mergeArray := make([]float64, len(array))
	copy(mergeArray, array)
	insertionArray := make([]float64, len(array))
	copy(insertionArray, array)
	selectionArray := make([]float64, len(array))
	copy(selectionArray, array)
	bubbleArray := make([]float64, len(array))
	copy(bubbleArray, array)
	mergeSort(mergeArray, 0, 10)
	fmt.Println(mergeArray)
	insertionSort(insertionArray)
	fmt.Println(insertionArray)
	selectionSort(selectionArray)
	fmt.Println(selectionArray)
	bubbleSort(bubbleArray)
	fmt.Println(bubbleArray)

}

// mergeSort sorts an array using the mergeSort method
func mergeSort(array []float64, start, end int) {
	if start < end {
		mid := (start + end) / 2
		mergeSort(array, start, mid)
		mergeSort(array, mid+1, end)
		merge(array, start, mid, end)
	}

}

func merge(array []float64, start, mid, end int) {
	L := make([]float64, (mid + 1 - start))
	copy(L, array[start:mid+1])
	R := make([]float64, (end + 1 - mid - 1))
	copy(R, array[mid+1:end+1])
	i, j, k := 0, 0, start
	for i < len(L) && j < len(R) {
		if L[i] <= R[j] {
			array[k] = L[i]
			i++
			k++

		} else {
			array[k] = R[j]
			k++
			j++
		}
	}

	if j == len(R) {
		for i := i; i < len(L); i++ {
			array[k] = L[i]
			k++
		}
	}
}

// Sorts an array using the insertion sort method
func insertionSort(array []float64) {
	for j := 1; j < len(array); j++ {
		i := j - 1
		key := array[j]
		for i >= 0 && key < array[i] {
			array[i+1] = array[i]
			i--
		}
		array[i+1] = key
	}
}

//Sorts an array using selection sort method
func selectionSort(array []float64) {
	n := len(array)
	for i := 0; i < n; i++ {
		minIndex := i
		for index := i; index < n; index++ {
			if array[index] < array[minIndex] {
				minIndex = index
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]

	}
}

// sorts an array using the bubble sort method
func bubbleSort(array []float64) {
	n := len(array)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}
