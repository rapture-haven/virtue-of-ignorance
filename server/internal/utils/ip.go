package utils

import "net/http"

func GetHttpRequestIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}

	// TODO: IPAddress = r.RemoteAddr можно просто так?

	return IPAddress
}
