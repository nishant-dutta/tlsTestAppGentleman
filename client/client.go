package main

import (
	"log"

	"gopkg.in/h2non/gentleman.v2"
)

func main() {
	cli := gentleman.New()

	// Define base URL
	cli.URL("http://localhost:8080/hello")

	// Create a new request based on the current client
	req := cli.Request()

	// Define the URL path at request level
	// req.Path("/hello")

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
