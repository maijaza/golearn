package interfaces

import (
	"fmt"
)

//pointer and slic

func Ex01(){
   // make slice 5 len 5 cap
   v := make([]int,5)	

   // make slice 0 len 5 cap
   v2 := make([]int, 0 , 5)

    v[0] = 1
   fmt.Println(v, v2)

   // addpend for add slice 
   // if append when len = cap , cap will extend = cap + cap such as len2 cap2 => len3 cap4
   v2 = append(v2,1,2)
   fmt.Println(v, v2, cap(v2))
}

func Ex02(){
	o := []int{1,2,3,4,5,6,7,8,9,10}
	
	for i,v := range o {
		fmt.Printf("index : %d value: %d \n", i ,v)
		v = 20
	}
  
	//for index only
	for i := range o {
		_ = i
	}

	//for value only
	for _ ,v := range o {
		_ = v
	}

	fmt.Println(o)
}

// map

func Ex03(){

 
	// create or initial map use make
	m := make(map[string]string)

	//add or update use =
	m["11"] = "AA"
	m["bb"] = "CC"
	fmt.Println(m)

	//delete use delete(map,key)
	delete(m,"11")
	fmt.Println(m)

	for _, v := range m{
		fmt.Println(v)
	}

	// check if has key will return data, but not will return false
	data, ok := m["11"]
	if !ok {
	   fmt.Println("error", ok)	
	}else {
	   fmt.Println(data)
	}
	

}

// function variable
func Ex04(){
  
	   fn := func (x,y int) int {
		   return x + y
	   }

	   fmt.Println(fn(1,2))

	   f := TTT()
	   fmt.Println(f())
	   fmt.Println(f())
}

func TTT() func() int {
	i := 10
	return func() int {
		j := i
		i = i + 1
		return j
	}
}


// https://golang.org/doc/
// https://tour.golang.org/methods/9