package main

import "fmt"

func main(){
	var a, b, c string
	var hasil bool

	fmt.Scanln(&a, &b, &c)
	hasil = a == b || a == c || c == b

	fmt.Println(hasil)
}