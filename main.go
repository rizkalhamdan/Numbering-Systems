package main

// import "RIZK/numbering-systems/utils"

import (
	"fmt"
	"RIZK/numbering-systems/decimal"
	"RIZK/numbering-systems/utils"
	"RIZK/numbering-systems/system"
)

func main() {
	decimal.Test()


	var rr []int

	rr = append(rr, 1, 1, 0, 0, 0)
	tyh := system.BinToDec(rr)

	println(tyh)

	var choise int
	var pri string
	var res []int

	for {
		// utils.DO("clear")
		res = res[:0]

		// welcome
		utils.Welcome()

		// get type id
		utils.PrintTypes()
		println("Your choise: ")

		choise = 0
		fmt.Scan(&choise)

		// get the number
		println("Enter the number: ")
		fmt.Scan(&pri)
		res = utils.GetRealNumber(pri)
		


		// do switch type
		switch choise {
		case 2:

		}

	
	}
}