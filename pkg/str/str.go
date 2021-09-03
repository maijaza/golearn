package str

import ( 
	"fmt" 
	"time"
)

type Person struct {
	name string
}

//Stringer for describe itself
func Ex01(){
	p := Person { name: "AAA"}
    fmt.Println(p) // with overide with Stringer
}

//stringer func overide Println whit stringer function 
func (p Person) String() string {
	//fmt.Println("String func")
	return "AA"
}

// error type can overide like stringer
// ==========================
type MyError struct{
	When time.Time
	What string
}


func (m MyError) Error() string {
	return fmt.Sprintf("$s %s", m.When, m.What)
}

func testErr() error{
	return &MyError{
		When: time.Now(),
		What: "TEST",
	}
}

func Ex02Err() {
   if err := testErr(); err != nil {
	   fmt.Println(err)
   } 

}
// ==========================


// Example error 2
// implement error func
type SqrError float64

func (s SqrError) Error() string {
	return fmt.Sprintf("Error : %v", float64(s))
}

func Ex03Err(i float64) (float64, error){

	if i < 0 {
       return i, SqrError(i)
	}
    return i, nil
}
// ========End Example error 2==================
