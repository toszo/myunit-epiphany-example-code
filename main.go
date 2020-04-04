package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
)

func handlerCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "golang")
}

func handlerHost(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err == nil {
        fmt.Fprintf(w, hostname)
    } else {
        fmt.Println(err)
    }
}

func handlerNumber(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    var ret int
    if r.Form["number"] != nil {
        ret, _ = strconv.Atoi(r.Form["number"][0])
        ret*=2
    } else {
        ret = 0
    }
    fmt.Fprintf(w, strconv.Itoa(ret))
}

func main() {
    http.HandleFunc("/check", handlerCheck)
    http.HandleFunc("/status", handlerCheck)
    http.HandleFunc("/host", handlerHost)
    http.HandleFunc("/number", handlerNumber)
    log.Fatal(http.ListenAndServe(":8080", nil))
}