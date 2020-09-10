SHELL = bash
CGO_FILES = Client.a Client.h
SWIG_FILES = Client.cs secrethub_wrap.c ClientPINVOKE.cs Secret.cs SecretVersion.cs
TEST_FILES = Client.cs ClientPINVOKE.cs Secret.cs SecretVersion.cs
OUT_FILES = secrethub_wrap.o libClient.so Client.dll
DOTNET_DIR = ./dotnet
DEPS = $(DOTNET_DIR)/secrethub_wrap.c $(DOTNET_DIR)/Client.h
OBJ = $(DOTNET_DIR)/secrethub_wrap.o $(DOTNET_DIR)/Client.a

lib: client swig compile
lib-win: client-win swig compile-win

.PHONY: client
client: secrethub_wrapper.go
	go build -o $(DOTNET_DIR)/Client.a -buildmode=c-archive secrethub_wrapper.go

.PHONY: client-win
client-win: secrethub_wrapper.go
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o $(DOTNET_DIR)/Client.a -buildmode=c-archive secrethub_wrapper.go

.PHONY: swig
swig:
	swig -csharp -namespace SecretHub $(DOTNET_DIR)/secrethub.i

.PHONY: compile
compile: $(DEPS)
	gcc -c -O2 -fpic -o $(DOTNET_DIR)/secrethub_wrap.o $(DOTNET_DIR)/secrethub_wrap.c
	gcc -shared -fPIC $(OBJ) -o $(DOTNET_DIR)/libClient.so

.PHONY: compile-win
compile-win: $(DEPS)
	x86_64-w64-mingw32-gcc -c -O2 -fpic -o $(DOTNET_DIR)/secrethub_wrap.o $(DOTNET_DIR)/secrethub_wrap.c
	x86_64-w64-mingw32-gcc -shared -fPIC $(OBJ) -o $(DOTNET_DIR)/Client.dll

.PHONY: dotnet-test
dotnet-test: lib lib-win
	cp $(addprefix $(DOTNET_DIR)/, $(TEST_FILES)) $(DOTNET_DIR)/test
	dotnet publish $(DOTNET_DIR)/test/secrethub.csproj -o $(DOTNET_DIR)/build --nologo
	mv $(DOTNET_DIR)/libClient.so $(DOTNET_DIR)/build 
	dotnet test $(DOTNET_DIR)/build/secrethub.dll --nologo
	make clean

.PHONY: nupkg
nupkg: lib lib-win
	dotnet pack $(DOTNET_DIR)/secrethub.csproj -o $(DOTNET_DIR)/build --nologo
	mv $(DIR)/build/SecretHub.*.nupkg .
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
	rm -f $(addprefix $(DOTNET_DIR)/, $(CGO_FILES) $(SWIG_FILES) $(OUT_FILES)) $(addprefix $(DOTNET_DIR)/test/, $(TEST_FILES))
	rm -rf $(addprefix $(DOTNET_DIR)/, build bin obj test/bin test/obj)
