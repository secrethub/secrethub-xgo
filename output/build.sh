go build -o secrethub.a -buildmode=c-archive ../secrethub_wrapper.go
swig -csharp ./secrethub.i
gcc -O2 -fPIC -c secrethub_wrap.c
gcc -shared  secrethub_wrap.o secrethub.a  -o libsecrethub.so
mono-csc -out:runme.exe *.cs
rm secrethub.a secrethub_wrap.o secrethub_wrap.c secrethub.h secrethub.cs secrethubPINVOKE.cs SecretVersion.cs
