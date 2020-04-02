package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
)

func handlerCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Ok")
}

func handlerHost(w http.ResponseWriter, r *http.Request) {
    hostname, err := os.Hostname()
    if err == nil {
        fmt.Fprintf(w, hostname)
    } else {
        fmt.Println(err)
    }
}

func handlerLastValue(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    ret, _ := strconv.Atoi(r.Form["number"][0])
    ret*=2
    fmt.Fprintf(w, strconv.Itoa(ret))
}

func main() {
    http.HandleFunc("/check", handlerCheck)
    http.HandleFunc("/status", handlerCheck)
    http.HandleFunc("/host", handlerHost)
    http.HandleFunc("/last", handlerLastValue)
    log.Fatal(http.ListenAndServe(":8080", nil))
}