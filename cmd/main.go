package main

import (
	"fmt"
	"log"

	"github.com/135yshr/go-confernce-mini-2022/registry"
	"github.com/135yshr/go-confernce-mini-2022/server"
)

func main() {
	fmt.Println("===== start =====")
	s := server.New(registry.NewRegistry())
	log.Fatal(s.Start())
}
