package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Who Go is nit")

	//http://localhost:999 will show this line in a web browser
}
func main() {
	http.HandleFunc("/", index_handler)
	http.ListenAndServe(":999", nil)
}
