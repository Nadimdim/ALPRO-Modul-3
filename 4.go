package main

import "fmt"

func main(){
	var tahun int
	var kabisat1, kabisat2, kabisat3, kabisat4 bool

	fmt.Scan(&tahun)
	kabisat1 = (tahun % 400 == 0 || (tahun % 4 == 0 && tahun % 100 != 0))
	fmt.Scan(&tahun)
	kabisat2 = (tahun % 400 == 0 || (tahun % 4 == 0 && tahun % 100 != 0))
	fmt.Scan(&tahun)
	kabisat3 = (tahun % 400 == 0 || (tahun % 4 == 0 && tahun % 100 != 0))
	fmt.Scan(&tahun)
	kabisat4 = (tahun % 400 == 0 || (tahun % 4 == 0 && tahun % 100 != 0))

	fmt.Println("Apakah semuanya tahun kabisat?:", kabisat1 && kabisat2 && kabisat3 && kabisat4)
}