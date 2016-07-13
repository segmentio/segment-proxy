package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type SegmentServer int

const (
	CDN SegmentServer = iota
	TrackingAPI
)

func TestSegmentReverseProxy(t *testing.T) {
	cases := []struct {
		url            string
		expectedServer SegmentServer
	}{
		{"/v1/projects", CDN},
		{"/analytics.js/v1", CDN},
		{"/v1/import", TrackingAPI},
		{"/v1/pixel", TrackingAPI},
	}
	for _, c := range cases {
		cdn := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if c.expectedServer == CDN {
				fmt.Fprintln(w, "Hello, client")
			} else {
				t.Errorf("CDN unexpected request: %f\n", r.URL)
			}
		}))

		trackingAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if c.expectedServer == TrackingAPI {
				fmt.Fprintln(w, "Hello, client")
			} else {
				t.Errorf("Tracking API unexpected request: %f\n", r.URL)
			}
		}))

		proxy := httptest.NewServer(NewSegmentReverseProxy(mustParseUrl(cdn.URL), mustParseUrl(trackingAPI.URL)))

		_, err := http.Get(proxy.URL + c.url)
		if err != nil {
			t.Fatal(err)
		}

		cdn.Close()
		trackingAPI.Close()
	}
}

func mustParseUrl(raw string) *url.URL {
	u, err := url.Parse(raw)
	if err != nil {
		log.Fatal(err)
	}
	return u
}
