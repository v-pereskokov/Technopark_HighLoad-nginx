FROM ubuntu:16.04

MAINTAINER Pereskokov Vladislav

# Env vars
ENV HOME /root
ENV GOPATH /root/go
ENV PATH /root/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

# GOPATH
RUN mkdir -p /root/go

RUN apt-get update
RUN apt-get install -y build-essential mercurial git subversion wget curl bzr

# go 1.8.3
RUN wget -qO- http://golang.org/dl/go1.8.3.linux-amd64.tar.gz | tar -C /usr/local -xzf -

RUN mkdir -p $GOPATH/src/github.com/vladpereskokov
RUN cd $GOPATH/src/github.com/vladpereskokov/; git clone https://github.com/vladpereskokov/Technopark_HighLoad-nginx.git
RUN cd $GOPATH/src/github.com/vladpereskokov/Technopark_HighLoad-nginx; git clone https://github.com/init/http-test-suite.git; mv ./http-test-suite/httptest ./

WORKDIR /root/go
EXPOSE 80

CMD cd $GOPATH/src/github.com/vladpereskokov/Technopark_HighLoad-nginx && make run
