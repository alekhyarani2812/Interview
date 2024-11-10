package mynexthirecodibility

import "fmt"

type Cat struct {
	IsSleeping bool
}
func (c Cat) Sleep() {
	//c.IsSleeping = true
	fmt.Printf("bill is jkdbm\n")
}

func BilliCat(){
	fmt.Print("Billi cat program.....")
	billi := Cat{}
	billi.Sleep()
	fmt.Printf("billi : %v\n",billi.IsSleeping)
}