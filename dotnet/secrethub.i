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

// Map the time type to System.DateTime.
%apply long long { time };
%typemap(cstype) time "System.DateTime"
%typemap(csvarout, excode=SWIGEXCODE) time %{
    get {
        System.DateTime ret = System.DateTimeOffset.FromUnixTimeSeconds($imcall).UtcDateTime;$excode
        return ret;
    }
%}
%typemap(csvarin, excode=SWIGEXCODE) time %{
    // time is read only
%}

// Map the uuid type to System.Guid.
%apply char* { uuid };
%typemap(cstype) uuid "System.Guid"
%typemap(csvarout, excode=SWIGEXCODE) uuid %{
    get {
        System.Guid ret = System.Guid.Parse($imcall);$excode
        return ret;
    }
%}
%typemap(csvarin, excode=SWIGEXCODE) uuid %{
    // uuids are read only
%}

// Map return value of ResolveEnv to map[string]string.
%typemap(cstype) char* ResolveEnv "System.Collections.Generic.Dictionary<string,string>"
%typemap(csout, excode=SWIGEXCODE) char* ResolveEnv {
    var res = Newtonsoft.Json.JsonConvert.DeserializeObject<System.Collections.Generic.Dictionary<string, string>>($imcall);
    $excode
    return res;
}


// extern struct SecretVersion Read(char* path, char** errMessage);
//extern char* ReadString(char* path, char** errMessage);
/*extern char* Resolve(char* path, char** errMessage);
extern char* ResolveEnv(char** errMessage);
extern bool Exists(char* path, char** errMessage);
extern void Remove(char* path, char** errMessage);
extern void Write(char* path, char* secret, char** errMessage);
*/

extern struct Client {
    int ID;
    %extend {
        Client(char **errMessage);
        ~Client();
        char* ReadString(char* path, char** errMessage);
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
