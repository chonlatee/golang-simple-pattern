package main

type nginx struct {
	application *application
	maxAllowRequest int
	rateLimiter map[string]int
}

func newNginxServer() *nginx {
	return &nginx{
		application:     &application{},
		maxAllowRequest: 2,
		rateLimiter:     make(map[string]int),
	}
}

func (n *nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}

func (n *nginx) checkRateLimiting(url string) bool {
	n.rateLimiter[url]++
	if n.rateLimiter[url] > n.maxAllowRequest {
		return false
	}
	return true
}