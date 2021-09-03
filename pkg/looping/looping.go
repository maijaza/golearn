package looping

import (
    "fmt"
)

func Test01() {
    // short if 
    // can do login and condition in same line
    i := 10
    if i = i /2; i < 10 {
        fmt.Println(i)
    }
    fmt.Println("01")
}

func Test02(i int) {
    // switch notice
    // multiple command without {}
    // case no need brake
    switch i {
    case 0:
        fmt.Println(0)
        fmt.Println(1)
    case 2:
        fmt.Println(1)
        fmt.Println(0)
    default:
        fmt.Println(2)
    }
      
}

func Test03(){
    // switch list
    // fallthrough for next case
    i := 0
    switch i {
    case 0,1 :
        fmt.Println(0)
        fallthrough
    case 2:
        fmt.Println(1)
        fallthrough
    default:
        fmt.Println(2)
    }
}

func Test04(){
    switch{
    case 1 == 1 :
        fmt.Println(0)
    }
}

func Test05(){
    defer fmt.Println("Test05")
    i := 1
    for {
        defer fmt.Println("Test" , i)
       if( i < 10){
           i++
           fmt.Println(i)
       }else{
           break
       }
    }
}