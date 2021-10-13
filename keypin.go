//usr/bin/env go run $0 "$@"; exit
package main

import (
	"os"
	"io"
	"fmt"
	"flag"	
	"bufio"
	"regexp"
	"strings"
	"net/http"
)



//Random user agent file & list:
var rua_file, rua_err = os.Open("/home/kali/project/keypin/db/rua.txt")
var lst_rua = make([]string,0)



//Run keypin with it's linked functions:
func main() {

	//Banner Design & option(parse) display ("-h, --help"):
	ShowBanner()
	
	flag.Usage = func() {
   	fmt.Fprintf(os.Stderr, "\nUsage: ./keypin -t <domain> -p <path> [OPTIONS] ...\n\n");
	
	fmt.Println(`
Keypin tries to bypass forbidden, unauthorized, unathenticated directories inside a domain.
It do so by using diffirent techniques by manupulating Headers, Methods, Verbs, Extensions etc.

	`)
	
	flag.PrintDefaults()
	fmt.Println("\n")
	}

	//UserOptions (Arguments):
	var url string;		flag.StringVar(&url, "u", "", "URL to test on [Ex: \"https://www.target.com\"]")
	var path string;	flag.StringVar(&path, "p", "/", "path to bypass [Ex: \"/admin\"]")
	var method string;	flag.StringVar(&method, "m", "GET", "HTTP method to use [Ex: \"GET\", \"POST\", \"HEAD\" ...]")
    var output string;	flag.StringVar(&output, "o", "keypin.txt", "output result")


	flag.Parse()

	
	//Checking that "url" & "path" is parsed otherwise exit:
	if path == "/" {
		fmt.Println("Use: ./keypin -u <Domain> -p <Path>... ")
		fmt.Print(":x: Path was not detected. Do you want to continue? [y/n]: ")

		var input_noPath string
    	fmt.Scanln(&input_noPath)
		if input_noPath != "y" && input_noPath != "Y" {
			os.Exit(0)
		}
	}
	

	//Checking that the url has a HTTP protocol:
	valid_url, _ := regexp.MatchString(`.*://`, url)
	if valid_url == false {
		fmt.Println("HTTP protocol is needed for:", url)
		os.Exit(0)
	}


	//Show the configured setup before starting:
	config(url, path, method, output)


	//Client configure & request command setup:
	client := http.Client{}
	req, _ := http.NewRequest(method, url+path, nil)


	//Starting the bypass process:
	for i := 1; i < 5; i++ {
		
		//If there is a newwork error:
		resp, req_err := client.Do(req)

		if req_err != nil {
			fmt.Println("Error:", req_err)
			os.Exit(0)
		}

		contentSize, _ := io.ReadAll(resp.Body)
		size := len(contentSize)

		fmt.Printf(":%v: %v%v - [ %v | %v | %v ]\n", i, url, path, method, resp.StatusCode, size)
	}

	fmt.Println("\n:: Process done.")
}


//Banner design:
func ShowBanner() {
	fmt.Println(`
				        ___
				      ,/ __¨\
	 __  __     ______     __  __ \\ \´\ \	 __     __   __    
	/\ \/ /    /\  ___\   /\ \_\ \ \\_¨¨ /  /\ \   /\ '-.\ \   
	\ \  _'-.  \ \  __\   \ \___, \ ¨”\\ \  \ \ \  \ \ \-.  \  
	 \ \_\ \_\  \ \_____\  \/\_____\   \\ l  \ \_\  \ \_\\'\_\ 
	  \/_/\/_/   \/_____/   \/_____/    \\ l  \/_/   \/_/ \/_/ 
					     \\_l
					      ¨¨¨
				Version: v1.0
				Author: Brumens
`)
}


func config(url, path, method, output string) {


	//Information & user configure output:
	fmt.Print(strings.Repeat("_", 78),"\n",
	"\n\r· url\t\t:\t", url,
	"\n\r· Path:\t\t:\t", path,
	"\n\r· Method:\t:\t", method,
	"\n\r· Output:\t:\t", output,
	"\n\r· User-Agents:\t:\t",len(lst_rua),
	"\n",
	strings.Repeat("_", 78),"\n\n")

}

//User agent to list:
func rua() {

	scanner := bufio.NewScanner(rua_file)
    
    //Adding all lines(user-agents) to a list:
    for scanner.Scan() {
        lst_rua = append(lst_rua, scanner.Text())
    }
}
/*

//HTTP protocol bypass:
func protocol() {
	fmt.Print("→ Protocol function running")
	wordlist := make([]int, 1, 2)

	fmt.Print("wordlist >>", wordlist[3])

	fmt.Println("::", *url)

}

//HTTP verb bypass:
func extension() {
	fmt.Println("::", *url)

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