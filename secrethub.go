// Package main wraps the SecretHub Golang Client with cgo exported
// functions so it can be called from other languages, e.g. C, Python,
// Ruby, Node, and Java.
package main

import "C"

import "github.com/secrethub/secrethub-go/pkg/secrethub"

//export Read
func Read(cCredential *C.char, cPassphrase *C.char, cPath *C.char) *C.char {
	path := C.GoString(cPath)
	credential := C.GoString(cCredential)
	passphrase := C.GoString(cPassphrase)

	cred, err := secrethub.NewCredential(credential, passphrase)
	if err != nil {
		panic(err)
	}

	result, err := secrethub.NewClient(cred, nil).Secrets().Versions().GetWithData(path)
	if err != nil {
		panic(err)
	}

	return C.CString(string(result.Data))
}

//export Exists
func Exists(cCredential *C.char, cPassphrase *C.char, cPath *C.char) bool {
	path := C.GoString(cPath)
	credential := C.GoString(cCredential)
	passphrase := C.GoString(cPassphrase)

	cred, err := secrethub.NewCredential(credential, passphrase)
	if err != nil {
		panic(err)
	}

	exists, err := secrethub.NewClient(cred, nil).Secrets().Exists(path)
	if err != nil {
		panic(err)
	}

	return exists
}

func main() {}
