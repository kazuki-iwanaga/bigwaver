package main

import (
 "crypto/hmac"
 "crypto/sha256"
 "crypto/subtle"
 "encoding/hex"
 "fmt"
 "io"
 "log"
 "net/http"
 "os"
 "time"
)

func signature(secret string, message string) string {
 mac := hmac.New(sha256.New, []byte(secret))
 mac.Write([]byte(message))
 return hex.EncodeToString(mac.Sum(nil))
}

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 x_hub_signature_256 := r.Header.Get("X-Hub-Signature-256")
 if x_hub_signature_256 == "" {
  w.WriteHeader(http.StatusBadRequest)
  fmt.Fprintf(w, "Bad Request")
  return
 }

 buf, err := io.ReadAll(r.Body)
 if err != nil {
  w.WriteHeader(http.StatusInternalServerError)
  fmt.Fprintf(w, "Internal Server Error")
  return
 }
 body := string(buf)

 // https://docs.github.com/en/webhooks/using-webhooks/validating-webhook-deliveries#validating-webhook-deliveries
 expected_signature := "sha256=" + signature(os.Getenv("SECRET"), body)
 if subtle.ConstantTimeCompare(
  []byte(expected_signature), []byte(x_hub_signature_256)) != 1 {
  w.WriteHeader(http.StatusBadRequest)
  fmt.Fprintf(w, "Bad Request")
  return
 }
}

func main() {
 // https://pkg.go.dev/net/http#hdr-Servers
 s := &http.Server{
  Addr:         ":8080",
  Handler:      &Handler{},
  ReadTimeout:  10 * time.Second,
  WriteTimeout: 10 * time.Second,
 }
 log.Fatal(s.ListenAndServe())
}