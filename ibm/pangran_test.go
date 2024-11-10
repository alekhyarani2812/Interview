package ibm

import (
	"fmt"
	"testing"
)

func TestPangram(t *testing.T) {
	str := []string{"abcdefghj iklmnoprqstuv wxyz", "The brown fox ", "Hello alekhya", "Abcdefghijklmnopqrstuvwxyz", "dffdfdfdsfdsfseredsfdsfdsfdsfsfsdfdsdsfefdsfdsdsfdfs"}
	result := Pangram(str)
	fmt.Printf("result :%v\n",result)
}