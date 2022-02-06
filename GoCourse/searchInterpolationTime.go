package main

import (
	"fmt"
	"time"
)

func main() {
	var toFind int = 0
	var posic int = 0
	//items := []int{2, 3, 5, 7, 11, 13, 17}
	//fmt.Printf("Datos indicados: %d\n", items)

	items := make([]int, 1000000)
	for i := 0; i < len(items); i++ {
		items[i] = i * 2
	}
	count := 100

	// Test de velocidad
	startTime := time.Now()
	toFind = 1234568

	for i := 0; i < count; i++ {
		posic = interpolationSearch(items, toFind)
	}

	deltaTime := time.Now().Sub(startTime)
	nanoseg := deltaTime.Nanoseconds() / int64(count)

	interpolationSearch(items, toFind)
	print(toFind, posic)
	fmt.Printf("La busqueda del numero: %d tardó: %d nanosegundos.\n", toFind, nanoseg)

}

func interpolationSearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1

	for (arr[low] < x) && (x < arr[high]) {
		var mid = low + ((x-arr[low])*(high-low))/(arr[high]-arr[low])

		if arr[mid] < x {
			low = mid + 1
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			return mid
		}
	}

	if arr[low] == x {
		return low
	}

	if arr[high] == x {
		return high
	}

	return -1
}

func print(num int, pos int) {
	if pos == -1 {
		fmt.Printf("El número: %d no fue encontrado!\n", num)
	} else {
		fmt.Printf("El número: %d se encuentra ubicado en la posición: %d\n", num, pos)
	}
}
