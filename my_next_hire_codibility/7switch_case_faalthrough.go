package mynexthirecodibility

import "fmt"

func SwitchFallthrough(){
	v := 15 % 4
	switch v {
	case 3 : fmt.Println(100)
	        fallthrough
	case 2 : fmt.Println(42)
	case 1 : fmt.Println(1)
			 fallthrough
	default : fmt.Println("default")
	}
}