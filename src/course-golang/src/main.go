package main

import "fmt"

func normalFuncion(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnDoubleValue(a int) int {
	return a * 2
}

func returnMoreOneValue(a int) (c, d int) {
	return a, a * 2
}

func main() {
	// normalFuncion("A normal message")
	// tripleArgument(2, 3, "Triple argumentos")
	// value := returnDoubleValue(10)
	// fmt.Println(value)
	value1, value2 := returnMoreOneValue(2)
	// value1, _ := returnMoreOneValue(2)
	fmt.Println(value1, value2)
}

/* ==============================================
===================== Apuntes ===================
=============================================  */

// ==== Declaracion de constantes ==== //
// const pi1 float64 = 3.14
// const pi2 = 3.1415

// fmt.Println("pi1:", pi1)
// fmt.Println("pi2:", pi2)

// ==== Declaracion de variables enteras ==== //
// base := 12
// var altura int = 10
// var area int

// fmt.Println("Var enteras", base, altura, area)

// ==== Zero values ==== //
// var a int
// var b float64
// var c string
// var d bool

// fmt.Println("Zero values", a, b, c, d)

// const baseCuadrado = 10
// areaCuadrado := baseCuadrado * baseCuadrado
// fmt.Println("Area cuadrado:", areaCuadrado)

// ==== Operadores aritmeticos ==== //
// x := 10
// y := 50

// result := x + y
// fmt.Println("Suma:", result)

// result = y - x
// fmt.Println("Resta:", result)

// result = y * x
// fmt.Println("Mult:", result)

// result = y / x
// fmt.Println("Division:", result)

// result = y % x
// fmt.Println("Modulo:", result)

// ==== Incremental ====
// x++
// fmt.Println(x)

// ==== Decremental ====
// x--
// fmt.Println(x)

// ==== Numeros enteros ====
// int = Depende del OS (32 o 64 bits)
// int8 = 8 bits = -128 a 127
// int16 = 16bits = -2^15 a 2^15-1
// int32 = 32bits = -2^31 a 2^31-1
// int64 = 64bits = -2^63 a 2^63-1

// ==== Numeros enteros positivos ====
// uint = Depende del OS (32 o 64 bits)
// uint8 = 8 bits = 0 a 2^8-1
// uint16 = 16bits = 0 a 2^16-1
// uint32 = 32bits = 0 a 2^32-1
// uint64 = 64bits = 0 a 2^64-1

// ==== Numeros decimales ==== //
// float32 = 32bits = +/- 1.18e^-38 a +/- 3.4e^38
// float64 = 64bits = +/- 2.23e^-308 a +/- 1.8e^308

// ==== Textos y boleanos ==== //
// string = ""
// bool = true / false

// ==== Paquete FM ==== //
// helloMessage := "Hello"
// worldMessage := "World"

// Println
// fmt.Println(helloMessage, worldMessage)
// fmt.Println(helloMessage, worldMessage)

// Printf
// nombre := "Platzi"
// cursos := 500
// siempre se recomienda colocar el tipo de dato
// fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
// fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

// Sprintf
// message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
// fmt.Println(message)

// Imprime el tipo de dato
// fmt.Printf("HelloMessage: %T\n", helloMessage)
// fmt.Printf("Cantidad cursos: %T\n", cursos)
