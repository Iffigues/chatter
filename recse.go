package main

import(
	"encoding/json"
	"net/http"
)

func header(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "localhost")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS,HEAD")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, X-Requested-With, X-Prototype-Version, Origin, Allow, *")
	w.Header().Set("Access-Control-Max-Age", "1728000")
	w.Header().Set("Access-Control-Allow-Credentials"   , "true")
}

func grap(r *http.Request,ou interface{})(err error){
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&ou)
	writeError(err)
	defer r.Body.Close()
	return
}

func sendJson(ar interface{},w http.ResponseWriter)(err error){
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(ar);
	writeError(err)
	return
}

func sendRes(aa int,res string,w http.ResponseWriter){
	w.WriteHeader(aa)
	w.Write([]byte(res));
}
