package reverz

import (
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

// Run starts of reverz
func (r *Reverz) Run(f func(w http.ResponseWriter, r *http.Request)) func(handler http.Handler) http.Handler  {
	u, _ := url.Parse(r.conf.URLs[0])
	reverseProxy := httputil.NewSingleHostReverseProxy(u)
	reverseProxy.ServeHTTP(proxy(f))
	return reverseProxy(method)
}

func proxy(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Host = r.URL.Host
		handler.ServeHTTP(w, r)
	})
}
