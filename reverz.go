package reverz

import (
	"log"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Reverz defines
type Reverz struct {
	conf *Config
}

// New provides initialization for Reverz
func New(conf *Config) (*Reverz, error) {
	if conf == nil {
		return nil, fmt.Errorf("config is not defined")
	}
	return &Reverz{
		conf: conf,
	}, nil
}

// Proxy defines endpoint for redirect
func (r *Reverz) Proxy(w http.ResponseWriter, req *http.Request)  {
	u, err := url.Parse(r.conf.URLs[0])
	if err != nil {
		log.Fatalf("unable to parse url: %v", err)
	}
	reverseProxy := httputil.NewSingleHostReverseProxy(u)
	req.URL.Host = u.Host
	req.URL.Scheme = u.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = u.Host
	reverseProxy.ServeHTTP(w, req)
}