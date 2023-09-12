package main

import "fmt"


func main() {

	cards := deck {"Ace of Diamonds", newCard()}

	cards = append(cards, "Seis de Espadas")

	for i,card := range cards{
		fmt.Println(i,card)
	}


	

	
}


func newCard() string{
	return "Cinco de copas" 
}



//var card string = "As de espadas"