# go-dns-benchmarks
Runs a DNS lookup 5 times, and averages out the time.

`$ go run main.go testdomain.com`

```
2019/09/21 14:11:08 Google        - 8.8.8.8        avg 22.801956ms
2019/09/21 14:11:08 Google2       - 8.8.4.4        avg 20.767106ms
2019/09/21 14:11:08 CloudFlare    - 1.1.1.1        avg 19.154577ms
2019/09/21 14:11:08 Quad 9s       - 9.9.9.9        avg 53.829064ms
2019/09/21 14:11:08 OpenDNS       - 208.67.222.222 avg 26.413981ms
2019/09/21 14:11:08 OpenDNS2      - 208.67.222.220 avg 20.318214ms
2019/09/21 14:11:08 DNS Advantage - 156.154.70.1   avg 65.551849ms
```