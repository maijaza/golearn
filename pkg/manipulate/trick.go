package manipulate

import (
	"sync"
	"fmt"
	"reflect"
	"time"
)

// multi value without length  we call Variadic Functions
func Multi(nums ...int){
	fmt.Println(nums)
	fmt.Println(reflect.TypeOf(nums))
}


func ExClosure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestClosure(){
   nextInt := ExClosure()
   fmt.Println(nextInt())
   fmt.Println(nextInt())
   fmt.Println(nextInt())
   nextInt = ExClosure()
   fmt.Println(nextInt())
}

//=========================
//test feature check type ofobject
// we can convert type like this  x.(int)
type person interface{
   walk() int	
}

type man struct{

}

type woman struct{

}

func (m man) walk() int{
	return 1
}

func (m woman) walk() int{
	return 2
}

func checkType(o person){
	switch o.(type) {
		case man : 
		   fmt.Println("MAN")
		case woman : 
		   fmt.Println("WOMAN")
	 }
}

func ExType(){
   m := man{}
   w := woman{}

   checkType(m)
   checkType(w)
}

//========================================
//// routine

func test(n int, s string){
	for i:=0; i< n; i++ {
		fmt.Println("loop ", s , ": ", i)
	}
}

// normal
func ExRou1(){
    test(5, "direct")
	go test(5 , "routine ")

	go func(msg string) {
        fmt.Println(msg)
    }("going")

	time.Sleep(time.Second)
}

// waiting group for wait for all route

func worker(i int, wg *sync.WaitGroup){
    defer wg.Done()
	fmt.Printf("Worker %d starting\n", i)
}

func ExRou2(){
   var wg sync.WaitGroup

   for i :=0; i < 10 ; i++{
     wg.Add(1)
	 go worker(i, &wg)
   }

   wg.Wait()

}


// channel Synchronization
func worker2(done chan bool){
	fmt.Println("workkkk")
	done <- true
}

// for blocking program waiting for change use <- xx
func ExRou3(){
	done := make(chan bool,1)
	go worker2(done)
	<- done
}

// ======================

// channel for sending only
func ping(ping chan<- string){
	ping <- "PING"
}

// ping channal for recieve and pong change for sending only 
func pong(ping <-chan string, pong chan<- string){
	s :=  <- ping
	pong <- s

}

func ppp(p chan string){
	p <- "A"
}

func ExRou4(){
    pings := make(chan string, 1)
	pongs := make(chan string, 1)
    ppps := make(chan string,1)
	ping(pings)
	pong(pings,pongs)
	ppp(ppps)
	fmt.Println(<-pongs)
	
	fmt.Println(<-ppps)
}

// select for channel
// select use for waiting channel
// use timeout with time.After for channel call external service
// user default for non-blocking channel, it will pass to default without waiting
func ExRou5(){
	c1 := make(chan string)
    c2 := make(chan string)
	go func() {
        time.Sleep(4 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(4 * time.Second)
        c2 <- "two"
    }()
	
	for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)

		// case <-time.After(3 * time.Second):
		// 	fmt.Println("timeout 2")
        }
    }
}