package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func main() {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("keygen error: ", err)
	}
	//pkcs1 := x509.MarshalPKCS1PrivateKey(key)
	pemdata := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)
	err = ioutil.WriteFile("key.pem", pemdata, 0600)
	if err != nil {
		log.Fatal("write error: ", err)
	}
}
