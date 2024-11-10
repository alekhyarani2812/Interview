package ibm

import (
	"fmt"
	"strings"
)

func Pangram(inputStrArr []string) string {
	allAlphabets := "abcdefghijklmnopqrstuvwxyz"
	result := ""
	
	fmt.Printf("allAlphabets :%v\n", string(allAlphabets))
	for _, inputStr := range inputStrArr {
		isPangram := true
		inputStr = strings.ToLower(inputStr)
		for _, letter := range allAlphabets {
			if !(strings.Contains(inputStr, string(letter))){
						isPangram = false			
			} 
		}

		if isPangram {
			result = result + "1"
		} else {
			result = result + "0"
		}

	}
	return result
}
