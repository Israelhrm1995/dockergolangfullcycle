package main;
import (
    "fmt"
    "log"
    "net/http"
)
func main() {
    http.HandleFunc("/FullCycleRocks", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Full Cycle Rocks!!")
    })
    fmt.Printf("Server running (port=8080), route: http://localhost:8080/FullCycleRocks\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

