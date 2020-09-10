SHELL = bash
SWIG_VERSION = 4.0.2
CGO_FILES = Client.a Client.h
SWIG_FILES = Client.cs secrethub_wrap.c ClientPINVOKE.cs Secret.cs SecretVersion.cs
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

.PHONY: nupkg
nupkg: lib lib-win
	dotnet pack $(DOTNET_DIR)/secrethub.csproj -o $(DOTNET_DIR)/build --nologo
	mv $(DOTNET_DIR)/build/SecretHub.*.nupkg .
	make clean

#.PHONY: nupkg-publish
#nupkg-publish: nupkg
	#dotnet nuget push *.nupkg --api-key <API_KEY> --source 	https://api.nuget.org/v3/index.json

.PHONY: deps
.ONESHELL:
deps:
	# install gcc
	sudo apt install gcc
	sudo apt install gcc-mingw-w64
	# install pcre
	sudo apt install libpcre3-dev
	# install swig
	sudo apt install swig
# 	wget https://downloads.sourceforge.net/swig/swig-$(SWIG_VERSION).tar.gz
# 	mkdir -p swig && tar -xzvf swig-$(SWIG_VERSION).tar.gz -C swig --strip-components 1
# 	cd swig
# 	./configure; sudo make; sudo make install
# 	cd ..; rm -rf swig; rm -f swig-$(SWIG_VERSION).tar.gz
# 	echo "export SWIG_PATH=usr/local/share/swig/bin" | sudo tee -a /etc/profile
# 	echo "export PATH=$(SWIG_PATH):$(PATH)" | sudo tee -a /etc/profile
# 	source /etc/profile


.PHONY: clean
clean:
	rm -f go.sum
	rm -f $(addprefix $(DOTNET_DIR)/, $(CGO_FILES) $(SWIG_FILES) $(OUT_FILES))
	rm -rf $(DOTNET_DIR)/build $(DOTNET_DIR)/bin $(DOTNET_DIR)/obj
