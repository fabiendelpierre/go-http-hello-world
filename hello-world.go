package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func get_hostname() (hostname string) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hostname
}

func get_date() (today string) {
	currentTime := time.Now()
	formattedDate := currentTime.Format("02-Jan-2006")
	return formattedDate
}

func get_time() (now string) {
	currentTime := time.Now().UTC()
	formattedTime := currentTime.Format("15:04:05")
	return formattedTime
}

func get_client_ip_addr(req *http.Request) string {
	client_ip := req.Header.Get("X-FORWARDED-FOR")
	if client_ip != "" {
		return client_ip
	}
	return req.RemoteAddr
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	client_ip := get_client_ip_addr(r)
	greeting := "Hello! You're on host " + get_hostname() + ", visiting from " + client_ip + ".\nToday's date is " + get_date() + " and the UTC time is " + get_time() + "."
	fmt.Fprintf(w, greeting+"\n\n")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
