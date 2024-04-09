package main

import (
	"log"
	"os"

	"crypto/tls"
	"crypto/x509"

	"gopkg.in/h2non/gentleman.v2"
	gtls "gopkg.in/h2non/gentleman.v2/plugins/tls"
)

func main() {
	caCert, err := os.ReadFile("./certificates/cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.LoadX509KeyPair("./certificates/cert.pem", "./certificates/key.pem")
	if err != nil {
		log.Fatal(err)
	}

	cli := gentleman.New()

	config := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}

	cli.Use(gtls.Config(config))

	// Define base URL
	cli.URL("https://localhost:8443/hello")

	// Create a new request based on the current client
	req := cli.Request()

	// Set a new header field
	req.SetHeader("Client", "gentleman")

	// Perform the request
	res, err := req.Send()
	if err != nil {
		log.Printf("Request error: %s\n", err)
		return
	}
	if !res.Ok {
		log.Printf("Invalid server response: %d\n", res.StatusCode)
		return
	}

	// Reads the whole body and returns it as string
	log.Printf("Body: %s", res.String())
}
