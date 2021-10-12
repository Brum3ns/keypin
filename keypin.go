//usr/bin/env go run $0 "$@"; exit
package main

import (
	"os"
	"fmt"
	"flag"	
	"bufio"
	"strings"
	//"net/http"
)



//User arguments - [User Input]:
var domain = flag.String("d", "", "Domain to test")
var path = flag.String("p", "", "Path to bypass")
var bypass = flag.String("b", "all", `Bypass technique to use`)
var header = flag.String("hc", "", "Header/cookie to add in requests")
var method = flag.String("m", "GET", "Method to use")
var agent = flag.String("a", "/home/kali/project/keypin/db/rua.txt", "user agent to use")
var output = flag.String("o", "", "Save output")

//Random user agent file & list:
var rua_file, rua_err = os.Open("/home/kali/project/keypin/db/rua.txt")
var lst_rua = make([]string,0)



//Run keypin with it's linked functions:
func main() {
	//Banner Design & option(parse) display ("-h, --help"):
	banner()
	flag.Parse()
	
	//Configure KeyPin:
	rua()

	config()




	//Enable user input args to all functions that will run in "main":



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


//Banner design:
func banner() {
	fmt.Println(`
	
	[BANNER]
	
- Version: v1.0
- Author: Brumens
	`)
}


func config() {

	fmt.Print("x")

	//Information & user configure output:
	fmt.Print(strings.Repeat("-", 60),"\n",
	"· Bypass: \n",
	"· Target: \n",
	"· Path: \n",
	"· Threads: \n",
	"· Timeout: \n",
	"· Delay: \n",
	"· Header: \n",
	"· Cookie: \n",
	"· User-Agents: [",len(lst_rua),"] Lines\n",
	strings.Repeat("-", 60),"\n")

}

//User agent to list:
func rua() {

	scanner := bufio.NewScanner(rua_file)
    
    //Adding all lines(user-agents) to a list:
    for scanner.Scan() {
        lst_rua = append(lst_rua, scanner.Text())
    }
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