package reverz

import (
	"net/url"
	"sync/atomic"
)

// Balancer provides implementation of balancing
type Balancer interface {
	Next() *url.URL
}

// RoundRobin provides implementation of round robin algorithms 
type RoundRobin struct {
	urls []*url.URL
	next uint32
}

// Next select teh next host
func (r *RoundRobin) Next() *url.URL {
	n := atomic.AddUint32(&r.next, 1)
	return r.urls[(int(n)-1)%len(r.urls)]
}