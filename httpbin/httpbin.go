package main

import (
	"encoding/json"
	"net"
	"net/http"
)

func getIP(rw http.ResponseWriter, request *http.Request) {
	ip := request.Header.Get("X-FORWARDED-FOR")
	if ip == "" {
		ip, _, _ = net.SplitHostPort(request.RemoteAddr)
	}
	data := map[string]string{
		"origin": ip,
	}
	serializedData, _ := json.Marshal(data)
	rw.Write(serializedData)
}

func main() {
	http.HandleFunc("/", getIP)
	http.ListenAndServe(":5000", nil)
}
