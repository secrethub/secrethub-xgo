%module secrethub
%{
#include "secrethub.h"
%}
%include exception.i
#include "secrethub.h"
#include <stdio.h>

%typemap(in, numinputs=0) char **errMessage (char *temp="") {
  $1 = &temp;
}

%typemap(argout, canthrow=1) char **errMessage {
  char *errMsg = *$1;
  if(strlen(errMsg) != 0) {
      SWIG_exception(SWIG_RuntimeError, errMsg);
  }
}

extern char *Read(char* path, char** errMessage);
extern char *Resolve(char* path, char** errMessage);
extern bool Exists(char* path, char** errMessage);
extern void Remove(char* path, char** errMessage);
extern void Write(char* path, char* secret, char** errMessage);
