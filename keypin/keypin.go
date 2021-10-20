//usr/bin/env go run $0 "$@"; exit
package main

import (
	"os"
	"fmt"
	"flag"
	"time"	
	"bufio"
	"regexp"
	"strings"
	"net/http"
	"io/ioutil"
)

/*===============[Global variables that all functions take data from]===============*/

//Colors:
var WHITE = "\033[0m"
var GREEN = "\033[32m"
var BLUE = "\033[36m"
var RED = "\033[31m"
var ORANGE = "\033[33m"

//Icons:
var SUCCESS = "\033[32m✔\033[0m"
var FAILED = "\033[31mx\033[0m"



//List & variables declaration:
type storage struct {

	payload string
	count int

	lst_rua []string
	lst_verbs []string
	lst_paths []string
}

//Store the list & variables:
func storage_declare() *storage {
	st := &storage{}

	st.payload = "" 
	st.count = 0

	st.lst_rua = []string{}
	st.lst_verbs = []string{}
	st.lst_paths = []string{}

	return st
}

//User arguments - [OPTIONS]
type options struct {

	url string
	path string
	method string
	delay string
	cacheT string
	output string
}

//UserOptions (Arguments):
func parse() *options{
	opt := &options{}

	flag.StringVar(&opt.url, "u", "", "URL to test on [Ex: \"https://www.target.com\"]")
	flag.StringVar(&opt.path, "p", "", "path to bypass [Ex: \"admin\"]")
	flag.StringVar(&opt.method, "m", "all", "HTTP method to use [Ex: \"GET\", \"POST\", \"HEAD\" ...]")
    flag.StringVar(&opt.delay, "d", "0", "Delay between requests")
	flag.StringVar(&opt.cacheT, "c", "false", "Cached response technique. KeyPin will use a random [param+value] at the end.\n [Ex: \"https://www.target.com/admin?qewnje=8542\"]")
    flag.StringVar(&opt.output, "o", "keypin.txt", "output result")
	flag.Parse()

	return opt
}

/*
===============[This is the main function where all other functions gets executed]===============
							Run KeyPin with it's linked functions:							   */
func main() {
	opt := parse()
	st := storage_declare()
	

	//Banner Design & option(parse) display ("-h, --help"):
	ShowBanner()
	flagUsage()


	//Check so everything is proberly configured by the user:
	if opt.url == "" {
		fmt.Println("Use: ./keypin -u <Domain> -p <Path> [OPTIONS] ...")
		fmt.Printf("View help menu for more options [ %v-h%v ]\n", ORANGE, WHITE)
		os.Exit(0)
	}

	//Checking that the url has a HTTP protocol:
	valid_url, _ := regexp.MatchString(".*://", opt.url)
	if valid_url == false {
		fmt.Println("HTTP protocol is needed for:", opt.url)
		os.Exit(0)
	}



	//Check if the path is given:
	if len(opt.path) > 0 {
		//Check if it starts with "/" otherwise add it:
		if opt.path[0:1] != "/" {
			opt.path = "/"+opt.path
		}

	//If path has length zero == nothing. Warn the user:
	}else {
		fmt.Printf("\n:%v: WARNING no \"path\" has been set.\n", FAILED)
		time.Sleep(2 * time.Second)
	}

	//Display the configuration to the user:
	rua(st)
	config(opt, st)
	


	// ====:[Starting KeyPin and it's bypass techniques]:====
	
	/* ( Verbs ) */
	if opt.method == "all" {
		//verbs(opt)
	}



	/* ( Protocols ) */

	/* ( Paths ) */

	/* ( Extensions ) */



	//Starting bypass request with given bypass techniques given:
	request(opt, st)

	fmt.Println("\n:: Process finished successfully.")
}


/*================================[Functions]================================*/

// [] Setting up custom help menu:
func flagUsage() {

	flag.Usage = func() {
   	fmt.Fprintf(os.Stderr, "\nUsage: ./keypin -t <domain> -p <path> [OPTIONS] ...\n\n");
	
	fmt.Println("Keypin tries to bypass forbidden, unauthorized and unathenticated directories inside a domain\nby using diffirent techniques that manupulating Headers, Methods, Verbs, Extensions etc.")
	
	flag.PrintDefaults()
	fmt.Println("")
	}

}


