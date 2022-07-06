package swordtech

import (
	"net/http"
	"regexp"
	"strings"
)

// Request is an enhanced http.Request
type Request struct {
	*http.Request
	IPAddress string
}

// NewRequest creates a new HiveRequest
func NewRequest(r *http.Request) *Request {
	ip := r.RemoteAddr
	re := regexp.MustCompile(`(?m)\[([^\]]*)\]`)

	if len(re.FindStringIndex(ip)) > 0 {
		ip = re.FindStringSubmatch(ip)[1]
	}

	if idx := strings.Index(ip, ":"); idx > -1 {
		ip = ip[0:idx]
	}
	return &Request{r, ip}
}
