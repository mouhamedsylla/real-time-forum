package gateway

import (
	"io"
	"log"
	"net/http"
	"real-time-forum/server/router"
)

var Gateway_EndPoint = map[string][]string{
	"9090": {
		"/api/message/private/:senderId/:receiverId",
		"/api/message/private/send/:receiverId",
		"/api/message/private/users/:userId",
	},
	"8080": {},
	"8181": {},
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

func (gtw *Gateway) BootstrapApp() {
	for port, endpoints := range Gateway_EndPoint {
		for _, endpoint := range endpoints {
			gtw.Router.Method(http.MethodGet).Handler(endpoint, gtw.Proxy(endpoint, "http://localhost:"+port))
		}
	}
	log.Fatal(http.ListenAndServe(":3000", gtw.Router))
}
