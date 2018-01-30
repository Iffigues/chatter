package main

import(
	"net/http"
	"github.com/gorilla/context"
)
type key string

func isExiste(r *http.Request, name key)(vrai bool) {
	_, vrai = context.GetOk(r, name)
	return
}



func delKey(r *http.Request,name key){
	context.Delete(r,name)
}

func createName(r *http.Request,name key)(vrai bool){
	vrai = true
	if !isExiste(r,name){
		vrai = false
		return
	}
	var MyKey key = name
	context.Set(r,MyKey,"")
	return
}
