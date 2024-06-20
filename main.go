package main

import (
	"fmt"
	pk "golang_introduccion/mypackage"
)

type car struct {
	brand string
	year  int
}

type pc struct {
	ram   int
	disk  int
	brand string
}

// Creacion de un "metodo" a un struct
func (myPC pc) ping() {
	fmt.Println(myPC.brand, "marca")
}

func (myPC *pc) duplicateRam() {
	myPC.ram = myPC.ram * 2
	fmt.Println(myPC.ram)
}

type noConfigurado struct {
	Numero   *int
	Caracter *string
}

func printPuntero(obj noConfigurado) {
	fmt.Println(obj.Caracter, obj.Numero)
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

	//Punteros

	a := 50
	b := &a

	fmt.Println(a)
	fmt.Printf("direccion de memoria %v", b)
	fmt.Println(*b)

	myPc := pc{ram: 16, disk: 200, brand: "HP"}

	fmt.Println(myPc)

	myPc.ping()

	myPc.duplicateRam()

	prueba := noConfigurado{}

	//Imprime contenido de struct como <nil>
	printPuntero(prueba)

	//Asignar dato a direccion de memoria del struct
	numeroAsignar := 21
	prueba.Numero = &numeroAsignar

	//Imprime el struct con direccion de memoria establecida
	printPuntero(prueba)

	//Imprime valor de memoria
	fmt.Println(*prueba.Numero)
}
