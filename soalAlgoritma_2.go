package main

import (
	"fmt"
	"sort"
	"math"
)

func main(){

	numbers := [32]int{1,2,33,44,55,33,44,22,44,33,2,77,66,1,2,3,4,5,6,7,89,3,3,8,9,75,4,3,2,2,1,3}
	var subNumbersLen int = len(numbers)/3
	subNumbers1 := numbers[0:subNumbersLen]
	subNumbers2 := numbers[subNumbersLen:subNumbersLen*2]
	subNumbers3 := numbers[subNumbersLen*2:len(numbers)]
	
	sort.Sort(sort.Reverse(sort.IntSlice(subNumbers1)))
	sort.Sort(sort.Reverse(sort.IntSlice(subNumbers2)))
	sort.Sort(sort.Reverse(sort.IntSlice(subNumbers3)))

	fmt.Printf("\nShorting \n")

	fmt.Println(subNumbers1)
	fmt.Println(subNumbers2)
	fmt.Println(subNumbers3)
	
	var sumSubNumbers1 = sum(subNumbers1)
	var sumSubNumbers2 = sum(subNumbers2)
	var sumSubNumbers3 = sum(subNumbers3)
	
	fmt.Printf("\nTotal Penjumlahan \n")

	fmt.Println(sumSubNumbers1)
	fmt.Println(sumSubNumbers2)
	fmt.Println(sumSubNumbers3)

	fmt.Printf("\nRata-Rata \n")
	
	fmt.Println(average(sumSubNumbers1, len(subNumbers1)))
	fmt.Println(average(sumSubNumbers2, len(subNumbers2)))
	fmt.Println(average(sumSubNumbers3, len(subNumbers3)))

	fmt.Printf("\nMin & Max \n")
	
	min, max := minmax(subNumbers1)
	fmt.Printf("Kelompok 1 : min=%d max=%d\n",min,max)
	min, max = minmax(subNumbers2)
	fmt.Printf("Kelompok 2 : min=%d max=%d\n",min,max)
	min, max = minmax(subNumbers3)
	fmt.Printf("Kelompok 3 : min=%d max=%d\n",min,max)

}

func sum(numbers []int) int{
	result := 0
	for _, number := range numbers{
		result += number
	}
	return result
}

func average(sum int, total int) float32{
	return float32(sum)/float32(total)
}

func minmax(numbers []int) (int, int){
	var min = math.Inf(1)
	var max = math.Inf(-1)

	for _, number := range numbers{
		if(float64(number) > max){
			max = float64(number)
		}
		if(float64(number) < min){
			min = float64(number)
		}
	}

	return int(min), int(max)
}