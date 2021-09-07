package manipulate

func ManiArr(){
	var a [5]int
	var b = [5]int{1,2,3,4,5}
	var c = [...]int{1,2,3,4,5}

	_, _, _ = a, b ,c

	// all of array or slice use append or copy only for manage
}

func ManiMap(){
	// map use make 
	m := make(map[string]int)
	// add 
	m["A"] = 1

	// delete
	delete(m, "A")

	// get 1
	value1 := m["A"]

	// get2 protect error  , value of map and haskey is boolen
	value, hasKey := m["A"]

	_, _, _ = value1, value, hasKey

}