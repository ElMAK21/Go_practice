package main

import "fmt"

func main() {
	age := 32 // regular variable

	var agePointer *int
	agePointer = &age
	fmt.Println("Age:", *agePointer)

	editAgetAdultYears(agePointer)

	fmt.Println(age)

}

func editAgetAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}
