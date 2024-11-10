package komalroud2ltimindtree

import (
	"fmt"
	"strings"
	//"strings"
)

func FindDuplicateString() {
	strMap := make(map[string]bool)
	dublicateArray := []string{}
	uniqueArray := []string{}

	stmt1 := "Alekhya lives in Hyderabad"
	stmt2 := "Satya lives in Banglore"
	sentence := stmt1 + " " + stmt2
	fmt.Printf("sentence :%v\n",sentence)
	stringArray := strings.Split(sentence, " ")

	for _,stringArrayElement := range stringArray {
		if strMap[stringArrayElement] {
			dublicateArray = append(dublicateArray, stringArrayElement)
		} else if !(strMap[stringArrayElement]) {
			strMap[stringArrayElement] = true
			uniqueArray = append(uniqueArray, stringArrayElement)
		}
	}
	fmt.Printf("dublicate array :%v",dublicateArray)
	fmt.Printf("uniqueArray :%v",uniqueArray)
}

// func main() {
// 	fmt.Println("Hello, 世界")
// 	//simple go rouine
// 	go SampleGoroutine()
// }

// func SampleGoroutine() {
// 	fmt.Println("It's a SampleGoroutine")
// }

// Alekhya resides in Hyderabad
// Karteek lives in Banglore

// func FillWordsInMapCalling() {
// 	fmt.Println("Hello")
// 	wordMap := make(map[string]int)
// 	statement1 := "Alekhya resides in Hyderabad"
// 	statement2 := "Karteek lives in Banglore"
// 	mergeStatements := statement1 + " " + statement2
// 	//split by space
// 	stmtArr1 := strings.Split(mergeStatements, " ")
// 	resultMap := FillWordsInMap(wordMap, stmtArr1)
// 	fmt.Printf(" resultMap : %v \n", resultMap)

// }

// func FillWordsInMap(wordMap map[string]int, arr []string) map[string]int {
// 	for _, value := range arr {
// 		fmt.Printf(" value: %v \n", value)
// 		wordCount, exists := wordMap[value]
// 		if exists {
// 			//it's duplicate
// 			wordCount = wordCount + 1
// 			fmt.Printf(" duplicate word : %v \n", wordCount)
// 		} else {
// 			wordMap[value] = 1
// 			//fmt.Printf(" fill in map  : %v \n", value)
// 		}

// 	}
// 	return wordMap
// }




// You can edit this code!
// Click here and start typing.
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // Alekhya resides in Hyderabad
// // Karteek lives in Banglore

// func main() {
// 	fmt.Println("Hello, 世界")
// 	wordMap := make(map[string]int)
// 	statement1 := "Alekhya resides in Hyderabad"
// 	statement2 := "Karteek lives in Banglore"
// 	mergeStatements := statement1 + " " + statement2
// 	//split by space
// 	stmtArr1 := strings.Split(mergeStatements, " ")
// 	resultMap := FillWordsInMap(wordMap, stmtArr1)
// 	//fmt.Printf(" resultMap : %v \n", resultMap)

// 	//print unique words
// 	for key, value := range resultMap{
// 		if value == 1 {
// 			//fmt.Printf("The unique words are %v \n", key)
// 		}else {
// 			fmt.Printf("The common words are %v \n", key)
// 		}

// 	}

// }

// func FillWordsInMap(wordMap map[string]int, arr []string) map[string]int {
// 	for _, value := range arr {
// 		fmt.Printf(" value: %v \n", value)
// 		wordCount, exists := wordMap[value]
// 		if exists {
// 			//it's duplicate
// 			wordCount = wordCount + 1
// 			wordMap[value] = wordCount
// 			//fmt.Printf(" duplicate word : %v \n", value)
// 		} else {
// 			wordMap[value] = 1
// 			//fmt.Printf(" fill in map  : %v \n", value)
// 		}

// 	}
// 	return wordMap
// }

// func main() {
// 	fmt.Println("Hello world.!")
// 	stmt1 := "Alekhya resides in Hyderabad"
// 	stmt2 := "Karteek lives in Banglore"
// 	words := stmt1 +" "+ stmt2
// 	fmt.Printf("words : %v\n",words)
// 	wordCountMap := FindUniqueDublicate(words)
// 	//map[Alekhya:1 Banglore:1 Hyderabad:1 Karteek:1 in:2 lives:1 resides:1]
// 	for keyWord, valueCount := range wordCountMap {
// 		if valueCount > 1 {
// 			fmt.Printf(" %v is dublicate\n",keyWord)
// 		} else {
// 			fmt.Printf(" %v is unique\n",keyWord)
// 		}
// 	}

// }
// func FindUniqueDublicate(words string) map[string]int{
// 	wordArr := strings.Split(words, " ")
// 	wordCountMap := make(map[string]int) //{Alekhya : 1}
// 	for _,word := range wordArr {
// 		//wordCountMap[word]++
// 		_, exists := wordCountMap[word]
// 		if !exists {
// 			wordCountMap[word] = 1
// 		}else{
// 			wordCountMap[word] += 1
// 		}
// 		//fmt.Printf("wordCountMap : %v\n", wordCountMap)
// 	}
// 	return wordCountMap
// }
