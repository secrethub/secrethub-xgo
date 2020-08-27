package secrethub_xgo

import (
	"C"
	"strings"

	"github.com/secrethub/secrethub-go/internals/api"
	"github.com/secrethub/secrethub-go/pkg/secrethub"
)

// Read retrieves a secret from SecretHub, given its path.
// It throws an error if it fails to initialize the SecretHub client
// or if it fails to retrieve the secret.
//export Read
func Read(path string) *api.Secret {
	client, err := secrethub.NewClient()
	if err != nil {
		setErr(err)
		return nil
	}
	secret, err := client.Secrets().Get(path)
	if err != nil {
		setErr(err)
		return nil
	}
	return secret
}

// Resolve fetches the values of a secret from SecretHub, when the `ref` parameter
// has the format `secrethub://<path>`. Otherwise it returns `ref` unchanged, as an array of bytes.
//export Resolve
func Resolve(ref string) []byte {
	client, err := secrethub.NewClient()
	if err != nil {
		setErr(err)
		return nil
	}
	bits := strings.Split(ref, "://")
	if len(bits) == 2 && bits[0] == "secrethub" {
		secret, err := client.Secrets().Read(bits[1])
		if err != nil {
			setErr(err)
			return nil
		}
		return secret.Data
	}
	return []byte(ref)
}

// ResolveEnv takes a map of environment variables and replaces the values of those
// which store references of secrets in SecretHub (`secrethub://<path>`) with the value
// of the respective secret. The other entries in the map remain untouched.
//export ResolveEnv
func ResolveEnv(envVars map[string]string) map[string]string {
	resolvedEnv := make(map[string]string, len(envVars))
	for key, value := range envVars {
		resolvedEnv[key] = string(Resolve(value))
	}
	return resolvedEnv
}