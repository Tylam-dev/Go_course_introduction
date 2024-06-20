package main

import (
	"fmt"
	pk "golang_introduccion/mypackage"
)

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	//Otra metodo de creacion
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

	var newCar pk.CarPublic
	newCar.Brand = "Ferrari"
	newCar.Year = 2024

	fmt.Println(newCar)

	//No se puede establaces debido a que es privado
	// var secondNewCar carPrivate

	//Los metodos y types con mayusculas son Publicos
	pk.PrintMessage()
}
