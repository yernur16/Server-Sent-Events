package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

type echo struct {
	// clients    map[chan string]string
	newClients chan chan string
	messages   chan string
}

func main() {
	app := fiber.New()
	app.Get("/echo", adaptor.HTTPHandler(handler(echoHandler)))
	app.Get("/say", adaptor.HTTPHandler(handler(sayHandler)))

	log.Println("Running on localhost 8080")
	app.Listen(":8080")
}

func (e *echo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		msg, open := <-e.messages

		if !open {
			break
		}

		fmt.Fprintf(w, "data: %s\n\n", msg)

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}

}

func handler(f http.HandlerFunc) http.Handler {
	return http.HandlerFunc(f)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println("error with parsing index html")
	}

}

func sayHandler(w http.ResponseWriter, r *http.Request) {
	go funkcia(client)

	timeout := time.After(1 * time.Second)

	select {
	case ev := <-client.messages:
	case <-timeout:
		fmt.Fprintf(w, ": nothing to sent\n\n")
	}

}
