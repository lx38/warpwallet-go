package main

import (
    "golang.org/x/crypto/pbkdf2"
    "golang.org/x/crypto/scrypt"
    "github.com/vsergeev/btckeygenie/btckey"
    "crypto/sha256"
    "bytes"
    "fmt"
    "os"
)

func main () {
    var passphrase string
    var salt string
    var priv btckey.PrivateKey
    var result bytes.Buffer
    
    if len(os.Args) < 3 {
        fmt.Printf("Usage: %s [Passphrase] [Salt]\n\n", os.Args[0])
        os.Exit(0)
    }
    
    passphrase = os.Args[1]
    salt = os.Args[2]
    
    fmt.Printf("Passphrase: %s\nSalt: %s\n", passphrase, salt)
    
    _passphrase := fmt.Sprint(passphrase, "\x01")
    _salt := fmt.Sprint(salt, "\x01")
    key, _ := scrypt.Key([]byte(_passphrase), []byte(_salt), 262144, 8, 1, 32)

    _passphrase = fmt.Sprint(passphrase, "\x02")
    _salt = fmt.Sprint(salt, "\x02")
    key2 := pbkdf2.Key([]byte(_passphrase), []byte(_salt), 65536, 32, sha256.New)

    for i := 0; i < len(key); i++ {
        result.WriteByte(key[i] ^ key2[i])
    }

    priv.FromBytes(result.Bytes())

    privkey := priv.ToWIF()
    address_uncompressed := priv.ToAddressUncompressed()
    fmt.Printf("Bitcoin Address: %s\n", address_uncompressed)
    fmt.Printf("Private Key: %s\n", privkey)
    os.Exit(0)

}

/*
--- USAGE ---
Windows: warpwallet [passphrase] [salt]
Linux: go run warpwallet.go [passphrase] [salt]
*/
