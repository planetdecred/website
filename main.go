package main

// To run: go run main.go 

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	hostnames := []string{
		"planetdecred.com",
	}
	tlsEnabled := 1
	filesDir := "."
	addr := ":443"
	fs := http.FileServer(http.Dir(filesDir))
	http.Handle("/", fs)
	if tlsEnabled != 1 {
		addr = ":8080"
	}
	s := &http.Server{
		Addr: addr,
	}
	if addr == ":443" {
		fmt.Printf("Serving TLS as %q..\n", hostnames)
		m := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			Cache:      autocert.DirCache("/etc/secrets/acme/"),
			HostPolicy: autocert.HostWhitelist(hostnames...),
		}
		s.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
		log.Fatal(s.ListenAndServeTLS("", ""))
	} else {
		fmt.Printf("Serving HTTP on %s..\n", addr)
		log.Fatal(s.ListenAndServe())
	}
}
