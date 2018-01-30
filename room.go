package main

import (
	"net/http"
	"fmt"
)

func init(){
	auth := Router.PathPrefix("/room").Subrouter()
	auth.Path("/create").Methods("POST").HandlerFunc(createRoom)
	auth.Path("/search").Methods("POST").HandlerFunc(searchRoom)
	auth.Path("/connect").Methods("POST").HandlerFunc(connectRoom)
}
type ail struct{
     Ai bool
     Nbr int
}




type typeRoom struct{
	Name string `json:"name"`
	Private string `json:"password"`
}



func createRoom(w http.ResponseWriter,r *http.Request){
	var(
		ali typeRoom
	)
	if !isRegister(r){
		http.Error(w,"bad",405)
		return
	}
	grap(r,&ali)
	fmt.Println(ali)
	name := ali.Name
	private := ali.Private
	var new *froom = new(froom)
	var mimi = hub{
		broadcast:  make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
		clients:    make(map[*client]bool),
		content:    nil,
		live :      false,
	}	
	new.room = mimi
	session, _ := store.Get(r, "session-name")
	nnn,_ := session.Values["user"].(string)
	new.user = append(new.user, userws{nnn,true,false})
	if name != ""{
		if private == ""{
			new.private = false
		}else{
			new.private=true
			new.password=private
		}
		if _,ok := stand[name];ok{
			http.Error(w,"bad",403)
			return
		}else{
			stand[name]= new
			go stand[name].run()
			sendRes(http.StatusAccepted,"haha",w)
		}
	}
}

type Search struct{
	Search string `json:"search"`
}

type sendRoom struct{
	Private bool `json:"private"`
	Name string `json:"name"`
	Number int `json:"number"`
}

func searchRoom(w http.ResponseWriter,r *http.Request){
	if !isRegister(r){
		return
	}
	var (
		b Search
		c []sendRoom
	)
	grap(r,&b)
	if(b.Search == ""){
		for x,y := range stand{
			c = append(c,sendRoom{y.private,x,len(y.user)})
		}
	}else{
	}
	sendJson(c,w)
}

func isInRoom(a *froom,b string){
	for i,ok := range a.user{
		if ok.name == b {
			a.user[i].connect = true
			return
		}
	}
	a.user = append(a.user,userws{b,true,false})
}


func connectRoom(w http.ResponseWriter,r *http.Request){
     	var (
	    sal typeRoom
	)
	if !isRegister(r){
		return
	}
	grap(r,&sal)
	session, _ := store.Get(r, "session-name")	
	if _,ok := stand[sal.Name];!ok{
		return
	}

	if stand[sal.Name].password != sal.Private{
		return
	}
	nnn,_:= session.Values["user"].(string)
	isInRoom(stand[sal.Name],nnn)
	sendRes(http.StatusAccepted,"ez",w)  
}
