package webauthn

import (
    "fmt"
    "net/http"
    "encoding/json"
)

func GetRegistrationOptions(w http.ResponseWriter, req *http.Request) {
    challenge, err := MakeChallenge()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Printf("Failed to generate challenge: %w\n", err)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    userName := req.URL.Query()["userName"][0]

    opts := regOpts(challenge, userName)
    
    err = json.NewEncoder(w).Encode(opts)

	if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Printf("Failed to marshal json %w", err)
		return
	}
}

func Register(w http.ResponseWriter, req *http.Request) {
    var credential Credential;
    err := json.NewDecoder(r.Body).Decode(&credential)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    //TODO: validate credential
    //TODO: save credential against userId
}