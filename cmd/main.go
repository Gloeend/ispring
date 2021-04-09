package main

import (
	"fmt"
	"ispring/pkg/transport"
	"net/http"
	"os"
	_ "strings"
)

func main()  {
	port := os.Getenv("PORT")

	http.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request)) {
		fmt.Fprintf(w, "World")
	}
	//r := transport.Router()
	//fmt.Println(http.ListenAndServe(":8000", r))
	http.ListenAndServe(":"+port, nil)

}