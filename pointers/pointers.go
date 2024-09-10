package main

import "fmt"

// Use pointers to avoid to crate copies on expensive objects or to directly mutate values
func main() {
	age := 37 // regular variable

	agePointer := &age
	fmt.Println(agePointer)
	fmt.Println(*agePointer)
	fmt.Println("Age: ", age)
	(getAdultYears(agePointer))
	fmt.Println("Age value mutated: ", age)

}

func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18 // mutate
}
