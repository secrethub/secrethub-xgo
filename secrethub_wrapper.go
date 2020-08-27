package secrethub_xgo

import (
	"C"
	"os"
	"strings"

	"github.com/secrethub/secrethub-go/internals/api"
	"github.com/secrethub/secrethub-go/pkg/secrethub"
)

//export Read
func Read(path string) (*api.Secret, error) {
	client, err := secrethub.NewClient()
	if err != nil {
		return nil, err
	}
	secret, err := client.Secrets().Get(path)
	if err != nil {
		return nil, err
	}
	return secret, nil
}

//export Resolve
func Resolve(ref string) ([]byte, error) {
	client, err := secrethub.NewClient()
	if err != nil {
		return nil, err
	}
	bits := strings.Split(ref, "://")
	if len(bits) == 2 && bits[0] == "secrethub" {
		secret, err := client.Secrets().Read(bits[1])
		if err != nil {
			return nil, err
		}
		return secret.Data, nil
	}
	return []byte(ref), nil
}

//export ResolveEnv
func ResolveEnv() (map[string]string, error) {
	env := os.Environ()
	resolvedEnv := make(map[string]string, len(env))
	for _, keyValuePair := range env {
		bits := strings.Split(keyValuePair, "=")
		resolvedVar, err := Resolve(bits[1])
		if err != nil {
			return nil, err
		}
		resolvedEnv[bits[0]] = string(resolvedVar)
	}
	return resolvedEnv, nil
}

// Workaround for error handling.

// errMessage stores the message of the last encountered error.
// If no error was encountered it is "".
var errMessage string

// setErr stores the message of the given error in errMessage,
// so it can be checked from the wrapper C code.
func setErr(err error) {
	errMessage = err.Error()
}

// checkErr returns the message of the last encountered error.
// It is called from the C wrapper code.
//export checkErr
func checkErr() *C.char {
	return C.CString(errMessage)
}

// clearErr clears errMessage.
//export clearErr
func clearErr() {
	errMessage = ""
}
