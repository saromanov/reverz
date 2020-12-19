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
	return &Reverz {
		conf: conf,
	}, nil
}

// Run starts of reverz
func (r *Reverz) Run(t func(handler http.Handler) http.Handler) error {
	u, _ := url.Parse(r.conf.URLs[0])
    reverseProxy := httputil.NewSingleHostReverseProxy(u)
    http.ListenAndServe(":5000", t(reverseProxy))
	return nil
}