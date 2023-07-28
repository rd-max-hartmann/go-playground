package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"undefined/apps/api/model"
	"undefined/libs/mylib"
)

func getHello(response http.ResponseWriter, request *http.Request) {
	res := model.Person{
		Name:    "Max Mustermann",
		Address: "Carlswerk 22"}
	res.Test()
	test := request.URL.Query().Get("test")
	fmt.Println(test)
	fmt.Println(mylib.Mylib("Max"))
	err := json.NewEncoder(response).Encode(res)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
}

func main() {
	http.HandleFunc("/hello", getHello)

	srv := &http.Server{Addr: ":8080"}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
