package main

import "fmt"

// variables y constantes

// asume por inferencia el tipo de la constante
const MiConstante = "Ten dollars ñandú" // no hay problemas con caracteres especiales

func main() {
	fmt.Println("Hola Mundo desde GO")

	// declaracion por inferencia
	var nombre string = "Gustavo"

	// declaracion rapido o corta
	nombre2 := "Angel"

	fmt.Println(nombre)
	fmt.Println(nombre2)

	fmt.Printf("El valor de mi constante es: %s", MiConstante)

}
