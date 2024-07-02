package gateway

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"real-time-forum/server/middleware"
	"real-time-forum/server/router"
	"strings"
)

var Gateway_EndPoint = map[string][]string{
	"9090": {
		"/chat/message/private/:senderId/:receiverId",
		"/chat/message/private/send/:receiverId",
		"/chat/message/private/users/:userId",
	},
	"8080": {
		"/auth/getUsers",
		"/auth/public/register",
		"/auth/public/login",
	},
	"8181": {
		"/posts/getAllPost",
		"/posts/createdpost/:userId",
		"/posts/:postId/comment",
		"/posts/:postId/getcomment",
	},
	"9191": {
		"/chat/message/private/getConnectedUser/:userId",
	},
}

type Gateway struct {
	Router *router.Router
}

func NewGateway() *Gateway {
	return &Gateway{
		Router: router.NewRouter(),
	}
}

func (gtw *Gateway) Proxy(path, target string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		targetURL := target + r.URL.Path

		if r.URL.RawQuery != "" {
			targetURL += "?" + r.URL.RawQuery
		}
		req, err := http.NewRequest(r.Method, targetURL, r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		req.Header = r.Header

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		w.WriteHeader(resp.StatusCode)

		// Copy the response body to the client
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
	})
}

func (gtw *Gateway) Authenticate() {

}

func Home(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("../../../../frontend/assets/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmp.Execute(w, nil)
}

func (gtw *Gateway) SubcribeHandler(port string, endpoint string) {
	target := "http://localhost:" + port
	if strings.Contains(endpoint, "public") {
		gtw.Router.Method(http.MethodPost, http.MethodGet).
			Handler(endpoint, gtw.Proxy(endpoint, target))
	} else {
		gtw.Router.Method(http.MethodPost, http.MethodGet).
			Middleware(middleware.Authenticate).
			Handler(endpoint, gtw.Proxy(endpoint, target))
	}
}

func (gtw *Gateway) BootstrapApp() {
	gtw.Router.SetDirectory("/frontend/", "../../../../frontend/")
	gtw.Router.Method(http.MethodGet).Handler("/frontend/", gtw.Router.StaticServe())
	gtw.Router.Method(http.MethodGet).Handler("/", http.HandlerFunc(Home))
	for port, endpoints := range Gateway_EndPoint {
		for _, endpoint := range endpoints {
			gtw.SubcribeHandler(port, endpoint)
		}
	}
	fmt.Println("Server is running on: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", gtw.Router))
}
