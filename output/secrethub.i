%module secrethub
%{
#include "secrethub.h"
%}
%include exception.i
#include "secrethub.h"

%typemap(in, numinputs=0) char **errMessage (char *temp="") {
  $1 = &temp;
}

%typemap(argout, canthrow=1) char **errMessage {
  char *errMsg = *$1;
  if(strlen(errMsg) != 0) {
      SWIG_exception(SWIG_RuntimeError, errMsg);
  }
}

%typemap(cstype) long long Secret::CreatedAt "System.DateTime"
%typemap(csvarout, excode=SWIGEXCODE) long long Secret::CreatedAt %{
    get {
        System.DateTime ret = System.DateTimeOffset.FromUnixTimeSeconds($imcall).UtcDateTime;$excode
        return ret;
    }
%}
%typemap(csvarin, excode=SWIGEXCODE) long long Secret::CreatedAt %{
    // Secret.CreatedAt is read only
%}

%typemap(cstype) char* Secret::SecretVersionID "System.Guid"
%typemap(csvarout, excode=SWIGEXCODE) char* Secret::SecretVersionID %{
    get {
        System.Guid ret = System.Guid.Parse($imcall);$excode
        return ret;
    }
%}
%typemap(csvarin, excode=SWIGEXCODE) char* Secret::SecretVersionID %{
    // Secret.SecretVersionID is read only
%}

extern struct Secret Read(char* path, char** errMessage);
extern char* ReadString(char* path, char** errMessage);
extern char* Resolve(char* path, char** errMessage);
extern bool Exists(char* path, char** errMessage);
extern void Remove(char* path, char** errMessage);
extern void Write(char* path, char* secret, char** errMessage);

extern struct Secret {
	char* SecretVersionID;
    int Version;
    char* Data;
    long long CreatedAt;
	char* Status;
};
