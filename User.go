package main

import(
	"net/http"
)

func init(){
	//auth := Router.PathPrefix("/user").Subrouter()
	/*auth.Path("/login").Methods("POST").HandlerFunc(connection)
	auth.Path("/logout").Methods("GET").HandlerFunc(deconnection)
	auth.Path("/checkin").Methods("POST").HandlerFunc(inscription)
	auth.Path("/isConnected").Methods("GET").HandlerFunc(isConn)*/
}

type Namer struct {
	Name string `json:"Name"`
}


func isConn(w http.ResponseWriter,r *http.Request){
	if isRegister(r){
		sendRes(http.StatusAccepted,"accepted",w)
		return
	}
	http.Error(w, "Method ", 405)
}


func findUser(utilisateur string)(vrai bool){
	db := openDB()
	db.QueryRow("SELECT IF(COUNT(*),'true','false') FROM user WHERE login = ?", utilisateur).Scan(&vrai)
	return
}

func verifUser(names,password string)(io bool){
	io = false
	var pwd []byte
	ser, err := dbPrepare("SELECT password FROM user WHERE login=? LIMIT 1");
	writeError(err)
	res,err := dbQueryOne(ser,names)
	writeError(err)
	if res == nil{
		return 
	}
	for res.Next(){
		res.Scan(&pwd)
	}
	if decrypt(pwd,[]byte(password)){
		io = true
	}
	return
}

func connection(w http.ResponseWriter,r *http.Request,user string,password string)(vrai bool){
	vrai = false
	if isRegister(r){
		vrai = false
		return
	}
	if verifUser(user,password){
		register(w,r)
		session, _ := store.Get(r, "session-name")
		session.Values["name"] = user
		session.Values["co"] = true
		session.Save(r,w)
		vrai = true
		return
	}
	return
}

func deconnection(w http.ResponseWriter,r *http.Request)(vrai bool){
	vrai = true
	if isRegister(r){
		unregister(w,r)
		return
	}
	vrai = false
	return
}

func addUser(a User)(vrai bool){
	passwording,_ := crypte(a.Password)
	user, err := dbPrepare("INSERT INTO user(login,mail,password) VALUES(?,?,?)")
	writeError(err)
	_,err = dbExec(user,a.Login,a.Mail,passwording)
	vrai = !writeError(err)
	return 
}

type User struct{
	Login string `json:"login"`
	Mail string `json:"mail"`
	Password string `json:"password"`
	Repass string `json:"repass"`
}


func registreError(r User)(tt bool){
	tt = true
	if r.Password != r.Repass {
		tt = false
	}
	if r.Login  == ""|| r.Password  == "" || r.Repass == "" {
		tt = false
	}
	return
}

func inscription(w http.ResponseWriter,r *http.Request,login string,password string,mail string)(vrai bool){
	vrai = false
	mule := User{login,mail,password,password}
	if registreError(mule) && !findUser(mule.Login){
		if addUser(mule){
			register(w,r)
			session, err := store.Get(r, "session-name")
			session.Values["name"] = mule.Login
			session.Values["co"] = true
			session.Save(r,w)
			writeError(err)
			vrai = true
			return
		}
		
	}
	return
}
