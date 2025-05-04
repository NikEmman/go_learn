package main

import "fmt"

//pointers
func main4(){

	// i:=10
	//  create datatype pointer, which hold the memory location where i (&i) is stored i.e 0xc000980040
	// p :=&i
	//  *p ( star before the pointer variable "dereferences" the pointer, which means get the value the memory location holds. In this case the value of i)
	// fmt.Println(*p)
	//  changing the value of *p changes the i variable's value because they both refer to the value held in the same memory location
	// *p = 11
	// fmt.Printf("the values of both p: %v, and i: %v have changed",*p, i )

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\n The memory location of the thing1 array is :%p",&thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is :%v", result)
	fmt.Printf("\n The value of thing1 is: %v",thing1)

	var result2 [5]float64 = square2(&thing1)
	fmt.Printf("\nThe second result is :%v", result2)
	fmt.Printf("\n The value of thing1 is now changed: %v",thing1)



}

func square(thing2 [5]float64) [5]float64{
	//this prints thing2 memory address, which differs from thing1, as shown above, which means we double the memory usage
	fmt.Printf("\n The memory location of the thing1 array is :%p",&thing2)

	for i:=range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return thing2
}

//instead of the above square func, we could have it take a pointer as an argument, thus not creating a copy of thing1, thus be more memory efficient
func square2(thing2 *[5]float64) [5]float64{
	//this prints thing2 memory address, which differs from thing1, as shown above, which means we double the memory usage
	fmt.Printf("\n The memory location of the thing2 array in square2 is same as thing1 :%p",thing2)

	for i:=range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
}

// pointers are usefull when passing in large parameters, where you don't have to create copies of data when you call a function