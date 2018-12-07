package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/handlers"
)

const segmentCDNHost string = "http://cdn.segment.com"
const segmentTrackingAPIHost string = "http://api.segment.io"

// Replace with real mirror host URLs
var mirrorHosts = []string{
	"http://localhost:8081",
	"http://localhost:8082",
}

func parseURL(hostName string) *url.URL {
	targetURL, err := url.Parse(hostName)
	if err != nil {
		log.Fatalf("Failed to parse url %s: %v", hostName, err)
	}

	return targetURL
}

// singleJoiningSlash is copied from httputil.singleJoiningSlash method.
func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func isCDNUrl(url *url.URL) bool {
	urlStr := url.String()
	return strings.HasPrefix(urlStr, "/v1/projects") ||
		strings.HasPrefix(urlStr, "/analytics.js/v1")
}

func copyRequestTo(client *http.Client, url *url.URL, req *http.Request) error {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	destURI := fmt.Sprintf("%s%s", url.String(), req.RequestURI)

	proxyReq, err := http.NewRequest(req.Method, destURI, bytes.NewReader(body))
	if err != nil {
		return err
	}

	// We may want to filter some headers, otherwise we could just use a shallow copy
	// proxyReq.Header = req.Header
	proxyReq.Header = req.Header
	resp, err := client.Do(proxyReq)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

// NewSegmentReverseProxy is adapted from the httputil.NewSingleHostReverseProxy
// method, modified to dynamically redirect to different servers (CDN or Tracking API)
// based on the incoming request, and sets the host of the request to the host of of
// the destination URL.
func NewSegmentReverseProxy(
	client *http.Client,
	cdn *url.URL,
	trackingAPI *url.URL,
	mirrorHostURLs []*url.URL,
) http.Handler {
	director := func(req *http.Request) {
		// Figure out which server to redirect to based on the incoming request.
		var target *url.URL
		if isCDNUrl(req.URL) {
			target = cdn
		} else {
			target = trackingAPI
		}

		targetQuery := target.RawQuery
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}

		// Set the host of the request to the host of of the destination URL.
		// See http://blog.semanticart.com/blog/2013/11/11/a-proper-api-proxy-written-in-go/.
		req.Host = req.URL.Host
		for _, mirrorURL := range mirrorHostURLs {
			err := copyRequestTo(client, mirrorURL, req)
			if err != nil {
				log.Printf(
					"WARNING: Failed to mirror request to %s: %v\n",
					mirrorURL.String(), err,
				)
			}
		}
	}

	return &httputil.ReverseProxy{Director: director}
}

func createMirrorHostsList() []*url.URL {
	mirrorHostURLs := make([]*url.URL, len(mirrorHosts))
	for i, hostName := range mirrorHosts {
		mirrorHostURLs[i] = parseURL(hostName)
	}

	return mirrorHostURLs
}

var port = flag.String("port", "8080", "bind address")
var debug = flag.Bool("debug", false, "debug mode")

func main() {
	flag.Parse()

	proxy := NewSegmentReverseProxy(
		&http.Client{},
		parseURL(segmentCDNHost),
		parseURL(segmentTrackingAPIHost),
		createMirrorHostsList(),
	)
	if *debug {
		proxy = handlers.LoggingHandler(os.Stdout, proxy)
		log.Printf("serving proxy at port %v\n", *port)
	}

	log.Fatal(http.ListenAndServe(":"+*port, proxy))
}
