package interfaces

import "fmt"

// how to create struct
// 1. new(structName) => return pointer of struct
// 2. structName{}
// 3. structName{filed:value}

type AA interface {
	test()
}

type BB struct {
	Id int
}

type CC struct {}

func (a *BB) test(){
	fmt.Println("test interface")
}

func (a *CC) test(){
	fmt.Println("test interface")
}

func Q(){
	fmt.Println("Q")
}

func Z(i AA){
	fmt.Println(i)
}


// type assertion
func Ex05() {

	var i interface{} = "hello"
    
	// if assertion true datatype it not error
	s := i.(string)  
	fmt.Println(s)

	// if is not true datatype should use 2 parameter 
	f, ok := i.(byte)
	fmt.Println(f , ok)
}

// swith on type
func Ex06(i interface{}){
	switch i.(type) {
	 case string: 
	    fmt.Println("string")
	 case int: 
	    fmt.Println("int")
	default : 
	fmt.Printf("I don't know about type %T!\n", i)
	}

}