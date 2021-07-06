package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	 temp, err := template.ParseFiles("index.html")
	 if err != nil {
	 	log.Fatalln("Some wrong happen in parsing template")
	 }


	 err = temp.Execute(w,nil)
	if err != nil {
		log.Fatalln("Some wrong happen in rendering template")
	}
}

type AskKey struct {
	K1 string
	K2 string
}

func handleAjaxReq(w http.ResponseWriter, r *http.Request) {
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln("read http request body panic")
	}

	var key AskKey
	err = json.Unmarshal(result, &key)
	if err != nil {
		log.Fatalln("askkey unmarsh panic")
	}

	w.Write([]byte(key.K1 + key.K2 + "end"))
}

func main() {
	http.HandleFunc("/",handleRoot)
	http.HandleFunc("/getContent",handleAjaxReq)
	err := http.ListenAndServe(":8090",nil)
	if err != nil {
		log.Fatalln(err)
	}
}
