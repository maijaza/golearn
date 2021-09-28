package main

import (
	"golearn/pkg/db"
	"fmt"
)
func main(){
    fmt.Println("")
	 


    // case
	//1. select 1 field
	fmt.Println("=========== 01 ===========")
	v := db.Select01()
    fmt.Println(v)


	//1.1 select 1 row with struct
	fmt.Println("=========== 1.1 ===========")
	v = db.Select01_1()
    fmt.Println(v)
	
	

	//1.2 select N row 
	fmt.Println("=========== 1.2 ===========")
	list, _ := db.Select01_2()
    fmt.Println(list)
	

	//1.3 select N result set
	fmt.Println("=========== 1.3 ===========")
	list, _ = db.Select01_3()
    fmt.Println(list)


	//2. insert
	fmt.Println("=========== 2 ===========")
    id, err := db.Insert()
	fmt.Println(id)
	fmt.Println(err)
	
	//3. update
	fmt.Println("\n=========== 3 ===========")
    id, err = db.Update()
	fmt.Println(id)
	fmt.Println(err)
	list, _ = db.Select01_3()
    fmt.Println(list)

	//4. delete
	fmt.Println("\n=========== 4 ===========")
    id, err = db.Delete()
	fmt.Println(id)
	fmt.Println(err)
	list, _ = db.Select01_3()
    fmt.Println(list)


	//5 transaction
	fmt.Println("\n=========== 5 ===========")
    id, err = db.ExecTrans()
	fmt.Println(id)
	fmt.Println(err)
	list, _ = db.Select01_3()
	fmt.Println(list)
	
	
	//6. store without return


	//7. store with return
	fmt.Println("\n=========== 7 ===========")
	data2, err := db.ExStore02()
	fmt.Println(data2)
	fmt.Println(err)

	defer db.Close()
}