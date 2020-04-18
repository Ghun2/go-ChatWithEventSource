package main

import (
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func postMessageHandler(w http.ResponseWriter, r *http.Request)  {
	msg := r.FormValue("msg")
	name := r.FormValue("name")
	log.Println("postMessageHandler", msg, name)
}

func main()  {
	mx := pat.New()

	mx.Post("/messages", postMessageHandler)

	n := negroni.Classic()
	n.UseHandler(mx)

	log.Fatal(http.ListenAndServe(":3000", n))
}
