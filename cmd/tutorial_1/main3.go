package main

import "fmt"

// structs and interfaces
// structs are custom data types, much like classes
type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
	//alternatively can be set just owner, not ownerInfo owner and can be called straight up like myEngine.name
}
type owner struct {
	name string
}
type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}
func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Printf("You can make it!\n")
	} else {
		fmt.Printf("You need to fuel up\n")
	}
}

type engine interface {
	milesLeft() uint8
}

// structs can have methods too
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func main3() {
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: owner{"Bob"}}
	// struct fields can also be assigned like (25,15, owner{"Bob"}) or myEngine.mg=25. fields can be anything, even other structs

	// structs can be anonymous and be defined straight up but are not reusable
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{45, 85}

	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	//we use the interface, can pass anything as an engine as long as it has a milesLeft method that returns a uint8 (which was defined in the interface)
	canMakeIt(myEngine, 50)
	var myEngine3 = electricEngine{25, 15}
	canMakeIt(myEngine3, 50)

}
