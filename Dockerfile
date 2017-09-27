FROM ubuntu:16.04

MAINTAINER Pereskokov Vladislav

# Env vars
ENV HOME /root
ENV GOPATH /root/go
ENV PATH /root/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ENV MONGO_URL <MONGO_URL>
RUN mkdir -p /root/.ssh
ADD id_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa
RUN echo "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

# GOPATH
RUN mkdir -p /root/go

RUN apt-get update
RUN apt-get install -y build-essential mercurial git subversion wget curl bzr

# go 1.3 tarball
RUN wget -qO- http://golang.org/dl/go1.8.3.linux-amd64.tar.gz | tar -C /usr/local -xzf -

RUN mkdir -p $GOPATH/src/github.com/vladpereskokov
RUN cd $GOPATH/src/github.com/vladpereskokov/; git clone git@github.com:vladpereskokov/Technopark_HighLoad-nginx.git

WORKDIR /root/go
EXPOSE 8080

CMD cd $GOPATH/src/github.com/vladpereskokov/Technopark_HighLoad-nginx && make run
