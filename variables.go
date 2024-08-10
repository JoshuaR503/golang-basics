package main

import "fmt"

func variables() {


	/// Conditionals



	/// Maps
	myMap := map[string]string {
		"Colombia": "Bogota",
		"Venezuela": "Caracas",
		"Brazil": "Brasilia",
		"Chile": "Santiago",
	}

	fmt.Println("The capital of Colombia is: ", myMap["Colombia"])

	fmt.Println("The map of countries is: ", myMap)

	delete(myMap, "Colombia")

	fmt.Println("The map of countries is: ", myMap)

	/// Arrays
	var fruits = [4]string{"Banna", "Orange"}

	fmt.Println("Fruit: ", fruits[1]);

	countries := [3]string{"Peru", "Brazil", "Venezuela"}
	
	fmt.Println(countries)

	
	/// Ways to declare variables.
	var name string = "Joshua"

	var namePerson  = "Joshua"

	namePerson2 := "Joshua"

	var n int8 = 127

	fmt.Println("Name: " + name);

	fmt.Println("Name: " + namePerson);

	fmt.Println("Name: " + namePerson2);

	fmt.Println("Number: ", n);
}