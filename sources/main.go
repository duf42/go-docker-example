package main

import (
    "fmt"
    //"html"
    //"log"
    //"net/http"
    "os"
)

func main() {

   fmt.Println("Current port is " + os.Getenv("PORT"));
/*
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
*/
}
