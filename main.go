package main
import(
	"net/http"
	"time"
	"github.com/urfave/negroni"
	"strconv"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/csrf"
	"github.com/GeertJohan/go.rice"
)

var(
	Router = NewRouter()
	ReadTimeOut,_ = strconv.Atoi(getKey(htp,"ReadTime"))
	WriteTimeOut ,_= strconv.Atoi(getKey(htp,"WriteTime"))
	Socket = getKey(htp,"Socket")
	CSRF = csrf.Protect([]byte("32-byte-long-auth-key"))
	n = negroni.Classic()
	headersOk = handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk = handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk = handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
)

func init(){
	Router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	n.UseHandler(Router)

}

func console(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header(w)
		next.ServeHTTP(w, r)
	})
}

func main() {

	s := &http.Server{
		Addr:           Socket,
		Handler:   console(n),
		ReadTimeout:    time.Duration(ReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(WriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	writeError(s.ListenAndServe())
}
