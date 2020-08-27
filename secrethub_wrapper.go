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
