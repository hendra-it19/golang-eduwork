// operator : aritmatika, logika, pembanding

package main

import "fmt"

func main() {

	a := 10
	b := 3

	c := false
	d := true

	fmt.Println("Penambahan =>  ", a, " + ", b, " = ", a+b)
	fmt.Println("Pengurangan ", a-b)
	fmt.Println("Perkalian ", a*b)
	fmt.Println("Pembagian ", a/b)
	fmt.Println("Sisa bagi ", a%b)

	fmt.Println("AND : ", c && d)
	fmt.Println("OR : ", c || d)
	fmt.Println("NOT : ",  !d)

	fmt.Println("Kurang Dari : ", a < b)
}
