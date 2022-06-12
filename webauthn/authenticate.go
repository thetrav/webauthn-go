package webauthn

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func Challenge(w http.ResponseWriter, req *http.Request) {
	//TODO: generate and send signing request
}

func Authenticate(w http.ResponseWriter, req *http.Request) {
	//TODO: parse and validate signed request
}