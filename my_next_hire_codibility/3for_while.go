package mynexthirecodibility

import "fmt"

func ForWhile(){
	fmt.Print("copies of array...\n")
	i := 0
	for i < 5 {
		fmt.Printf("i...%v\n",i)
		i += 1
	}
}