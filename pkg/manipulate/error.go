package manipulate

import (
	"fmt"
	"errors"
)

// normally return error msg
func Err1(i int) (bool, error){
	if i < 0 {
		return false, errors.New("Error!")
	}
	return true, nil
}

// custom error like this
// can custo error number, error code ,error value or anything
type argError struct{
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func Err2(i int) (bool, error){
	if i < 0 {
		return false, &argError{i, "ERROR!"}
	}
	return true, nil
}

// example for convert type error to struct for using field
func Err2_1(){
	_, er := Err2(-1)
	if e1, ok := er.(*argError); ok {
		fmt.Println(e1.arg)
		fmt.Println(e1.prob)
	}
}