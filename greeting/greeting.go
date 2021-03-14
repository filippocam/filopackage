// Package greeting : descrizione package greeting
package greeting

import (
	"fmt"
	"rsc.io/quote"
	"unicode/utf8"
)

// MyConstant fsdsfd
const MyConstant = 5

// Hello   daasd
// 	es.
// 		greeting.Hello()
//			restituisce la stringa Hello!
func Hello() {
	test()
	fmt.Println("Hello!")
	fmt.Println("quote.Hello(): ", quote.Hello())
	fmt.Println("len(quote.Hello()): ", utf8.RuneCountInString(quote.Hello()))
}

//Prova dsfsfds
func Prova() {
	fmt.Println("Provaaaaaa!")
}

//Prova dsfsfds
func Prova2() {
	fmt.Println("Provaaaaaa22222!")
}

// Hi dfsdfs
func Hi() {
	test()
	fmt.Println("Hi!")
}

// test dfsdfs
func test() {
	fmt.Println("test!")
}
