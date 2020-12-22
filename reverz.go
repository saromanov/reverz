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

// convertURLs provides converting of urls from slice of strings
// to slice of urls
func (r *Reverz) convertURLs(rawURLs []string) ([]*url.URL, error) {
	urls := make([]*url.URL, len(rawURLs))
	for _, u := range rawURLs {
		urlResp, err := url.ParseRequestURI(u)
		if err != nil {
			return nil, fmt.Errorf("unable to parse url: %v", err)
		}
		urls = append(urls, urlResp)
	}
	return urls, nil
}