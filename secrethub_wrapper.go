package main

/*
typedef long long time;
typedef char* uuid;

struct Secret {
	uuid SecretID;
	uuid DirID;
	uuid RepoID;
	char* Name;
	char* BlindName;
	int VersionCount;
	int LatestVersion;
	char* Status;
	time CreatedAt;
};

struct SecretVersion {
	uuid SecretVersionID;
	struct Secret Secret;
	int Version;
	char* Data;
	time CreatedAt;
	char* Status;
};

struct MyClient {
	int ID;
};
#include <stdbool.h>
#include <stdlib.h>
*/
import "C"

import (
	"github.com/secrethub/secrethub-go/pkg/secrethub"
	"unsafe"
)

var clients = make(map[int]secrethub.ClientInterface)
var nextClientID = 1;

//export new_MyClient
func new_MyClient(errMessage **C.char) *C.struct_MyClient{
	options := []secrethub.ClientOption{
		secrethub.WithAppInfo(&secrethub.AppInfo{
			Name:    "secrethub-xgo",
			Version: "0.1.0",
		}),
	}
	client, err := secrethub.NewClient(options...)
	if err != nil {
		*errMessage = C.CString(err.Error())
		return nil
	}
	cClient := (*C.struct_MyClient)(C.malloc(C.size_t(unsafe.Sizeof(C.struct_MyClient{}))))
	cClient.ID = C.int(nextClientID)
	clients[nextClientID] = client
	nextClientID++
	return cClient
}

//export delete_MyClient
func delete_MyClient(cClient *C.struct_MyClient) {
	delete(clients, int(cClient.ID))
	C.free(unsafe.Pointer(cClient))
}

func GetGoClient(cClient *C.struct_MyClient) secrethub.ClientInterface {
	return clients[int(cClient.ID)]
}

/*
// Read retrieves a secret by its path.
//export Read
func Read(path *C.char, errMessage **C.char) C.struct_SecretVersion {
	client, err := Client()
	if err != nil {
		*errMessage = C.CString(err.Error())
		return C.struct_SecretVersion{}
	}
	secret, err := client.Secrets().Read(C.GoString(path))
	if err != nil {
		*errMessage = C.CString(err.Error())
		return C.struct_SecretVersion{}
	}
	return C.struct_SecretVersion{
		SecretVersionID: C.CString(secret.SecretVersionID.String()),
		Secret: C.struct_Secret{
			SecretID:      C.CString(secret.Secret.SecretID.String()),
			DirID:         C.CString(secret.Secret.DirID.String()),
			RepoID:        C.CString(secret.Secret.RepoID.String()),
			Name:          C.CString(secret.Secret.Name),
			BlindName:     C.CString(secret.Secret.BlindName),
			VersionCount:  C.int(secret.Secret.VersionCount),
			LatestVersion: C.int(secret.Secret.LatestVersion),
			Status:        C.CString(secret.Secret.Status),
			CreatedAt:     C.longlong(secret.Secret.CreatedAt.Unix()),
		},
		Version:   C.int(secret.Version),
		Data:      C.CString(string(secret.Data)),
		CreatedAt: C.longlong(secret.CreatedAt.Unix()),
		Status:    C.CString(secret.Status),
	}
}
*/
// ReadString retrieves a secret as a string.
//export MyClient_ReadString
func MyClient_ReadString(cClient *C.struct_MyClient, path *C.char, errMessage **C.char) *C.char {
	client := GetGoClient(cClient)
	secret, err := client.Secrets().Read(C.GoString(path))
	if err != nil {
		*errMessage = C.CString(err.Error())
		return nil
	}
	return C.CString(string(secret.Data))
}
/*
// Resolve fetches the values of a secret from SecretHub, when the `ref` parameter
// has the format `secrethub://<path>`. Otherwise it returns `ref` unchanged, as an array of bytes.
//export Resolve
func Resolve(ref *C.char, errMessage **C.char) *C.char {
	lowercaseRef := strings.ToLower(C.GoString(ref))
	prefix := "secrethub://"
	if strings.HasPrefix(lowercaseRef, prefix) {
		client, err := Client()
		if err != nil {
			*errMessage = C.CString(err.Error())
			return nil
		}
		secret, err := client.Secrets().Read(strings.TrimPrefix(lowercaseRef, prefix))
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
		envVar := strings.SplitN(value, "=", 2)
		if len(envVar) < 2 {
			continue
		}
		key := envVar[0]
		value := C.CString(envVar[1])
		resolvedValue := Resolve(value, errMessage)
		resolvedEnv[key] = C.GoString(resolvedValue)
	}
	encoding, err := json.Marshal(resolvedEnv)
	if err != nil {
		*errMessage = C.CString(err.Error())
		return nil
	}
	return C.CString(string(encoding))
}

// Exists checks if a secret exists at `path`.
//export Exists
func Exists(path *C.char, errMessage **C.char) C.bool {
	client, err := Client()
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
	client, err := Client()
	if err != nil {
		*errMessage = C.CString(err.Error())
		return
	}
	err = client.Secrets().Delete(C.GoString(path))
	if err != nil {
		*errMessage = C.CString(err.Error())
		return
	}
}

// Write writes a secret containing the contents of `secret` at `path`.
//export Write
func Write(path *C.char, secret *C.char, errMessage **C.char) {
	client, err := Client()
	if err != nil {
		*errMessage = C.CString(err.Error())
		return
	}
	_, err = client.Secrets().Write(C.GoString(path), []byte(C.GoString(secret)))
	if err != nil {
		*errMessage = C.CString(err.Error())
		return
	}
}
*/
func main() {}
