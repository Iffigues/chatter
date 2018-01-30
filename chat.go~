package main

import(
	"net/http"
)

func init(){
	auth := Router.PathPrefix("/chat").Subrouter()
	auth.Path("/").Methods("GET").HandlerFunc(chat)
}

func chat(w http.ResponseWriter,r *http.Request){
	jointure(r,w,donne{},"layout.html","chat.html","chatform.html")
}
