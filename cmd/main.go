package main

import (
	"fmt"
	"log"

	"github.com/135yshr/go-confernce-mini-2022/handlers"
	"github.com/135yshr/go-confernce-mini-2022/server"
)

func main() {
	fmt.Println("Hello, World!")
	s := server.New()
	s.AddHandle("/hello", handlers.NewHelloHandler())
	log.Fatal(s.Start())
}
