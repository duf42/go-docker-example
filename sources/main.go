package main

import (
    "fmt"
    //"html"
    "log"
    "net/http"
    "os"
    "io/ioutil"
    "model"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    model.Initialize()

    http.Handle("/", http.FileServer(http.Dir("/web")))

    http.HandleFunc("/step", func(w http.ResponseWriter, r *http.Request){
        
        target, _  := strconv.ParseFloat(r.URL.Query().Get("target"), 64)
        current, _ := strconv.ParseFloat(r.URL.Query().Get("current"), 64)

        model.SetInput("target",   target)
        model.SetInput("current",  current)
        model.Step()
        fmt.Fprintf(w,"Command = %f", model.GetOutput("command"))

    })

    http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request){
        
        ver, err := ioutil.ReadFile("/config/VERSION")
        check(err)
        fmt.Fprintf(w,string(ver))

    })

    log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))

}
