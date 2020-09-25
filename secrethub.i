%module SecretHubXGO
%{
#include "Client.h"
%}
%include exception.i
#include "Client.h"

// Handle error message output parameters.
%typemap(in, numinputs=0) char **errMessage (char *temp="") {
  $1 = &temp;
}

%typemap(argout, canthrow=1) char **errMessage {
  char *errMsg = *$1;
  if(strlen(errMsg) != 0) {
      SWIG_exception(SWIG_RuntimeError, errMsg);
  }
}

extern struct Client {
    %extend {
        Client(char **errMessage);
        ~Client();
        struct SecretVersion Read(char* path, char** errMessage);
        char* ReadString(char* path, char** errMessage);
        char* Resolve(char* path, char** errMessage);
        char* ResolveEnv(char** errMessage);
        bool Exists(char* path, char** errMessage);
        void Remove(char* path, char** errMessage);
        void Write(char* path, char* secret, char** errMessage);
    }
};

extern struct Secret {
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

extern struct SecretVersion {
	uuid SecretVersionID;
	struct Secret Secret;
	int Version;
	char* Data;
	time CreatedAt;
	char* Status;
};
