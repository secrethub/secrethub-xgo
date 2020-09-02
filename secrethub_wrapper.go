package main

// #include <stdbool.h>
import "C"

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/secrethub/secrethub-go/pkg/secrethub"
)

// Read retrieves a secret from SecretHub, given its path.
// It throws an error if it fails to initialize the SecretHub client
// or if it fails to retrieve the secret.
//export Read
func Read(path *C.char, errMessage **C.char) *C.char {
	client, err := secrethub.NewClient()
	if err != nil {
		*errMessage = C.CString(err.Error())
		return nil
	}
	secret, err := client.Secrets().Read(C.GoString(path))
	if err != nil {
		*errMessage = C.CString(err.Error())
		return nil
	}
	return C.CString(string(secret.Data))
}

// Resolve fetches the values of a secret from SecretHub, when the `ref` parameter
// has the format `secrethub://<path>`. Otherwise it returns `ref` unchanged, as an array of bytes.
//export Resolve
func Resolve(ref *C.char, errMessage **C.char) *C.char {
	client, err := secrethub.NewClient()
	if err != nil {
		*errMessage = C.CString(err.Error())
		return nil
	}
	bits := strings.Split(C.GoString(ref), "://")
	if len(bits) == 2 && bits[0] == "secrethub" {
		secret, err := client.Secrets().Read(bits[1])
		if err != nil {
			*errMessage = C.CString(err.Error())
			return nil
		}
		return C.CString(string(secret.Data))
	}
	return ref
}

// ResolveEnv takes a map of environment variables and replaces the values of those
// which store references of secrets in SecretHub (`secrethub://<path>`) with the value
// of the respective secret. The other entries in the map remain untouched.
//export ResolveEnv
func ResolveEnv(errMessage **C.char) *C.char {
	envVars := os.Environ()
	resolvedEnv := make(map[string]string, len(envVars))
	for _, value := range envVars {
		keyValue := strings.Split(value, "=")
		resolvedEnv[keyValue[0]] = C.GoString(Resolve(keyValue[1], errMessage))
	}
	encoding, err := json.Marshal(resolvedEnv)
	if err != nil {
		*errMessage = C.CString(err.Error())
		return nil
	}
	return C.CString(encoding)
}

// Exists checks if a secret exists at `path`.
//export Exists
func Exists(path *C.char, errMessage **C.char) C.bool {
	client, err := secrethub.NewClient()
	if err != nil {
		*errMessage = C.CString(err.Error())
		return C.bool(false)
	}
	exists, err := client.Secrets().Exists(C.GoString(path))
	if err != nil {
		*errMessage = C.CString(err.Error())
		return C.bool(false)
	}
	return C.bool(exists)
}

// Remove deletes the secret found at `path`, if it exists.
//export Remove
func Remove(path *C.char, errMessage **C.char) {
	client, err := secrethub.NewClient()
	if err != nil {
		*errMessage = C.CString(err.Error())
		return
	}
	err = client.Secrets().Delete(C.GoString(path))
	if err != nil {
		*errMessage = C.CString(err.Error())
	}
}

// Write writes a secret containing the contents of `secret` at `path`.
//export Write
func Write(path *C.char, secret *C.char, errMessage **C.char) {
	client, err := secrethub.NewClient()
	if err != nil {
		*errMessage = C.CString(err.Error())
		return
	}
	_, err = client.Secrets().Write(C.GoString(path), []byte(C.GoString(secret)))
	if err != nil {
		*errMessage = C.CString(err.Error())
	}
}

func main() {}
