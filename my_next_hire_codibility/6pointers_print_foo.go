package mynexthirecodibility

import "fmt"

type S struct {
	m string
} 
func f() *S{
	return &S{"foo"}
}
func PrintFoo(){
	p := *f()
	fmt.Printf("print foo: %v\n", p.m)
}