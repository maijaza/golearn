package gort

import (
	"sync"
	"fmt"
	"time"
)

type Demo struct{
	mu sync.Mutex
	v int
}
// need package sync
// mutex
// use for ascess variable at time = mutual exclusion
func Ex04(){
 
	c := Demo{}
	for i:=0; i<20; i++ {
     go c.TestMux(i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(c) 
}

func (d *Demo) TestMux(i int){
	//d.mu.Lock()
	fmt.Println("=====start======" , i) 
	if i % 2 == 0 {
		time.Sleep(time.Second)
	}
    d.v = i
	fmt.Println("====End======" , i) 
	//d.mu.Unlock()
}