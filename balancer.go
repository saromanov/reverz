package reverz

// Balancer provides implementation of balancing
type Balancer interface {
	Next() string
}