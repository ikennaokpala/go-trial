package main

import (
	"github.com/albrow/negroni-json-recovery"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"bitbucket.org/globalfoodbook/api/controllers"
)

func main() {

	router := mux.NewRouter()
	posts := controllers.Posts{}

	router.HandleFunc("/posts", posts.Index).Methods("GET")

	n := negroni.New(negroni.NewLogger())
	n.Use(recovery.JSONRecovery(true))
	n.UseHandler(router)

	n.Run(":3421")
}
