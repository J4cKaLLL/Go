package main

import (
	"fmt"
	"sort"
)

func main() {

	var arr []int = []int{7, 7, 7, 7, 7, 7, 7}
	mapa1 := make(map[int][]int, len(arr))
	//mapa2 := make(map[int][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		mapa1[arr[i]] = append(mapa1[arr[i]], len(arr))
	}
	fmt.Println(mapa1)

	var values []int
	var valuesCount int
	for v := range mapa1 {
		values = append(values, len(mapa1[v]))
		valuesCount += len(mapa1[v])
	}

	sort.Ints(values)

	var count int = 0
	for i := 0; i < len(values)-1; i++ {
		for j := 0; j < len(values)-1; j++ {
			if (values[i] + values[j+1]) < valuesCount/2 {
				count++
			}
		}
	}
	fmt.Println("count", count)

}
