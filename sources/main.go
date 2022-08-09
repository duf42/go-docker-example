package main

import (
    "fmt"
    //"html"
    "log"
    "net/http"
    "os"
)

func main() {

    http.Handle("/", http.FileServer(http.Dir("/web")))

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))

}
