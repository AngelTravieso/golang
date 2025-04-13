package main

import "fmt"

func VariablesConstantes() {
	fmt.Println("Hola Mundo desde GO")

	// declaracion por inferencia
	var nombre string = "Gustavo"

	// declaracion rapido o corta
	nombre2 := "Angel"

	fmt.Println(nombre)
	fmt.Println(nombre2)

	fmt.Printf("El valor de mi constante es: %s \n", MiConstante)
}
