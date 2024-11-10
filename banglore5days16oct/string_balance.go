package banglore5days16oct

import "strings"

import (
	"fmt"
)

func StringBalance(inputStrings []string) []string {
	// "{[(}){})})]}"
	var result []string
	for _, inputString := range inputStrings {
		fmt.Printf("inputString is : %v\n", inputString)
		for _, inputBrace := range inputString {
			if inputBrace == '{' {
				inputString = strings.Replace(inputString, "{", "", 1)
				inputString = strings.Replace(inputString, "}", "", 1)
			} else if inputBrace == '(' {
				inputString = strings.Replace(inputString, "(", "", 1)
				inputString = strings.Replace(inputString, ")", "", 1)
			} else if inputBrace == '[' {
				inputString = strings.Replace(inputString, "[", "", 1)
				inputString = strings.Replace(inputString, "]", "", 1)
			}
		}
		if len(inputString) == 0 {
			fmt.Printf("The given string is :%v balanced string\n", inputString)
			result = append(result, "1")
		} else {
			fmt.Printf("The given string is :%v not a balanced string\n", inputString)
			result = append(result, "0")
		}
	}
	return result
}
