package proxy

import "fmt"

type Proxy struct {
	Application       IServer
	MaxAllowedRequest int
	RateLimiter       map[string]int
}

//
func NewProxy() *Proxy {
	return &Proxy{
		Application:       &Application{},
		MaxAllowedRequest: 2,
		RateLimiter:       make(map[string]int),
	}
}

func (p *Proxy) IsRateLimited(key string) bool {
	if p.RateLimiter[key] >= p.MaxAllowedRequest {
		return false
	}
	p.RateLimiter[key]++
	return true
}

func (p *Proxy) HandleRequest(url, method string) (int, string) {
	if !p.IsRateLimited(url + ":" + method) {
		return 403, "Not Allowed"
	}
	return p.Application.HandleRequest(url, method)
}

func TestProxy() {
	proxy := NewProxy()
	url := "user"
	var body string
	var HttpCode int
	HttpCode, body = proxy.HandleRequest(url, "GET")
	fmt.Printf("\nURL: %s HttpCode: %d Body: %s", url, HttpCode, body)
	cnt := 0
	for {
		cnt++
		if cnt >= 5 {
			break
		}
		HttpCode, body = proxy.HandleRequest(url, "POST")
		fmt.Printf("\nURL: %s HttpCode: %d Body: %s", url, HttpCode, body)
	}
}
