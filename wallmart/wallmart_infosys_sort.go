package wallmart

import (
	"fmt"
	"sort"
)

// Arr1 = [3,76,43,8]
// Arr2=[12,5,67,54,7]
// Res_arr = [3,5,7,8,12,43,54,67,76]
func WithoutPredefinesMethods() {
	fmt.Println("Hello sorting...")	
	arr1 := []int{3, 76, 43, 8}
	arr2 := []int{12, 5, 67, 54, 7}
	arr3 := make([]int, len(arr1)+len(arr2))
	
	// for i:=0; i<len(arr1); i++ {
	// 	arr3[i] = arr1[i]
	// }
	copy(arr3, arr1)
	for i:=0; i<len(arr2); i++ {
		arr3[len(arr1)+i] = arr2[i]
	}
	//copy(arr3, arr2)

	var temp int 
	fmt.Printf("merged arrays : %v\n",arr3)
	for i:=0; i<len(arr3); i++ {
		for j:=i; j<len(arr3); j++ {
			if arr3[i] > arr3[j] {
				temp = arr3[i]
				arr3[i] = arr3[j]
				arr3[j] = temp			
			}
		}
	}
	fmt.Printf("sorted array : %v\n",arr3)
}

func PredefinesMethods() {
	fmt.Println("Hello Sorting...")	
	arr1 := []int{3, 76, 43, 8}
	arr2 := []int{12, 5, 67, 54, 7}
	arr3 := append(arr1, arr2...)
	fmt.Printf("merged arrays : %v\n",arr3)
	sort.Ints(arr3)
	fmt.Printf("sorted array : %v\n",arr3)
}