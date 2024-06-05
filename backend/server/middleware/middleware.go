package middleware

import (
	"log"
	"net/http"
	"os"
	"real-time-forum/utils"
	"real-time-forum/utils/jwt"

	"github.com/mouhamedsylla/term-color/color"
)

var (
	infosLog *log.Logger
	clr = color.Color().SetText("[INFO] ")
 	Jwt = jwt.JWT{}
)

func init() {
	infosLog = log.New(os.Stdout, clr.Colorize(clr.Green), log.Ldate|log.Ltime|log.Lshortfile)
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			// Gérer les requêtes prévol
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("forum")
		if err != nil {
			utils.ResponseWithJSON(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		publicKey := utils.GetPublicKey()
		_, err = Jwt.VerifyToken(session.Value, publicKey)

		if err != nil {
			utils.ResponseWithJSON(w, "Unauthorized", http.StatusUnauthorized)
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
