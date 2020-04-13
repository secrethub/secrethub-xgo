// Package main wraps the SecretHub Golang Client with cgo exported
// functions so it can be called from other languages, e.g. C, Python,
// Ruby, Node, and Java.
package main

import (
	"C"

	"encoding/base64"

	"github.com/secrethub/secrethub-go/pkg/secrethub"
	"github.com/secrethub/secrethub-xgo/bridge"
	"google.golang.org/protobuf/proto"
)

//export Read
func Read(cRequest *C.char) *C.char {
	req := &bridge.ReadRequest{}

	enc, err := base64.StdEncoding.DecodeString(C.GoString(cRequest))
	if err != nil {
		panic(err)
	}

	err = proto.Unmarshal(enc, req)
	if err != nil {
		panic(err)
	}

	client, err := secrethub.NewClient()
	if err != nil {
		panic(err)
	}

	res := &bridge.ReadResponse{}
	secret, err := client.Secrets().ReadString(req.Path)
	if err == nil {
		res.Secret = &bridge.Secret{
			Path: req.Path,
			Data: []byte(secret),
		}
	} else {
		res.Error = err.Error()
	}

	resBytes, err := proto.Marshal(res)
	if err != nil {
		panic(err)
	}

	return C.CString(base64.StdEncoding.EncodeToString(resBytes))
}

func main() {}
