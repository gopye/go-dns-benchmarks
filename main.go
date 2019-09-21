package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// if no commandline args use a default
func main() {
	var lookupURL string
	if len(os.Args) > 1 {
		lookupURL = os.Args[1]
	} else {
		lookupURL = "www.example.com"
	}

	google := DNSResolver{name: "Google", server: "8.8.8.8"}
	google2 := DNSResolver{name: "Google2", server: "8.8.4.4"}
	cloudFlare := DNSResolver{name: "CloudFlare", server: "1.1.1.1"}
	quad9s := DNSResolver{name: "Quad 9s", server: "9.9.9.9"}
	openDNS := DNSResolver{name: "OpenDNS", server: "208.67.222.222"}
	openDNS2 := DNSResolver{name: "OpenDNS2", server: "208.67.222.220"}
	dnsAtvantage := DNSResolver{name: "DNS Advantage", server: "156.154.70.1"}

	resolvers := []DNSResolver{google, google2, cloudFlare, quad9s, openDNS, openDNS2, dnsAtvantage}

	for i, res := range resolvers {
		r := res.resolver()
		resolvers[i].avgTime = runLookups(&res, lookupURL, &r)
	}

	log.Printf("%s        - %s        avg %s", google.name, google.server, resolvers[0].avgTime)
	log.Printf("%s       - %s        avg %s", google2.name, google2.server, resolvers[1].avgTime)
	log.Printf("%s    - %s        avg %s", cloudFlare.name, cloudFlare.server, resolvers[2].avgTime)
	log.Printf("%s       - %s        avg %s", quad9s.name, quad9s.server, resolvers[3].avgTime)
	log.Printf("%s       - %s avg %s", openDNS.name, openDNS.server, resolvers[4].avgTime)
	log.Printf("%s      - %s avg %s", openDNS2.name, openDNS2.server, resolvers[5].avgTime)
	log.Printf("%s - %s   avg %s", dnsAtvantage.name, dnsAtvantage.server, resolvers[6].avgTime)
}

type DNSResolver struct {
	name, server, avgTime string
}

type lookupResults struct {
	ipAddr      []net.IPAddr
	elapsedTime time.Duration
}

func (d DNSResolver) dailer(ctx context.Context, network, address string) (net.Conn, error) {
	dial := net.Dialer{}
	return dial.DialContext(ctx, "udp", d.server+":53")
}

func (d DNSResolver) resolver() net.Resolver {
	return net.Resolver{
		PreferGo: true,
		Dial:     d.dailer,
	}
}

func lookupIPAddr(r *net.Resolver, cont context.Context, host string) (ipaddr []net.IPAddr, lookupTime time.Duration) {
	start := time.Now()
	ipaddr, err := r.LookupIPAddr(cont, host)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	return ipaddr, elapsed
}

func runLookups(d *DNSResolver, url string, r *net.Resolver) string {
	ctx := context.Background()
	lookUps := make(map[int]lookupResults)
	for i := 1; i <= 5; i++ {
		ipaddr, lookupTime := lookupIPAddr(r, ctx, url)
		log.Printf("%s - %s lookup %s took %s", d.name, d.server, ipaddr, lookupTime)
		lookUps[i] = lookupResults{ipaddr, lookupTime}
	}

	fmt.Println("")
	sum := 0
	for i := range lookUps {
		sum += int(lookUps[i].elapsedTime)
	}
	avg := time.Duration(sum / len(lookUps)).String()
	return avg
}
