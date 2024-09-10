package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// var productNames [4]string = [4]string{"a book"}
	// productNames[2] = "a third value"
	// prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	// fmt.Printf("prices: %v\n", prices)
	// fmt.Println(productNames)
	// fmt.Println(productNames[1:3])

	// // Dynamic arrays
	// dynamics := []float64{10.99, 2.99}
	// fmt.Println(dynamics[0:2])

	// // to reassign dynamics = append(dynamics, 5.99)
	// appended := append(dynamics, 5.99)
	// fmt.Println(appended)

	// Time to practice what you learned!

	//  1. Create a new array (!) that contains three hobbies you have
	hobbies := [3]string{"Sports", "Cooking", "Reading"}
	//     Output (print) that array in the command line.
	fmt.Println(hobbies)
	//  2. Also output more data about that array:
	//     - The first element (standalone)
	fmt.Println(hobbies[0])
	//     - The second and third element combined as a new list
	fmt.Println(hobbies[1:3]) // last index is excluded, 1: would select until end
	//  3. Create a slice based on the first element that contains
	//     the first and second elements.
	mainHobbies := hobbies[0:2]
	fmt.Println(mainHobbies)
	//     Create that slice in two different ways (i.e. create two slices in the end)
	//  4. Re-slice the slice from (3) and change it to contain the second
	//     and last element of the original array.
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)
	//  5. Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Go", "Basics"}
	fmt.Println(courseGoals)
	//  6. Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Learn all the details"
	courseGoals = append(courseGoals, "Learn all the gundas")
	fmt.Println(courseGoals)
	//  7. Bonus: Create a "Product" struct with title, id, price and create a
	//     dynamic list of products (at least 2 products).
	products := []Product{{"First product", "a First prusto", 12.99}, {"second product", "a second prusto", 121.99}}
	fmt.Println(products)
	newProduct := Product{
		"third",
		"product",
		11.99,
	}
	products = append(products, newProduct)
	fmt.Println(products)
	//     Then add a third product to the existing list of products.
}
