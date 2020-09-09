SHELL = bash
CGO_FILES = Client.a Client.h go.sum
SWIG_FILES = Client.cs secrethub_wrap.c ClientPINVOKE.cs Secret.cs SecretVersion.cs
OUT_FILES = secrethub_wrap.o
MONO_FILES = runme.exe
DOTNET_FILES = secrethub secrethub.deps.json secrethub.dll secrethub.pdb secrethub.runtimeconfig.json
DOTNET_DIRS = $(ODIR)/bin $(ODIR)/obj
SWIG = swig
CC = gcc
ODIR = ./output
DEPS = $(ODIR)/secrethub_wrap.c $(ODIR)/Client.h
OBJ = $(ODIR)/secrethub_wrap.o $(ODIR)/Client.a

lib: client swig compile
lib-win: client-win swig compile-win

.PHONY: client
client: secrethub_wrapper.go
	go build -o output/Client.a -buildmode=c-archive secrethub_wrapper.go

.PHONY: client-win
client-win: secrethub_wrapper.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o output/Client.a -buildmode=c-archive secrethub_wrapper.go

.PHONY: swig
swig:
	$(SWIG) -csharp -namespace SecretHub $(ODIR)/secrethub.i

.PHONY: compile
compile: $(DEPS)
	$(CC) -c -O2 -fpic -o $(ODIR)/secrethub_wrap.o $(ODIR)/secrethub_wrap.c
	$(CC) -shared -fPIC $(OBJ) -o $(ODIR)/libClient.so

.PHONY: compile-win
compile-win: $(DEPS)
	x86_64-w64-mingw32-gcc -c -O2 -fpic -o $(ODIR)/secrethub_wrap.o $(ODIR)/secrethub_wrap.c
	x86_64-w64-mingw32-gcc -shared -fPIC $(OBJ) -o $(ODIR)/libClient.dll

.PHONY: dotnet-test
dotnet: $(ODIR)/libClient.so
	dotnet publish $(ODIR)/secrethub.csproj -o $(ODIR)
	rm -r $(ODIR)/bin $(ODIR)/obj
# 	dotnet $(ODIR)/secrethub.dll

.PHONY: mono-test
mono: $(ODIR)/libClient.so
	mono-csc -out:$(ODIR)/runme.exe $(ODIR)/*.cs
# 	mono ./$(ODIR)/runme.exe

.PHONY: nupkg
nupkg: lib lib-win
	mkdir nuget
	cp $(ODIR)/{libClient.dll,libClient.so,Secret.cs,Client.cs,ClientPINVOKE.cs,SecretVersion.cs,secrethub.csproj} ./nuget/
	dotnet pack nuget/secrethub.csproj
	mv ./nuget/bin/Debug/SecretHub.*.nupkg .
	rm -r ./nuget
	rm $(ODIR)/libClient.so
	make clean
#.PHONY: nupkg-publish
#nupkg-publish: nupkg
	#dotnet nuget push $(ODIR)/*.nupkg --api-key <API_KEY> --source 	https://api.nuget.org/v3/index.json

.PHONY: clean
clean:
	rm -f go.sum
	rm -f $(addprefix $(ODIR)/, $(CGO_FILES) $(SWIG_FILES) $(OUT_FILES)) 
	rm -f $(addprefix $(ODIR)/, $(MONO_FILES) $(DOTNET_FILES) libClient.so libClient.dll)
