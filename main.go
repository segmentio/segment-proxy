package main

import (
    "encoding/json"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)
 
func main() { 
    http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        res := struct{ Status string }{"ok"}
        jData, _ := json.Marshal(&res)

        w.WriteHeader(http.StatusOK)
        w.Write(jData)
    })
 
    u, _ := url.Parse("https://api.segment.io/")
    http.Handle("/", httputil.NewSingleHostReverseProxy(u))
 
    // Start the server
    http.ListenAndServe(":80", nil)
    log.Printf("Proxy Initialized")
}