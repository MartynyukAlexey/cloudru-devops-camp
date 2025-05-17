package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"
	"strings"
)

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

// a techinque of passing dependencies via closures (function returns function)
// instead of building handlers as methods on top of a struct with dependencies.
// check the original post:
// https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#maker-funcs-return-the-handler
func handleIndex(author string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(content, "static/index.html")
		if err != nil {
			http.Error(w, "couldn't parse static files: "+err.Error(), http.StatusInternalServerError)
			return
		}

		hostname, err := os.Hostname()
		if err != nil {
			hostname = "couldn't get hostname"
		}

		ip, err := getClientIP(r)
		if err != nil {
			ip = "couldn't get client's ip"
		}

		err = tmpl.Execute(w, struct {
			Hostname string
			IP       string
			Author   string
		}{
			Hostname: hostname,
			IP:       ip,
			Author:   author,
		})
		if err != nil {
			http.Error(w, "couldn't execute template: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

// getting client's real ip isn't as trivial as to check r.RemoteAddr.
// the client may be behind proxy or NAT
func getClientIP(r *http.Request) (string, error) {
	ip := r.Header.Get("X-Real-IP")
	if netIP := net.ParseIP(ip); netIP != nil {
		return ip, nil
	}

	// when a request goes through proxy/NAT its ip is appended to the X-Forwarded-For's value.
	// the first ip in this header belongs to the original clien
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		for _, ip := range ips {
			ip = strings.TrimSpace(ip)
			if netIP := net.ParseIP(ip); netIP != nil {
				return ip, nil
			}
		}
	}

	// if no other option, fallback to r.RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	if netIP := net.ParseIP(ip); netIP != nil {
		return ip, nil
	}

	return "", fmt.Errorf("no ip address provided")
}
