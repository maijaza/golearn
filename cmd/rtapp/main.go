package main

import (
	"golearn/pkg/manipulate"
	_"golearn/pkg/gort"
   "fmt"
)

func main(){
	//gort.Ex01()
   //gort.Ex03()

   //gort.Ex04()
  
   //manipulate.Multi(1,2,3)
   // manipulate.TestClosure()

   //manipulate.ExType()

   _, err := manipulate.Err1(-1)
   if err != nil {
      fmt.Println(err)
   }

   _, err = manipulate.Err2(-1)
   if err != nil {
      fmt.Println(err)
   }
   
   //manipulate.Err2_1()
   fmt.Println("=============================================")

   //manipulate.ExRou1()
   fmt.Println("=============================================")
   //manipulate.ExRou2()
   //manipulate.ExRou3()
   fmt.Println("=============================================")
   //manipulate.ExRou4()
   fmt.Println("=============================================")
   //manipulate.ExRou5()

   fmt.Println("=============================================")
   //manipulate.ExTimer()

   //manipulate.ExTicker()
   fmt.Println("=============================================")

   manipulate.ExSort01()
   manipulate.ExStr()
   fmt.Println("=============================================")
   manipulate.ExStr02()
   fmt.Println("=============================================")
   manipulate.ExStr03()
}