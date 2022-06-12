package main

import (
    "fmt"
    "net/http"

    "github.com/thetrav/webauthn-go/webauthn"
)

func main() {
    http.HandleFunc("/registration", webauthn.GetRegistrationOptions)
    http.HandleFunc("/register", webauthn.Register)
    http.HandleFunc("/challenge", webauthn.Challenge)
    http.HandleFunc("/authenticate", webauthn.Authenticate)

    fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

    fmt.Printf("visit http://localhost:8090/index.html\n")
    http.ListenAndServe(":8090", nil)
}