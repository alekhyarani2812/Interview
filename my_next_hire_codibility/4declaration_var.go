package mynexthirecodibility

import "fmt"

func VariableDeclaration() {
	fmt.Println("Hello world....!")
	var a int
	a, b := fun()
	fmt.Printf("a val : %v\n", a)
	fmt.Printf("b val : %v\n", b)
}
func fun() (int, int) {
	return 1, 2
}
