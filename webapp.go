package main

import (
	"fmt"
	"net/http"
)

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Who Go is nit")

	//http://localhost:999 will show this line in a web browser
}
func abouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is how it change")

	//http://localhost:999/abouut/ will show this line in a web browser
}

func main() {
	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/abouut/", abouthandler)
	http.ListenAndServe(":999", nil)
}
