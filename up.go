package main


import(
	"net/http"
	"html/template"
	"fmt"
	"math/rand"
	"time"
	"crypto/md5"
	"os"
	"mime/multipart"
	"strconv"
	"io"
)

func uploads(w http.ResponseWriter,r *http.Request){
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))	
	t, _ := template.ParseFiles("upload.gtpl")
	t.Execute(w, token)
	
}

func uploadth(w http.ResponseWriter,r *http.Request){
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
		}
	defer file.Close()
	mimeType := handler.Header.Get("Content-Type")
	switch mimeType {
	case "application/json":
		saveFile(file,handler,".json")
	case "text/csv":
		saveFile(file,handler,".csv")
	default:
		println(78)
	}
}


func saveFile(file multipart.File, handler *multipart.FileHeader, a string ){
	name := tokenGenerator()
	if len(handler.Header["Content-Type"]) > 1 {
		println(69)
	}
	f, err := os.OpenFile("./test/"+name+a, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}


func tokenGenerator() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
