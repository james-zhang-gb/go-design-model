package main

import "fmt"

type isSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}


func getFactory(brand string)(isSportsFactory,error){
	if brand=="nike"{
		return &nikeFactory{},nil
	}
	if brand=="adidas"{
		return &adidaseFactory{},nil
	}
	return nil,fmt.Errorf("don't have %s's factory",brand)
}