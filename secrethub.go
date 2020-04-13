// Package main wraps the SecretHub Golang Client with cgo exported
// functions so it can be called from other languages, e.g. C, Python,
// Ruby, Node, and Java.
package main

import (
	"C"
	"encoding/json"

	"github.com/secrethub/secrethub-go/pkg/secrethub"
)

type ReadRequest struct {
	Path string `json:"path"`
}

type ReadResponse struct {
	Error  error  `json:"error"`
	Secret string `json:"secret"`
}

//export Read
func Read(cRequest *C.char) *C.char {
	req := &ReadRequest{}
	err := json.Unmarshal([]byte(C.GoString(cRequest)), req)
	if err != nil {
		panic(err)
	}

	client, err := secrethub.NewClient()
	if err != nil {
		panic(err)
	}

	secret, err := client.Secrets().ReadString(req.Path)
	res := &ReadResponse{
		Secret: secret,
		Error:  err,
	}

	resJson, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	return C.CString(string(resJson))
}

func main() {}
