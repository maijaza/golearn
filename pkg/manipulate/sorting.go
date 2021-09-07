package manipulate

import (
	"fmt"
	"sort"
)


func ExSort01(){
   //sort string
   s := []string{"Z","A","c"}
   sort.Strings(s)
   fmt.Println(s)

   //sort int
   i := []int{5,3,2,1,2,3,5,6}
   sort.Ints(i)
   fmt.Println(i) 

   fmt.Println("Sorted :" , sort.IntsAreSorted(i))

// sort struct
   p := []struct{
        ID int
		Name string
        }{
			{1,"A"},{2,"C"},{5,"Z"}, {4, "R"},
		}
	
    sort.Slice(p, func(i, j int) bool { return p[i].ID < p[j].ID})
	fmt.Println(p) 

    // search will return indexs
	c := sort.Search(len(i), func(q int) bool { return i[q] > 1})
	fmt.Println(c)
}