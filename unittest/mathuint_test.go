package unittest

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T){
	result := Add(10,20)
	fmt.Printf("addition of 2 nos :%v",result)
	expected := 40
	if result != expected {
		t.Errorf("Add(3, 5) = %d; want %d", result, expected)
	} 
}
func TestSub(t *testing.T){
	result := Sub(70,20)
	fmt.Printf("result sub of 2 nos :%v\n",result)

	expected := 40 
	if result != expected {
		t.Errorf("Sub(70, 20) = %d; want = %d\n", result, expected)
	}
}