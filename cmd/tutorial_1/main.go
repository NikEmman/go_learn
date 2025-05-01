package main

import "fmt"

func main() {
	intArr := [...]int32{1, 2, 3}
	fmt.Println((intArr))

	var intSlice []int32 = []int32{4, 5, 6}

	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("The length is %v with capacity %v \n", len(intSlice), cap(intSlice))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println((myMap))
	var myMap2 = map[string]uint8{"Adam": 28, "Sarah": 45}
	fmt.Println(myMap2)
	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v", name, age)
	}
	for i, v := range intArr {
		fmt.Printf("Index:%v, Value:%v", i, v)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
