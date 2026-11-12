package main

import "fmt"

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	pivot := arr[(l+r)/2]
	i, j := l, r
	for i <= j {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	quickSort(arr, l, j)
	quickSort(arr, i, r)
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	quickSort(a, 0, n-1)
	for idx, v := range a {
		if idx > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
}
