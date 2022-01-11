package main
import (
    "fmt"
    "net/http"
)
type msg string
func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
   fmt.Fprint(resp, m)
}
func main() {
    msgHandler := msg("Hello world 2")
    fmt.Println("Server is listening...")
    http.ListenAndServe("0.0.0.0:8181", msgHandler)
}
