//usr/bin/env go run $0 "$@"; exit
package main

import (
	"fmt"
	"flag"
	//"net/http"
)

//User arguments - [User Input]:
var domain = flag.String("domain", "xx", "Target to test on")


func main() {
	var banner = fmt.Println(`
		[BANNER]
		`)

	//Enable user input args to all functions that will run in "main":
	flag.Parse()

	protocol()



	//Setting up the request client:
	/*
	client := http.Client{}
	req, err := http.NewRequest("GET", domain, nil)


	//If there is a newwork error:
	if err != nil {
		fmt.Println("error:", err)
	}

	/*
	for i := 0; i < 10; i++ {
		fmt.Printf(":%v:\n", i)
	}
	*/

	fmt.Println(":: Process done.")
}

//HTTP protocol bypass:
func protocol() {
	fmt.Print("→ Protocol function running")
	wordlist := make([]int, 1, 2)

	fmt.Print("wordlist >>", wordlist[3])

	fmt.Println("::", *domain)

}

//HTTP verb bypass:
func extension() {
	fmt.Println("::", *domain)

}

/*

//Paths bypass:
func paths(err) {

}

//Extension bypass
func extension() {

}

//Header bypass
func headers() {
	
}

//Cache posioning bypass
func cache() {

}

//Port bypass
func ports() {

}
*/