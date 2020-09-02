package main

/*
struct Secret {
	char* SecretVersionID;
	int Version;
	char* Data;
	long long CreatedAt;
};
#include <stdbool.h>
*/
import "C"

import (
	"strings"

	"github.com/secrethub/secrethub-go/pkg/secrethub"
)

// Read retrieves a secret by its path.
//export Read
func Read(path *C.char, errMessage **C.char) C.struct_Secret {
	client, err := secrethub.NewClient()
	if err != nil {
		*errMessage = C.CString(err.Error())
		return C.struct_Secret{}
	}
	secret, err := client.Secrets().Read(C.GoString(path))
	if err != nil {
		*errMessage = C.CString(err.Error())
		return C.struct_Secret{}
	}
	return C.struct_Secret{
		SecretVersionID: C.CString(secret.SecretVersionID.String()),
		Version:         C.int(secret.Version),
		Data:            C.CString(string(secret.Data)),
		CreatedAt:       C.longlong(secret.CreatedAt.Unix()),
	}
}

// ReadString retrieves a secret as a string.
//export ReadString
func ReadString(path *C.char, errMessage **C.char) *C.char {
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

/*
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
}*/

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
