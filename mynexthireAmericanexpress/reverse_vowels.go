package mynexthireamericanexpress

import (
	"fmt"
	"strings"
)

func ReverseVowels(inputString string) string {
	vowels := "aeiou"
	matchedVowels := []string{}
	inputString = strings.ToLower(inputString)
	for _, inputVal := range inputString {
		if strings.Contains(vowels, string(inputVal)) {
			matchedVowels = append(matchedVowels, string(inputVal))
		 	//fmt.Printf("matchedVowels :%v\n",matchedVowels)
		}
	}

	//prepare result
	reversedVowels := ReverseStr(matchedVowels)
	result := ""
	vowelIndex := 0
	fmt.Printf("reversedVowels :%v\n",reversedVowels)
	for i, inputValue := range inputString {
		singleChar := string(inputValue)
		if strings.Contains(vowels, singleChar) {
			//vowels
			fmt.Printf("input string index :%v, input string's singleChar :%v, vowelIndex: %v \n",i, singleChar, vowelIndex)
			result = result + reversedVowels[vowelIndex]
			vowelIndex++
			
		}else{
			//consonants
			result = result + singleChar
		}
	}
	
	return result
}




func ReverseStr(matchedVowels []string) []string{
	reverse := []string{}
	for i:=len(matchedVowels)-1; i>=0; i-- {
		reverse = append(reverse, matchedVowels[i])
	}
	//fmt.Printf("reverced vowels :%v\n",result)
	return reverse
}