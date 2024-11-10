package longcommonprefix

import (
	"fmt"
	"testing"
)

func TestLongCommonPrefix(t *testing.T){
	str := []string{"flower", "flow", "flew", "flown"}
	result := LongCommonPrefix(str)
	fmt.Printf("result is :%v  \n", result)
}

func TestLongCommonPrefix_negative(t *testing.T){
	str := []string{"lower", "slow", "law", "flown"}
	result := LongCommonPrefix(str)
	fmt.Printf("result is :%v  \n", result)
}

func TestLongCommonPrefix2(t *testing.T){
	str := []string{"flower", "flow", "flew", "flown"}
	result := LongCommonPrefix2(str)
	fmt.Printf("result is :%v  \n", result)
}