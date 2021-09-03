package main

import (
	"golearn/pkg/str"
	"unsafe"
	"fmt"
	_ "golearn/pkg/looping"
    "golearn/pkg/interfaces"
	"reflect"
)

func main(){
	defer fmt.Println("Hello World!")
 
	//looping.Test01()
    //line()
    //looping.Test02(2)
    //looping.Test05()
	// v := BB
	// BB.Test()
    interfaces.Q()
	v := interfaces.BB{}
	v2 := new(interfaces.CC)
	fmt.Printf("%T",v)
	fmt.Println(reflect.TypeOf(v2))
	//interfaces.Z(&v)
	//interfaces.Z(v2)
    var a int
	p := &a
	fmt.Printf("a: %T, %d\n", p, unsafe.Sizeof(p))
	line()
	interfaces.Ex04();
	
	line()
	interfaces.Ex05()

	line()
	interfaces.Ex06(&interfaces.BB{})

	line()
	str.Ex01()

	line()
	str.Ex02Err()

	line()
    if _, err := str.Ex03Err(10); err != nil {
		fmt.Println(err)
	}

}
   

func line(){
	fmt.Println("====================================")
}