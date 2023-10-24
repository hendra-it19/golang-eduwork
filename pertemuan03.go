// type data
// 1. integer : bilangan bulat
// 2. float : bilangan desimal
// 3. string : type teks
// 4. boolean (bool) : type data false or true
// 5. byte : representasi data dalam bentuk byte
// 6. nil : menunjukkan ketiadaan nilai atau null atau kosong

package main

import "fmt"

func main(){
	// int
	umur := 23

	// string 
	nama := "Hendra"

	// float 
	suhu := 37.6

	// byte
	var binaryData byte = 0b11010101

	// bool
	isMentor := true

	// nil
	var noValue interface{}

	fmt.Println("Usia : ", umur)
	fmt.Println("Nama : ", nama)
	fmt.Println("Suhu : ", suhu)
	fmt.Println("Binary Data : ", binaryData)
	fmt.Println("Adalah Mentor : ", isMentor)
	fmt.Println("Nil value : ", noValue)
}