package mynexthirecodibility

import "fmt"

const end = 4
func DeferFunc(){
	for i:=1; i<end; i++ {
		fmt.Printf("i is :%v ", i)
		defer PrintNum(i)
	}
}

func PrintNum(i int){
	fmt.Printf("i is :%v ", i)
}