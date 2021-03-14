// Package greeting : descrizione package greeting
package greeting

import (
	"fmt"
	quoteV3 "rsc.io/quote/v3"
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
	fmt.Println("len(quoteV3.Hello()): ", utf8.RuneCountInString(quoteV3.HelloV3()))
	fmt.Println("quoteV3: ", quoteV3.HelloV3())
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
