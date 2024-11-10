package mynexthireamericanexpress

import "fmt"

func SortedArray() {
	arr := []int{1, 2, 4, 6, 8, 9}
	fmt.Printf("this is sorted array %v", arr)
	target := 5
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				fmt.Printf("The pair for target %v is %v-%v \n", target, arr[i], arr[j])
				break
			}
		}
	}
}