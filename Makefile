SHELL = bash
CGO_FILES = Client.a Client.h
SWIG_FILES = Client.cs secrethub_wrap.c ClientPINVOKE.cs Secret.cs SecretVersion.cs
OUT_FILES = secrethub_wrap.o libClient.so Client.dll
DOTNET_DIR = ./dotnet
DEPS = $(DOTNET_DIR)/secrethub_wrap.c $(DOTNET_DIR)/Client.h
OBJ = $(DOTNET_DIR)/secrethub_wrap.o $(DOTNET_DIR)/Client.a

lib: client swig compile
	@echo "Library Ready ^-^"

ifeq ($(OS),Windows_NT)
.PHONY: client
client: secrethub_wrapper.go
	@echo "Making the C library from Go files..."
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o $(DOTNET_DIR)/Client.a -buildmode=c-archive secrethub_wrapper.go

.PHONY: compile
compile: $(DEPS)
	@echo "Compiling using gcc..."
	@x86_64-w64-mingw32-gcc -c -O2 -fpic -o $(DOTNET_DIR)/secrethub_wrap.o $(DOTNET_DIR)/secrethub_wrap.c
	@x86_64-w64-mingw32-gcc -shared -fPIC $(OBJ) -o $(DOTNET_DIR)/Client.dll
else
.PHONY: client
client: secrethub_wrapper.go
	@echo "Making the C library from Go files..."
	@go build -o $(DOTNET_DIR)/Client.a -buildmode=c-archive secrethub_wrapper.go

.PHONY: compile
compile: $(DEPS)
	@echo "Compiling..."
	@gcc -c -O2 -fpic -o $(DOTNET_DIR)/secrethub_wrap.o $(DOTNET_DIR)/secrethub_wrap.c
	@gcc -shared -fPIC $(OBJ) -o $(DOTNET_DIR)/libClient.so
endif

.PHONY: swig
swig:
	@echo "Generating swig files..."
	@swig -csharp -namespace SecretHub $(DOTNET_DIR)/secrethub.i

.PHONY: nupkg
nupkg: lib
	@echo "Making the NuGet Package..."
	@dotnet pack $(DOTNET_DIR)/secrethub.csproj -o $(DOTNET_DIR)/build --nologo
	@mv $(DOTNET_DIR)/build/SecretHub.*.nupkg .
	@make clean
	@echo "NuGet Package Ready ^-^"

#.PHONY: nupkg-publish
#nupkg-publish: nupkg
#	dotnet nuget push *.nupkg --api-key <API_KEY> --source 	https://api.nuget.org/v3/index.json

.PHONY: deps
deps:
	@sudo apt install gcc
	@sudo apt install gcc-mingw-w64

.PHONY: clean
clean:
	@rm -f go.sum
	@rm -f $(addprefix $(DOTNET_DIR)/, $(CGO_FILES) $(SWIG_FILES) $(OUT_FILES))
	@rm -rf $(DOTNET_DIR)/build $(DOTNET_DIR)/bin $(DOTNET_DIR)/obj
