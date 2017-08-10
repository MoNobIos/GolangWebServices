package main

import(
	"strconv"
	"fmt"
)

type Personal struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func newPersonal(_name string,_age string) *Personal {
	age,err := strconv.Atoi(_age);
	if err != nil{
		fmt.Println("Error")
		return nil
	}
	if (len(_name) == 0) || (len(_age) == 0){
		return nil
	}
	p := &Personal{_name,age}
	return p
}