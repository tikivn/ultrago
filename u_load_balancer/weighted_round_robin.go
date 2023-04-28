package u_load_balancer

import (
	"fmt"

	lb2 "github.com/hedzr/lb"
	"github.com/hedzr/lb/lbapi"
)

func NewWeightedRoundRobinBalancer(nodes []*Node) (lbapi.Balancer, error) {
	if len(nodes) == 0 {
		return nil, fmt.Errorf("missing hosts")
	}
	lb := lb2.New(lb2.WeightedRoundRobin)
	for _, node := range nodes {
		lb.Add(node)
	}
	return lb, nil
}

func NewPanicWeightedRoundRobinBalancer(nodes []*Node) lbapi.Balancer {
	lb, err := NewWeightedRoundRobinBalancer(nodes)
	if err != nil {
		panic(err)
	}
	return lb
}

type Node struct {
	Addr     string
	Weighted int
}

func (s *Node) String() string { return s.Addr }
func (s *Node) Weight() int    { return s.Weighted }
