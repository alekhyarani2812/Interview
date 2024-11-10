package mynexthireamericanexpress

import (
	"fmt"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	inputStrArray := "Hello Alekhya" //output: Halle Alokhye
	// []string{"Hello Alekhya", "How are you", "What are you doing", "Good Bye"}
	result := ReverseVowels(inputStrArray)
	fmt.Printf("result is %v \n", result)
}
