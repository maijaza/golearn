package gort

import (
	"fmt"
)

func sum(i int,c chan int){
   s := 0
   for x := 0 ; x < i; x ++{
       s = s + x
   }
   c <- s
}

func Ex01(){
   // go keyword for go routien
   // channel for send and receive data
   // channel use with go rootine
   c := make(chan int)
   go sum(10, c)
 
   fmt.Println(<- c)
}

//channel can be buffer
func Ex02(){
   ch := make(chan int,2) // channel can use only 2 channel in same time 
   ch <- 10
   ch <- 20
   //ch <- 30	   // this error becouse 3 channel

}


// Ex03 range and close channel
// example with fibo

func fibo(n int, c chan int ){
   x, y := 0,1 

   for i := 0; i< n; i++{
	   <- c
	   x, y = y , x+y
   }
   close(c)  // close cannel that tell no more value
}

func Ex03(){
   c := make(chan int, 10)   
   go fibo(cap(c), c)
   for v := range c {
	   fmt.Println(v)
   }
   // https://tour.golang.org/concurrency/4
}


//**
//  size of channel effect per amount of channel
//  if size < amount then use go rootine


// Example for go tour
// package main

// import (
// 	"fmt"
// 	"golang.org/x/tour/tree"
// )

// // Walk walks the tree t sending all values
// // from the tree to the channel ch.
// func Walk2(t *tree.Tree, ch chan int) {
// 	if t != nil {
// 		Walk2(t.Left, ch)
// 		ch <- t.Value
// 		Walk2(t.Right, ch)
// 	}
// }

// func Walk(t *tree.Tree, ch chan int) {
// 	Walk2(t, ch)
// 	close(ch)
// }

// // Same determines whether the trees
// // t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool {

// 	var st, st2 []int
// 	ch1, ch2 := make(chan int), make(chan int)

// 	go Walk(t1, ch1)
// 	for c := range ch1 {
// 		st = append(st, c)
// 	}
// 	go Walk(t2, ch2)
// 	for c := range ch2 {
// 		st2 = append(st2, c)
// 	}

//     fmt.Println(st)
// 	fmt.Println(st2)
//     for i,v := range st {
// 	   if st2[i] != v {
// 	      return false
// 	   }
// 	}

// 	return true
// }

// func main() {

// 	ch := make(chan int)
// 	go Walk(tree.New(1), ch)
// 	for c := range ch {
// 		fmt.Print(c)
// 	}
// 	fmt.Println("")
// 	fmt.Print(Same(tree.New(2), tree.New(2)))
// }


