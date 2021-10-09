package main

import "fmt"

type Product struct {
	title string
	id string
	price int
}

func main() {
	//#1 origin slice
	//create string slice without define length
	hobbies := []string{"football", "basketball", "coding"}

	//#2   
	//print all data
	fmt.Println(hobbies)
	//print only the first one
	fmt.Println(hobbies[0])
	//print the secound and third
	fmt.Println(hobbies[1:])
	
	//#3
	//create a slice that contain 1st and 2nd elements
	littleHobbies := hobbies[:2]
	anotherLittleHobbies := [2]string{}
	anotherLittleHobbies[0] = hobbies[0]
	anotherLittleHobbies[1] = hobbies[1]
	fmt.Println(littleHobbies)
	fmt.Println(anotherLittleHobbies)

	//#4
	//Reslice #3 to contain secound and the last element from origin
	littleHobbies = littleHobbies[1:3]  //it will go back to origin slice of littleHobbies and go further to 3rd element
	fmt.Println(littleHobbies)

	//#5 create goal slice
	courseGoal := []string{"own 100 bitcoins","own 100 eth"}  

	//#6
	courseGoal = append(courseGoal,"own 100 sol")
	fmt.Println(courseGoal)

	productList := [] Product {{"New product1","new-prod-1",99},{"New product2","new-prod-2",199}}
	productList = append(productList,Product{"New product3","new-prod-3",1999})
	fmt.Println(productList)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.