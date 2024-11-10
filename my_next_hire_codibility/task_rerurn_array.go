package mynexthirecodibility

import "fmt"

func ReturnArray(){
	fmt.Println("Hello main ....!")
	n := 5
	res := CallingRetuernArray(n)	
	fmt.Printf("result is : %v\n",res)
}
func CallingRetuernArray(n int) []int {
	arr := []int{30,20,7,9,13}
	var result []int
	for i:=0; i<len(arr); i++ {
		result = append(arr, i)
	}
	return result
}