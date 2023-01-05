package main 

import "fmt"

func main(){
	var kg1, kg2, kg3, kg4 float64
	var pon1, pon2, pon3, pon4 float64
	var ons1, ons2, ons3, ons4 float64
	var gr1, gr2, gr3, gr4 float64

	fmt.Scan(&gr1)
	fmt.Scan(&gr2)
	fmt.Scan(&gr3)
	fmt.Scan(&gr4)

	kg1 = gr1 / 1000.0
	kg2 = gr2 / 1000.0
	kg3 = gr3 / 1000.0
	kg4 = gr4 / 1000.0

	pon1 = gr1 / 453.592
	pon2 = gr2 / 453.592
	pon3 = gr3 / 453.592
	pon4 = gr4 / 453.592

	ons1 = gr1 / 28.3495
	ons2 = gr2 / 28.3495
	ons3 = gr3 / 28.3495
	ons4 = gr4 / 28.3495

	fmt.Println(kg1, pon1, ons1)
	fmt.Println(kg2, pon2, ons2)
	fmt.Println(kg3, pon3, ons3)
	fmt.Println(kg4, pon4, ons4)


}