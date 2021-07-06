package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
)

func handleHtmlFile(w http.ResponseWriter, r *http.Request) {
	temp,err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln("parse html error:",err)
	}
	temp.Execute(w,nil)
}

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received message from client: " + reply)

		msg := "Received:  " + reply
		fmt.Println("Sending message to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send message")
			break
		}
	}
}

func main() {

	http.HandleFunc("/",handleHtmlFile)
	http.Handle("/ws",websocket.Handler(Echo))
	log.Println("开始监听8080端口")
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatalln("listen error:",err)
	}
}
