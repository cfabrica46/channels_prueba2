package main

import "fmt"

func main() {

	c := make(chan int)

	go contar(c)

	go imprimir(c)

	var fin string
	fmt.Scan(&fin)

	close(c)
}

func contar(c chan int) {

	var x int

	for {
		c <- x
		x++

		fmt.Println("-----------------------------------")
	}

}

func imprimir(c chan int) {

	var valor int

	for {

		valor = <-c
		fmt.Println(valor)

	}

}
