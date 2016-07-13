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

func main() {
	cdnURL, err := url.Parse("http://cdn.segment.com")
	if err != nil {
		log.Fatal(err)
	}
	cdnProxy := httputil.NewSingleHostReverseProxy(cdnURL)
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
