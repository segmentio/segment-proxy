package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

// patchHost sets the host of the request to the host of of the destination URL.
// See http://blog.semanticart.com/blog/2013/11/11/a-proper-api-proxy-written-in-go/.
func patchHost(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Host = r.URL.Host
		handler.ServeHTTP(w, r)
	})
}

func main() {
	cdnURL, err := url.Parse("http://cdn.segment.com")
	if err != nil {
		log.Fatal(err)
	}
	cdnProxy := patchHost(httputil.NewSingleHostReverseProxy(cdnURL))
	cdnProxyServer := httptest.NewServer(cdnProxy)
	defer cdnProxyServer.Close()

	// Test code.
	url := cdnProxyServer.URL + "/v1/projects/DQf6nqU1PaMbcVQenYzYSRd6nkUL21b8/settings"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", b)
}
