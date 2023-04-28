package u_load_balancer

import (
	"fmt"
	"net/url"
	
	roundrobin "github.com/hlts2/round-robin"
)

func NewRoundRobinBalancer(hosts []string) (roundrobin.RoundRobin, error) {
	if len(hosts) == 0 {
		return nil, fmt.Errorf("missing hosts")
	}
	var urls = make([]*url.URL, 0, len(hosts))
	for _, host := range hosts {
		u, err := url.Parse(host)
		if err != nil {
			return nil, err
		}
		urls = append(urls, u)
	}
	return roundrobin.New(urls...)
}

func NewPanicRoundRobinBalancer(hosts []string) roundrobin.RoundRobin {
	rr, err := NewRoundRobinBalancer(hosts)
	if err != nil {
		panic(err)
	}
	return rr
}
