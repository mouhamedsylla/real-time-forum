package gateway

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"real-time-forum/server/router"
)

var Gateway_EndPoint = map[string][]string{
	"9090": {
		"/chat/message/private/:senderId/:receiverId",
		"/chat/message/private/send/:receiverId",
		"/chat/message/private/users/:userId",
	},
	"8080": {
		"/auth/getGroupUser/:userId",
		"/auth/checkToken",
		"/auth/register",
		"/auth/login",
	},
	"8181": {
		"/posts/getAllPost",
		"/posts/createdpost/:userId",
		"/posts/:postId/comment",
		"/posts/:postId/getcomment",
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
	tmp, err := template.ParseFiles("../../../frontend/assets/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmp.Execute(w, nil)
}

func (gtw *Gateway) BootstrapApp() {
	gtw.Router.SetDirectory("/frontend/", "../../../frontend/")
	gtw.Router.Method(http.MethodGet).Handler("/frontend/", gtw.Router.StaticServe())
	gtw.Router.Method(http.MethodGet).Handler("/", http.HandlerFunc(Home))
	for port, endpoints := range Gateway_EndPoint {
		for _, endpoint := range endpoints {
			gtw.Router.Method(http.MethodPost, http.MethodGet).
				Handler(endpoint, gtw.Proxy(endpoint, "http://localhost:"+port))
		}
	}
	log.Fatal(http.ListenAndServe(":3000", gtw.Router))
}
