package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func getIp(r *http.Request) (string, error) {
	if client_ip := r.Header.Get("X-Real-Ip"); client_ip != "" {
		return client_ip, nil
	}
	if client_ip := r.Header.Get("X-Forwarded-For"); client_ip != "" {
		return client_ip, nil
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	return host, nil
}

func giveIp(w http.ResponseWriter, r *http.Request) {
	client_ip, err := getIp(r)

	if err != nil {
		client_ip = err.Error()
	}
	fmt.Fprintln(w, client_ip)
}

func main() {
	http.HandleFunc("/", giveIp)
	if err := http.ListenAndServe(":4723", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
