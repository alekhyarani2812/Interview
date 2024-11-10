package mynexthirecodibility

import "fmt"

func CallBulbSeries() {
	A := []int{3, 2, 1, 4 ,5}
	//A := []int{1, 3, 2, 5, 4} //return 5
	//A := []int{1,2,3,5,4}
	//A := []int{3,4,1,2,5}
	moments := BulbSeries(A)
	fmt.Printf("The Number of bulbs are lighting are : %v\n", moments)
}
func BulbSeries(A []int) int {
	var result int
	mx := 0
	for i:=0; i<len(A); i++ {
		mx = max(mx,A[i])
		if mx == i+1 {
			result++
		}
	}
	return result
}
