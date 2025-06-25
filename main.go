package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/AfterShip/email-verifier"
    "github.com/gorilla/mux"
)

type VerificationResult struct {
    Success bool        `json:"success"`
    Result  interface{} `json:"result,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func main() {
    verifier := emailverifier.NewVerifier().
        EnableSMTPCheck().
        EnableDomainSuggest().
        EnableAutoUpdateDisposable()

    r := mux.NewRouter()
    r.HandleFunc("/v1/{email}/verification", func(w http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        email := vars["email"]

        result, err := verifier.Verify(email)
        w.Header().Set("Content-Type", "application/json")
        if err != nil {
            json.NewEncoder(w).Encode(VerificationResult{Success: false, Error: err.Error()})
            return
        }
        json.NewEncoder(w).Encode(VerificationResult{Success: true, Result: result})
    }).Methods("GET")

    fmt.Println("API server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
} 