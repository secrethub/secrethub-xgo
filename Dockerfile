from mono:latest

ENV SWIG_VERSION 4.0.2
ENV GO_VERSION 1.15

# Install gcc
RUN apt-get update && \
    printf "y" | apt-get install build-essential

# Install PCRE
RUN apt-get install -y libpcre3-dev

# Install swig
RUN apt-get update && printf "y" | apt-get install wget && \
    wget https://downloads.sourceforge.net/swig/swig-$SWIG_VERSION.tar.gz && \
    mkdir swig && tar -xzvf swig-$SWIG_VERSION.tar.gz -C swig --strip-components 1

WORKDIR /swig/
RUN ./configure --prefix=/opt/swig --without-maximum-compile-warnings && \
	make && make install

WORKDIR /
RUN rm -rf swig swig-$SWIG_VERSION.tar.gz

ENV SWIG_DIR /opt/swig/share/swig/$SWIG_VERSION/
ENV SWIG_EXECUTABLE /opt/swig/bin/swig
ENV PATH $PATH:/opt/swig/bin/

# Install Go
RUN wget https://dl.google.com/go/go$GO_VERSION.linux-amd64.tar.gz && \
	tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz && \
	rm go$GO_VERSION.linux-amd64.tar.gz

ENV GOROOT /usr/local/go/
ENV PATH $PATH:$GOROOT/bin/

# Prepare build
RUN mkdir build
COPY error_handling.go secrethub_wrapper.go go.mod /build/
COPY output /build/output/
WORKDIR /build/output

