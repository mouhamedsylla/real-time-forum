package middleware

import (
	"log"
	"net/http"
	"os"
	"github.com/mouhamedsylla/term-color/color"
)

var infosLog *log.Logger
var clr = color.Color().SetText("[INFO] ")

func init() {
	infosLog = log.New(os.Stdout, clr.Colorize(clr.Green), log.Ldate|log.Ltime|log.Lshortfile)
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		infosLog.Printf("%s---%s-%s-%s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}
