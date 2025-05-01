package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "mixedCharsλξω"
	var indexed = myString[0]
	fmt.Printf("%v,%T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	var strSlice = []string{"l", "o", "c", "a", "l"}
	var catStr = ""
	//this way we are creating a new string on each iteration, which is inefficient
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("%v\n", catStr)
	// a better way to to use built-in string Builder method
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("%v\n", catStr2)
}