// [] Show the user configure that the tool will relay on:
func config(opt *options, st *storage) {


	//Information & user configure output:
	fmt.Print(strings.Repeat("_", 70),"\n")

	fmt.Printf(`
· url                  : %v
· Path:                : %v
· Method:              : %v
· Output:              : %v
· User-Agents:         : %v
`, opt.url, opt.path, opt.method, opt.output, len(st.lst_rua))

	fmt.Print(strings.Repeat("_", 70),"\n")

}


// [] Requesting the target:
func request(opt *options, st *storage) {

	client := http.Client{}

	//Configure request query from "x" function & send request:
	req, _ := http.NewRequest(opt.method, opt.url+opt.path, nil)
	
	//Setting up headers for the client ("Request query"):
	resp, req_err := client.Do(req)	

	if req_err != nil {
		fmt.Println("Error:", req_err)
		os.Exit(0)
	}

  	//Calculate the response length:
	contentSize, _ := ioutil.ReadAll(resp.Body)
	size := len(contentSize)
	verbose := fmt.Sprintf(":%v: %v%v - [ Mehod: %v | Status: %v | Size: %v | Payload: %v 🔒 ]\n", st.count, opt.url, opt.path, opt.method, resp.StatusCode, size, st.payload)
	
	//Output verbose information that came from the response:
	fmt.Print(verbose)
}

// [] User agent to list:
func rua(st *storage) {
	file_rua, _ := os.Open("db/rua.txt")

    //Adding all lines(user-agents) to a list - [Random-Agent]:
	scanner := bufio.NewScanner(file_rua)    
    for scanner.Scan() {
        st.lst_rua = append(st.lst_rua, scanner.Text())
    }
}


// [] HTTP protocol bypass: Ex: "PUT, GET, POST, GeT, PoSt"
func protocol() {
	fmt.Print("→ Protocol function running")
	wordlist := make([]int, 1, 2)

	fmt.Print("wordlist >>", wordlist[3])

	//fmt.Println("::", *url)

}

func verbs(opt *options, st *storage) {

	file_verb, _ := os.Open("db/path_bypass.txt")
	
	//Gather all verb methods and add them to list:
	scanner := bufio.NewScanner(file_verb)    
    for scanner.Scan() {
        st.lst_verbs = append(st.lst_verbs, scanner.Text())
    }
}



// [] Paths bypass: Ex "/./admin, ///admin///"
func paths(st *storage) {
	file_paths, _ := os.Open("db/path_bypass.txt")
	//Configure all path payloads with the given path:
	scanner := bufio.NewScanner(file_paths)
    for scanner.Scan() {
        st.lst_paths = append(st.lst_paths, scanner.Text())
    }
}

/*

//Extension bypass Ex: "admin?.css":
func extension() {

}

//Header bypass Ex: "Host: 127.1, X-Forwarded-For: 127.0.0.1"
func headers() {
	
}

//Cache posioning bypass (Cookie/Session or header needed) Ex: (1) "/admin - (not logged in) == 403 ", (2) "/admin?.css - (logged in) == 200[CACHED]"
func cachePosioning() {

}

//Port bypass Ex: "Host: www.example.com:80", "Host: www.example.com:8443 "
func ports() {

}
*/

// [] Banner design:
func ShowBanner() {
	fmt.Printf(`
                                   %v___
                                 ,/ __¨\%v
   __  __     ______     __  __  %v\\ \´\ \%v   __     __   __    
  /\ \/ /    /\  ___\   /\ \_\ \  %v\\_¨¨ /%v  /\ \   /\ '-.\ \   
  \ \  _'-.  \ \  __\   \ \___, \  %v¨”\\ \%v  \ \ \  \ \ \-.  \  
   \ \_\ \_\  \ \_____\  \/\_____\    %v\\ l%v  \ \_\  \ \_\\'\_\ 
    \/_/\/_/   \/_____/   \/_____/     %v\\ l%v  \/_/   \/_/ \/_/ 
                                        %v\\_l%v
                         Version: v1.0   %v¨¨¨%v
                         Author: Brumens
`, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE)
}