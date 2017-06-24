# warpwallet-go
An implementation of WarpWallet in Go.

*The goal is to eventually be portable with no 3rd party libraries.*

## Dependencies
golang.org/x/crypto/pbkdf2
golang.org/x/crypto/scrypt
github.com/vsergeev/btckeygenie/btckey

## Guide
1) Set your GOPATH directory
2) Download and install all required dependencies
3) Build with go build warpwallet.go
4) Run with either
Windows: `warpwallet [passphrase] [salt]` 
Linux: `go run warpwallet.go [passphrase] [salt]` 
