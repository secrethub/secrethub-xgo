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

struct Client {
	int ID;
};
#include <stdbool.h>
#include <stdlib.h>
*/
import "C"

import (
	"encoding/json"
	"github.com/secrethub/secrethub-go/pkg/secrethub"
	"os"
	"strings"
	"unsafe"
)

var clients = make(map[int]secrethub.ClientInterface)
var nextClientID = 1;

//export new_Client
func new_Client(errMessage **C.char) *C.struct_Client{
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
	cClient := (*C.struct_Client)(C.malloc(C.size_t(unsafe.Sizeof(C.struct_Client{}))))
	cClient.ID = C.int(nextClientID)
	clients[nextClientID] = client
	nextClientID++
	return cClient
}

//export delete_Client
func delete_Client(cClient *C.struct_Client) {
	delete(clients, int(cClient.ID))
	C.free(unsafe.Pointer(cClient))
}

func GetGoClient(cClient *C.struct_Client) secrethub.ClientInterface {
	return clients[int(cClient.ID)]
}


// Client_Read retrieves a secret by its path.
//export Client_Read
func Client_Read(cClient *C.struct_Client, path *C.char, errMessage **C.char) C.struct_SecretVersion {
	client := GetGoClient(cClient)
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

// Client_ReadString retrieves a secret as a string.
//export Client_ReadString
func Client_ReadString(cClient *C.struct_Client, path *C.char, errMessage **C.char) *C.char {
	client := GetGoClient(cClient)
	secret, err := client.Secrets().Read(C.GoString(path))
	if err != nil {
		*errMessage = C.CString(err.Error())
		return nil
	}
	return C.CString(string(secret.Data))
}

// Client_Resolve fetches the values of a secret from SecretHub, when the `ref` parameter
// has the format `secrethub://<path>`. Otherwise it returns `ref` unchanged, as an array of bytes.
//export Client_Resolve
func Client_Resolve(cClient *C.struct_Client, ref *C.char, errMessage **C.char) *C.char {
	lowercaseRef := strings.ToLower(C.GoString(ref))
	prefix := "secrethub://"
	if strings.HasPrefix(lowercaseRef, prefix) {
		client := GetGoClient(cClient)
		secret, err := client.Secrets().Read(strings.TrimPrefix(lowercaseRef, prefix))
		if err != nil {
			*errMessage = C.CString(err.Error())
			return nil
		}
		return C.CString(string(secret.Data))
	}
	return ref
}

// Client_ResolveEnv takes a map of environment variables and replaces the values of those
// which store references of secrets in SecretHub (`secrethub://<path>`) with the value
// of the respective secret. The other entries in the map remain untouched.
//export Client_ResolveEnv
func Client_ResolveEnv(cClient *C.struct_Client, errMessage **C.char) *C.char {
	envVars := os.Environ()
	resolvedEnv := make(map[string]string, len(envVars))
	for _, value := range envVars {
		envVar := strings.SplitN(value, "=", 2)
		if len(envVar) < 2 {
			continue
		}
		key := envVar[0]
		value := C.CString(envVar[1])
		resolvedValue := Client_Resolve(cClient, value, errMessage)
		resolvedEnv[key] = C.GoString(resolvedValue)
	}
	encoding, err := json.Marshal(resolvedEnv)
	if err != nil {
		*errMessage = C.CString(err.Error())
		return nil
	}
	return C.CString(string(encoding))
}

// Client_Exists checks if a secret exists at `path`.
//export Client_Exists
func Client_Exists(cClient *C.struct_Client, path *C.char, errMessage **C.char) C.bool {
	client := GetGoClient(cClient)
	exists, err := client.Secrets().Exists(C.GoString(path))
	if err != nil {
		*errMessage = C.CString(err.Error())
		return C.bool(false)
	}
	return C.bool(exists)
}

// Client_Remove deletes the secret found at `path`, if it exists.
//export Client_Remove
func Client_Remove(cClient *C.struct_Client, path *C.char, errMessage **C.char) {
	client := GetGoClient(cClient)
	err := client.Secrets().Delete(C.GoString(path))
	if err != nil {
		*errMessage = C.CString(err.Error())
		return
	}
}

// Client_Write writes a secret containing the contents of `secret` at `path`.
//export Client_Write
func Client_Write(cClient *C.struct_Client, path *C.char, secret *C.char, errMessage **C.char) {
	client := GetGoClient(cClient)
	_, err := client.Secrets().Write(C.GoString(path), []byte(C.GoString(secret)))
	if err != nil {
		*errMessage = C.CString(err.Error())
		return
	}
}

func main() {}
