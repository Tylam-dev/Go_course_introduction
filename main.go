package main

import "fmt"

func main() {
	var array [4]int
	array[0] = 0
	array[1] = 1
	fmt.Println(array[3], len(array), cap(array))

	//Crear array poblado
	slice := []int{2, 4, 21, 5, 10, 3, 2, 5}
	fmt.Println(slice, len(slice), cap(slice))

	//Primer elemento inclusivo, ultimo no inclusivo
	fmt.Println(slice[0:1])

	// metodo append
	newSlice := []int{2, 6, 4, 8, 3, 1}
	newSlice = append(newSlice, 1, 2, 3)
	slice2 := []int{}
	// Si no se especifica el tamaño del array es dinamico (slice)
	var nuevoArray []int
	nuevoArray = append(nuevoArray, 2, 3)
	fmt.Println(nuevoArray)
	slice2 = append(slice2, newSlice...)

	//recorre array
	for _, i := range slice {
		fmt.Println(i)
	}
}
