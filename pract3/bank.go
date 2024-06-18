package main

import (
	"fmt"

	"example.com/bank/fileops"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----")
		panic("Cant contrinue sorry ")
	}
	fmt.Println("Bienvenido al banco de los NACOS")
	fmt.Println("Hablame bb", randomdata.PhoneNumber())
	for {

		presentOptions()

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Tienes un total de: ", accountBalance)
		case 2:
			fmt.Println("Cuanto quieres depositar pa? ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalido tiene que ser un valor mayor a 0")
				//return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("El balance fue actualizado a: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Println("Cuanto quieres sacar pa? ")
			var withdrawtAmount float64
			fmt.Scan(&withdrawtAmount)

			if withdrawtAmount <= 0 {
				fmt.Println("Invalido tiene que ser un valor mayor a 0")
				continue
			}

			if withdrawtAmount > accountBalance {
				fmt.Println("Cantidad invalida no puedes sacar lo que no tienes ")
				continue
			}

			accountBalance -= withdrawtAmount
			fmt.Println("El balance fue actualizado a: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Te la lavas ")
			fmt.Print("Gracias por escoger al naco banco ")
			return
			//break
		}

	}

}
