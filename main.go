package main

import "fmt"

// generics


func mainEx() {
	var intSlice = []int{1,2,3}
	fmt.Println(sumIntSlice(intSlice))


	var float32Slice = []float32{1,2,3}
	fmt.Println(sumFloat32Slice(float32Slice))
	
}

func sumIntSlice(slice []int) int{
	var sum int
	for  _, v :=range slice{
		sum+=v
	}
	return sum
}
func sumFloat32Slice(slice []float32) float32{
	var sum float32
	for  _, v :=range slice{
		sum+=v
	}
	return sum
}

// in the example above we have to write the same code for each data type (int,float32 etc)
//to avoid this, we can use instead a generic as argument to the sumSlice function
func main(){
	var intSlice = []int{1,2,3}
	// the proper syntax is sumSlice[argumentType](argument)
	//in this case the [argumentType] can be omitted, cuz the compiler is smart enough, but if it was diff types of structs, it would need the [argumentType]
	fmt.Println(sumSlice[int](intSlice))


	var float32Slice = []float32{1,2,3}
	fmt.Println(sumSlice(float32Slice))
	var float64Slice = []float64{}
	fmt.Println(sumSlice(float64Slice))

	fmt.Println(isEmpty(intSlice))

	fmt.Println(isEmpty(float64Slice))

}
func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _,v :=range slice{
		sum+=v
	}
	return sum
}
// apart for the generic like above where we specify the different types, we might not care what data type the parameter will be, so we can use the any keyword
func isEmpty[T any](slice []T)bool{
	return len(slice)==0
}