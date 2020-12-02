package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

//HelloServer is a test server function
func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//JSONServer is a json test server function
func JSONServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    //name := req.PostFormValue("alias")
    jsonVariable := map[string]interface{}{"Success": 1}
    js, _ := json.Marshal(jsonVariable)
    w.Write(js)
}

func main() {
    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/json", JSONServer)
    http.ListenAndServe(":8080", nil)
}
