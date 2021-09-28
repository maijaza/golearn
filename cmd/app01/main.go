package main

import (
	"fmt"
	 
	"net/http"
	"encoding/json"
	"encoding/xml"
)

type Emp struct{
	Id int `json:"ID"`
	Name string `xml:"Name",json:"Name"`
	Group int `xml:"grp,attr"`
	Age int `json:"Age,omitempty"`
}

func main(){

	http.HandleFunc("/exam01",service1)
	http.HandleFunc("/exam02",service2)
	http.HandleFunc("/exam03",service3)

	http.ListenAndServe(":8800",nil)
}

func service1(w http.ResponseWriter, r *http.Request){
    e := &Emp{
		Id: 1,
		Name: "Name 1",
		Group: 99,
		Age: 30,
	}

    w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(e)
}

func service2(w http.ResponseWriter, r *http.Request){
    e := &Emp{
		Id: 1,
		Name: "Name 1",
		Group: 99,
		Age: 30,
	}

    w.Header().Set("Content-type","application/xml")
	xml.NewEncoder(w).Encode(e)
}

func service3(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	o := &Emp{}
	if err := decoder.Decode(&o); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println(o)

	// u, err := json.Marshal(o)
	// if err !=nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// }
    // fmt.Println(u)


}