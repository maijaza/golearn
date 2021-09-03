package iopkg

import (
	"fmt"
	"io"
	"strings"
)

func Ex01(){
   s := "Hello World"

   r := strings.NewReader(s)

   b := make([]byte, 8)
   for{
	  n, err := r.Read(b)
	  fmt.Println("String : " , n , string(b[:n]))
	  if err == io.EOF{
		  break
	  }
   }
   fmt.Println("End Read")
}