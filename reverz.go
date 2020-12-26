package reverz

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Reverz defines
type Reverz struct {
	conf     *Config
	urls     []*url.URL
	balancer Balancer
}

// New provides initialization for Reverz
func New(conf *Config) (*Reverz, error) {
	if conf == nil {
		return nil, fmt.Errorf("config is not defined")
	}
	urls, err := convertURLs(conf.URLs)
	if err != nil {
		return nil, err
	}
	return &Reverz{
		conf:     conf,
		urls:     urls,
		balancer: selectBalancer(conf.Balancer, urls),
	}, nil
}

// Proxy defines endpoint for redirect
func (r *Reverz) Proxy(w http.ResponseWriter, req *http.Request) {
	u := r.balancer.Next()
	reverseProxy := httputil.NewSingleHostReverseProxy(u)
	req.URL.Host = u.Host
	req.URL.Scheme = u.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = u.Host
	reverseProxy.ServeHTTP(w, req)
}

// convertURLs provides converting of urls from slice of strings
// to slice of urls
func convertURLs(rawURLs []string) ([]*url.URL, error) {
	if len(rawURLs) == 0 {
		return nil, fmt.Errorf("urls is not defined")
	}
	urls := make([]*url.URL, len(rawURLs))
	for i, u := range rawURLs {
		urlResp, err := url.ParseRequestURI(u)
		if err != nil {
			return nil, fmt.Errorf("unable to parse url: %v", err)
		}
		urls[i] = urlResp
	}
	return urls, nil
}

// selectBalancer provides selecting of load balancer
func selectBalancer(b string, urls []*url.URL) Balancer {
	if b == "rr" {
		return &RoundRobin{urls: urls}
	}
	return &RoundRobin{urls: urls}
}
