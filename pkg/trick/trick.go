package trick

import (
	"fmt"
)

func T01(){
	// 3 type of array
	x := []int{1,2,3}	  // slice  but canbe create with make too
	y := [3]int{1,2,3}    // fix array
	z := [...]int{1,2,3}  // complier count array size

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
  
    

}