SHELL = bash
CGO_FILES = Client.a Client.h
SWIG_FILES = Client.cs secrethub_wrap.c ClientPINVOKE.cs Secret.cs SecretVersion.cs
OUT_FILES = secrethub_wrap.o libClient.so Client.dll
ODIR = ./dotnet
DEPS = $(ODIR)/secrethub_wrap.c $(ODIR)/Client.h
OBJ = $(ODIR)/secrethub_wrap.o $(ODIR)/Client.a

lib: client swig compile
lib-win: client-win swig compile-win

.PHONY: client
client: secrethub_wrapper.go
	go build -o $(ODIR)/Client.a -buildmode=c-archive secrethub_wrapper.go

.PHONY: client-win
client-win: secrethub_wrapper.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o $(ODIR)/Client.a -buildmode=c-archive secrethub_wrapper.go

.PHONY: swig
swig:
	swig -csharp -namespace SecretHub $(ODIR)/secrethub.i

.PHONY: compile
compile: $(DEPS)
	gcc -c -O2 -fpic -o $(ODIR)/secrethub_wrap.o $(ODIR)/secrethub_wrap.c
	gcc -shared -fPIC $(OBJ) -o $(ODIR)/libClient.so

.PHONY: compile-win
compile-win: $(DEPS)
	x86_64-w64-mingw32-gcc -c -O2 -fpic -o $(ODIR)/secrethub_wrap.o $(ODIR)/secrethub_wrap.c
	x86_64-w64-mingw32-gcc -shared -fPIC $(OBJ) -o $(ODIR)/Client.dll

.PHONY: dotnet-test
dotnet-test: $(ODIR)/libClient.so
	dotnet publish $(ODIR)/secrethub.csproj -o $(ODIR)/build -f netcoreapp3.1 --nologo
	mv $(ODIR)/libClient.so $(ODIR)/build 
# 	dotnet $(ODIR)/build/secrethub.dll

.PHONY: nupkg
nupkg: lib lib-win
	dotnet pack $(ODIR)/secrethub.csproj -o $(ODIR)/build --nologo
	mv $(ODIR)/build/SecretHub.*.nupkg .
	make clean

#.PHONY: nupkg-publish
#nupkg-publish: nupkg
	#dotnet nuget push *.nupkg --api-key <API_KEY> --source 	https://api.nuget.org/v3/index.json

.PHONY: deps
deps:
	sudo apt install gcc
	sudo apt install gcc-mingw-w64

.PHONY: clean
clean:
	rm -f go.sum
	rm -f $(addprefix $(ODIR)/, $(CGO_FILES) $(SWIG_FILES) $(OUT_FILES))
	rm -rf $(ODIR)/build $(ODIR)/bin $(ODIR)/obj
