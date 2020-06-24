package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\n\nHowdy Folks..xD!!!\n\n")
	fmt.Fprintf(w, "My name is Ameya Makarand Mahajani\n")
	ipAddress, _, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Fprintf(w, "IP address is : %s\n\n", ipAddress)
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
	//srv := &http.Server{Addr: ":8080"}
	//go func() {
	//	srv.ListenAndServe()
	//}()
	//time.Sleep(10 * time.Second)
	//srv.Shutdown(context.TODO())
	//s.Shutdown(context.Background())
}
