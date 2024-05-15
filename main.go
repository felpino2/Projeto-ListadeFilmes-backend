package main

import "net/http"

func main() {

	Run()
}


func Run() {
	
	s := http.Server{
		Addr: "localhost:8080",
	} 
	
	s.ListenAndServe()
}


