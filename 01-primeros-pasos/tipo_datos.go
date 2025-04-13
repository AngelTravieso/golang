package main

import "fmt"

func TipoDatos() {
	var string1 string = "texto"
	textoGrande := `
		Adipisicing ea eiusmod quis culpa enim. Tempor eiusmod laborum quis et proident. Id irure labore adipisicing in adipisicing enim laborum. Velit qui nulla est ut tempor et incididunt non et ut culpa anim id.
	`
	fmt.Println(string1)
	fmt.Println(textoGrande)

	fmt.Println("--------------------------------------------------------------------------")

	var estado bool = true
	var flotante32 float32 = 32.33
	var flotante64 float64 = 32.33654654
	var entero int = 1234
	var entero_int8 int8 = 123 // -128 a 127
	var entero_int16 int16 = 123
	var entero_int32 int32 = 4556123
	var entero_int64 int64 = 65465
	var entero_uint uint = 65465
	var entero_uint8 uint8 = 71
	var entero_uint16 uint16 = 65465
	var entero_uint32 uint32 = 65465
	var entero_uint64 uint64 = 65465

	fmt.Println(estado)
	fmt.Println(flotante32)
	fmt.Println(flotante64)
	fmt.Println(entero)
	fmt.Println(entero_int8)
	fmt.Println(entero)
	fmt.Println(entero_int16)
	fmt.Println(entero_int32)
	fmt.Println(entero_int64)
	fmt.Println(entero_uint)
	fmt.Println(entero_uint8)
	fmt.Println(entero_uint16)
	fmt.Println(entero_uint16)
	fmt.Println(entero_uint32)
	fmt.Println(entero_uint64)

	// con usar int para enteros basta, los demas que sean para calculos grandes/matematicos

}
