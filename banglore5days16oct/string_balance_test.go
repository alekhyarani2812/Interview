package banglore5days16oct

import (
	"fmt"
	"testing"
)

func TestStringBalance(t *testing.T) {
	stringArray := []string{"{[(}){})})]}", "(]{(()})", "{[({(]})})][", "(){}({[]})"}
	result := StringBalance(stringArray)
	fmt.Printf("The result is %v\n",result)
}