package main

import (
    "crypto/rand"
    "fmt"
    "golang.org/x/crypto/nacl/sign"
    "io/ioutil"
)

func main() {
    // Generate a new key pair
    publicKey, privateKey, err := sign.GenerateKey(rand.Reader)
    if err != nil {
        panic(err)
    }

    // Read the package file
    data, err := ioutil.ReadFile("main.go")
    if err != nil {
        panic(err)
    }

    // Sign the package
    signedMessage := sign.Sign(nil, data, privateKey)
    fmt.Printf("Signed message: %x\n", signedMessage)

    // Save the signed message to a file
    err = ioutil.WriteFile("main_signed.go", signedMessage, 0644)
    if err != nil {
        panic(err)
    }

    fmt.Println("Package signed and saved to main_signed.go")

    // Verify the signed message
    message, ok := sign.Open(nil, signedMessage, publicKey)
    if !ok {
        fmt.Println("Invalid signature")
        return
    }
    fmt.Printf("Verified message: %s\n", message)
}
