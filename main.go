package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
	"github.com/antage/eventsource"
	"log"
	"net/http"
	"time"
)

func postMessageHandler(w http.ResponseWriter, r *http.Request)  {
	msg := r.FormValue("msg")
	name := r.FormValue("name")
	sendMessage(name, msg)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("name")
	sendMessage("", fmt.Sprintf("add user: %s", username))
}

func leftUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	sendMessage("", fmt.Sprintf("left user: %s", username))
}

type Message struct {
	Name 	string `json:"name"`
	Msg 	string `json:"msg"`
}

var msgCh chan Message

func sendMessage(name, msg string) {
	msgCh <- Message{name, msg}
}

func processMsgCh(es eventsource.EventSource) {
	for msg := range msgCh {
		data, _ := json.Marshal(msg)
		es.SendEventMessage(string(data), "", string(time.Now().Nanosecond()))
	}
}

func main()  {
	msgCh = make(chan Message)
	es := eventsource.New(nil, nil)
	defer es.Close()

	go processMsgCh(es)

	mx := pat.New()
	mx.Post("/messages", postMessageHandler)
	mx.Handle("/stream", es)
	mx.Post("/users", addUserHandler)
	mx.Delete("/users", leftUserHandler)


	n := negroni.Classic()
	n.UseHandler(mx)

	log.Fatal(http.ListenAndServe(":3000", n))
}
