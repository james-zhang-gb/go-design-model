package main

import (
	"fmt"
	"log"
)

func main() {
	nikeFactory, err := getFactory("nike")
	if err != nil {
		log.Fatal(err)
	}
	adidaseFactory, err := getFactory("adidas")
	if err != nil {
		log.Fatal(err)
	}

	nikeShirt := nikeFactory.makeShirt()
	nikeShirt.setSize(98)
	shoe := adidaseFactory.makeShoe()
	fmt.Printf("%+v", shoe)
	fmt.Printf("%+v", nikeShirt)

}
