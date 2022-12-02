package main

import (
  "encoding/base32"
  "fmt"
  "time"

  "github.com/pquerna/otp"
  "github.com/pquerna/otp/totp"
)

func main() {
  var secretString string
  fmt.Scan(&secretString)
  secret := base32.StdEncoding.EncodeToString([]byte(secretString))
  passcode, _ := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
    Period:    30,
    Skew:      1,
    Digits:    10,
    Algorithm: otp.AlgorithmSHA512,
  })
  // fmt.Println(err)
  fmt.Println(passcode)
}