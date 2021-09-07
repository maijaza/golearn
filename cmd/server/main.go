package main

import (
	"bytes"
	"bufio"
	"fmt"
	"net/http"
	"time"
)

func main(){

	//client()
    //runServer01()
	runServer02()
 
	http.ListenAndServe(":8800",nil)
} 

//url.QueryEscape(s) for encode param
func client(){

	// for fix timeout
	// c := http.Client { Timeout: time.Duration(1) * time.Second}
	
	// res, err := http.Get("https://www.google.co.th")	//for get
	// if err != nil {
	// 	panic(err)
	// }
 
	res, err := http.Post("https://httpbin.org/post","application/text",  bytes.NewBuffer([]byte{}))  // for post
		if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	fmt.Println("Response status : " , res.Status)

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err !=nil {
		panic(err)
	}
	
}

// for serve service with handleFunc
// *** if you use Handle => you should implement ServeHTTP in interface too
func hello(w http.ResponseWriter, req *http.Request){
	fmt.Fprint(w, "Hello!")
}

func header(w http.ResponseWriter, req *http.Request){
	for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func runServer01(){
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/header", header)
}

//  demo  handler with context
func hello2(w http.ResponseWriter,req *http.Request){

	ctx := req.Context();
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler Stopped")

	select {
		case <- time.After(10 * time.Second):
			fmt.Fprintf(w, "hello\n")
		case <- ctx.Done():
			err := ctx.Err()
			fmt.Println("server:", err)
    		internalError := http.StatusInternalServerError
    		http.Error(w, err.Error(), internalError)
		}
}

func runServer02(){
    http.HandleFunc("/hello", hello2)
}