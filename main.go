package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	//Declaracion de variables enteras

	base := 12

	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area de un cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println(areaCuadrado)

	//Declaracion de variables
	helloMessage := "Hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)

	// Tipo de dato
	fmt.Printf("nombre : %T\n", nombre)

	normalFunction("hola elam")
	tripleArgument(1, 2, "elam")

	value1, value2 := dobleReturn((2))

	fmt.Println("value1", value1)

	fmt.Println("value2", value2)

	//Ciclo for condicional

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// for while
	counter := 0

	for counter < 10 {
		counter++
		fmt.Println(counter)
	}

	//For forver
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
