package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func checkEqualDistance(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0
	}
	distance := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != distance {
			return false, 0
		}
	}
	return true, distance
}

func main() {
	var input int
	var data []int

	fmt.Println("Masukkan sekumpulan bilangan bulat (diakhiri dengan bilangan negatif):")
	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	insertionSort(data)

	isEqualDistance, distance := checkEqualDistance(data)

	for _, val := range data {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	if isEqualDistance {
		fmt.Printf("Data berjarak %d\n", distance)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
