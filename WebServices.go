package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"reflect"
)

func main(){
	http.HandleFunc("/api/hello", request)
	http.ListenAndServe(":8080",nil)
}

func request(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")
	ageString := r.URL.Query().Get("age")
	p := newPersonal(name,ageString)

	if p == nil {
		w.Write([]byte("error"))
	} else {
		b, _ := json.Marshal(p);
		fmt.Println(string(b))
		w.Header().Set("Content-Type","application/json");
		w.Write(b);
	}
	val := reflect.ValueOf(p).Elem()
        fmt.Println(val.NumField())

}
