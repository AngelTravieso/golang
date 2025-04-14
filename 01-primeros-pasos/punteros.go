package main

import "fmt"

func Punteros() {
	color := "rojo"

	// imprimir direccion de memoria del puntero
	fmt.Printf("%v: %s", &color, color)
}
