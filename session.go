package main

import(
	"net/http"
	"github.com/gorilla/sessions"
)

var(
	store = sessions.NewCookieStore([]byte("Maybe)I(<script>Have-*Say$ThatWas&Wrong/*</script>p"))
)


func register(w http.ResponseWriter,r *http.Request){
	session, err := store.Get(r, "session-name")
	if err == nil {
		session.Values["co"] = true
		session.Save(r, w)
	}
}

func haveName(r *http.Request)(vrai bool){
	vrai = false
	session, err := store.Get(r, "session-name")
	if err == nil{
		if _,ok :=session.Values["name"];ok{
			vrai = true
		}
	}
	return
}


func unregister(w http.ResponseWriter,r *http.Request){
	session, err := store.Get(r, "session-name")
	if err == nil {
		println("eee")
		session.Values["co"] = false
		session.Save(r, w)
	}
}

func isRegister(r *http.Request)(vrai bool){
	vrai = false
	session, err := store.Get(r, "session-name")
	if err == nil{
		if aa,ok :=session.Values["co"];ok{
			if aa == true {
				vrai = true
			}
		}
	}
	return 
}
