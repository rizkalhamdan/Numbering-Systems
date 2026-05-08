package main

// import "RIZK/numbering-systems/utils"

import (
	"fmt"
	"RIZK/numbering-systems/decimal"
	"RIZK/numbering-systems/utils"
	
)

func main() {
	decimal.Test()

	tryit := utils.GetRealNumber("7f")



	fmt.Println(tryit)


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
		


		// do switch type
		switch choise {
			case 1: {
				
				break;
			}

		}

	
	}
}