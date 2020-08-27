SWIG_FILES=secrethub.a secrethub.h secrethub.cs secrethub_wrap.c secrethubPINVOKE.cs
SWIG=swig
CC=gcc
CFLAGS=-I.
ODIR=./app
LDIR =./app
DEPS = secrethub.h
OBJ = $(ODIR)/secrethub_wrap.o
LIBS=-lm

$(ODIR)/%.o: %.c $(DEPS)
	$(CC) -c -fpic -o $@ $< $(CFLAGS) -o $@

example: $(OBJ)
	$(CC) -shared -fpic -o $@ $^ secrethub.a $(CFLAGS) -o $(ODIR)/libsecrethub.so

.PHONY: swig
swig:
	$(SWIG) -csharp secrethub.i

.PHONY: clean
clean:
	rm -f $(SWIG_FILES) $(ODIR)/secrethub.o $(ODIR)/secrethub_wrap.o $(ODIR)/libexample.so
