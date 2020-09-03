CGO_FILES=secrethub.a secrethub.h go.sum
SWIG_FILES=secrethub.cs secrethub_wrap.c secrethubPINVOKE.cs Secret.cs SecretVersion.cs
OUT_FILES=secrethub_wrap.o
MONO_FILES=runme.exe
DOTNET_FILES=secrethub secrethub.deps.json secrethub.dll secrethub.pdb secrethub.runtimeconfig.json
SWIG=swig
CC=gcc
ODIR=./output
DEPS=$(ODIR)/secrethub_wrap.c $(ODIR)/secrethub.h
OBJ=$(ODIR)/secrethub_wrap.o $(ODIR)/secrethub.a

all: client swig compile

.PHONY: client
client: secrethub_wrapper.go
	go build -o output/secrethub.a -buildmode=c-archive secrethub_wrapper.go

.PHONY: swig
swig:
	$(SWIG) -csharp $(ODIR)/secrethub.i

.PHONY: compile
compile: $(DEPS)
	$(CC) -c -O2 -fpic -o $(ODIR)/secrethub_wrap.o $(ODIR)/secrethub_wrap.c
	$(CC) -shared -fPIC $(OBJ) -o $(ODIR)/libsecrethub.so

.PHONY: dotnet
dotnet: client swig compile
	dotnet publish $(ODIR)/secrethub.csproj -o $(ODIR)
	rm -r $(ODIR)/bin $(ODIR)/obj
	$(MAKE) clear
# 	dotnet $(ODIR)/secrethub.dll

.PHONY: mono
mono: client swig compile
	mono-csc -out:$(ODIR)/runme.exe $(ODIR)/*.cs
	$(MAKE) clear
# 	mono ./$(ODIR)/runme.exe

.PHONY: clear
clear:
	rm -f go.sum
	rm -f $(addprefix $(ODIR)/, $(CGO_FILES) $(SWIG_FILES) $(OUT_FILES)) 

.PHONY: clean
clean: clear
	rm -f $(addprefix $(ODIR)/, $(MONO_FILES) $(DOTNET_FILES) libsecrethub.so)