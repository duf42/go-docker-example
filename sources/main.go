package main

import (
    "fmt"
    //"html"
    "log"
    "net/http"
    "os"
    "io/ioutil"
    "model"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    model.Initialize()

    http.Handle("/", http.FileServer(http.Dir("/web")))

    http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request){
        
        ver, err := ioutil.ReadFile("/config/VERSION")
        check(err)
        fmt.Fprintf(w,string(ver))

    })

    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))

}
