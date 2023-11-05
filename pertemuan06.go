// struktur data
// array, slice

package main

import "fmt"

func main() {
	// array
	var fruits [3]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"

	fmt.Println("Fruit : ", fruits)
	fmt.Println("Second fruit : ", fruits[1])

	// slice
	fruits2 := []string{"Apple", "Banana", "Orange"}
	fruits2 = append(fruits2, "Grape")
	fmt.Println("Fruits2 : ", fruits2)
	fmt.Println("Last fruit2 is", fruits2[3])

	// map
	fruitPrices := map[string]int{
		"Apple" : 1000,
		"Banana" : 5000,
		"Orange" : 7000,
	}
	fmt.Println("Prices all of fruits are : ", fruitPrices)
	fmt.Println("Price of apple is : ", fruitPrices["Apple"])
}
