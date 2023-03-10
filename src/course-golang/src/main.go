package main

import (
	"fmt"
)

func main() {

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

// ===== Funciones ===== //
// func normalFuncion(message string) {
// 	fmt.Println(message)
// }

// func tripleArgument(a, b int, c string) {
// 	fmt.Println(a, b, c)
// }

// func returnDoubleValue(a int) int {
// 	return a * 2
// }

// func returnMoreOneValue(a int) (c, d int) {
// 	return a, a * 2
// }

// normalFuncion("A normal message")
// tripleArgument(2, 3, "Triple argumentos")
// value := returnDoubleValue(10)
// fmt.Println(value)
// value1, value2 := returnMoreOneValue(2)
// value1, _ := returnMoreOneValue(2)
// fmt.Println(value1, value2)

// ===== Ciclo For ===== //
// for conditional
// for i := 0; i <= 10; i++ {
// 	fmt.Println(i)
// }

// fmt.Printf("\n")

// for while
// counter := 0
// for counter < 10 {
// 	fmt.Println(counter)
// 	counter++
// }

// for forever
// counterForever := 0
// for {
// 	fmt.Println(counterForever)
// 	counterForever++
// }

// ===== Condicionales ===== //
// valor1 := 1
// valor2 := 2

// if valor1 == 1 {
// 	fmt.Println("Es 1")
// } else {
// 	fmt.Println("No es 1")
// }

// with AND
// if valor1 == 1 && valor2 == 2 {
// 	fmt.Println("Es verdad AND")
// } else {
// 	fmt.Println("No es verdad AND")
// }

// with OR
// if valor1 == 0 || valor2 == 2 {
// 	fmt.Println("Es verdad OR")
// }

// convertir texto a numero
// value, err := strconv.Atoi("fads")
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(value, err)

// modulo := 4 % 2

// switch modulo {
// case 0:
// 	fmt.Println("Es par")
// default:
// 	fmt.Println("Es impar")
// }

// switch modulo2 := 4 % 2; modulo2 {
// case 0:
// 	fmt.Println("Es par")
// default:
// 	fmt.Println("Es impar")
// }

// value2 := 200
// switch {
// case value2 > 100:
// 	fmt.Println("Es mayor a 100")
// case value2 < 0:
// 	fmt.Println("Es menor a 0")
// default:
// 	fmt.Println("No condicion")
// }

// defer
// defer fmt.Println("Hola")
// fmt.Println("Mundo")

// // continue y break
// for i := 0; i < 10; i++ {
// 	fmt.Println(i)

// 	if i == 2 {
// 		fmt.Println("Es 2")
// 		continue
// 	}

// 	if i == 8 {
// 		fmt.Println("Break")
// 		break
// 	}
// }

// ==== Array y Slices ===== //
// arrays
// var array [4]int
// array[0] = 1
// array[1] = 2
// fmt.Println(array, len(array), cap(array))

// slices
// slice := []int{0, 1, 2, 3, 4, 5, 6}
// fmt.Println(slice, len(slice), cap(slice))

// metodos en el slice
// fmt.Println(slice[0])
// fmt.Println(slice[:3])
// fmt.Println(slice[2:5])
// fmt.Println(slice[4:])

// append
// slice = append(slice, 32)
// fmt.Println(slice)

// append nueva lista
// newSlice := []int{9, 13}
// slice = append(slice, newSlice...)
// fmt.Println(slice)
