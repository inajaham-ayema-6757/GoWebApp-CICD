package main

import (
	"fmt"
	"net"
	"net/http"
)

func greeting() string {
	return "Howdy Folks..xD!!!" 		//Greeting Message
}

func introduction() string {
	return "My name is Ameya Makarand Mahajani"		//Intro Message
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\n\n\n")
	fmt.Fprintf(w, greeting())
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, introduction())
	fmt.Fprintf(w, "\n")
	IPAddr, _, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Fprintf(w, "IP address is : %s", IPAddr)		//Printing Current IP Address
}

func main() {
	fmt.Println("Server is running")
	fmt.Println("Ctrl+C = close the server. ")
	http.HandleFunc("/", helloWorld)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
