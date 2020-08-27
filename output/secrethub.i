%module secrethub
%{
#include "secrethub.h"
%}
%include exception.i
#include "secrethub.h"

%exception {
    clearErr();
    $action
    char *errMsg = checkErr();
    if(strlen(errMsg)) {
        SWIG_exception(SWIG_RuntimeError, errMsg);
    }
}
extern char *Read(char*);
extern char *Resolve(char*);
