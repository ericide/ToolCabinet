package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("You need to give the path")
		return
	}

	serverWorkdir := os.Args[1]
	port := "10086"

	if len(os.Args) >= 3 {
		port = os.Args[2]
	}

	log.Println("Server path: " + serverWorkdir + " on Port:" + port)

	p, _ := filepath.Abs(filepath.Dir(serverWorkdir + "\\"))
	http.Handle("/", http.FileServer(http.Dir(p)))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
