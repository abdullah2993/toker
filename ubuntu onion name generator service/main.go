package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/asn1"
	"encoding/base32"
	"encoding/pem"
	"flag"
	"html/template"
	"log"
	"net/http"
	"strings"
)

//ServiceDesc contains services private key as well as hostname
type ServiceDesc struct {
	PrivateKey string
	Hostname   string
}

var addr = flag.String("http", "localhost:8080", "Address to listen on")
var gens = flag.Int("gen", 1, "Number of generators")
var queue = flag.Int("queue", 1, "Number of elemens in queue")

var results = make(chan *ServiceDesc, *queue)

func main() {
	defer close(results)
	for i := 0; i < *gens; i++ {
		go generator()
	}
	http.HandleFunc("/", handleIndex)
	log.Fatalf("server failed: %v", http.ListenAndServe(*addr, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./index.gohtml")
	if err != nil {
		log.Printf("unable to parse template: %v", err)
		http.Error(w, "unknow error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, <-results)
}

func generator() {
	const bits = 10
	const PEM = "RSA PRIVATE KEY"
	for {
		key, err := rsa.GenerateKey(rand.Reader, 1024)
		if err != nil {
			log.Printf("unable to generate private key: %v", err)
			continue
		}
		pub, err := asn1.Marshal(key.PublicKey)
		if err != nil {
			log.Printf("unable to encode public key: %v", err)
			continue
		}
		hashBytes := sha1.Sum(pub)
		hash := base32.StdEncoding.EncodeToString(hashBytes[:])
		exportedPriv := &pem.Block{
			Type:  PEM,
			Bytes: x509.MarshalPKCS1PrivateKey(key)}
		privateKey := pem.EncodeToMemory(exportedPriv)
		results <- &ServiceDesc{PrivateKey: string(privateKey), Hostname: strings.ToLower(hash[:16]) + ".onion"}

	}
}
