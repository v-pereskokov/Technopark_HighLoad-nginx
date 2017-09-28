# Highload course

## HTTP server

[![Build Status](https://travis-ci.org/vladpereskokov/Technopark_HighLoad-nginx.svg?branch=develop)](https://travis-ci.org/vladpereskokov/Technopark_HighLoad-nginx)  

### Table of contents
  * [Development stack](#dstack)  
  * [Clone](#clone)  
  * [Local Run](#lrun)  
  * [Docker run](#drun)  
  * [Unit tests](#utest)  
  * [Http test suite](#htest)  
  * [Author](#author)  

<a name="dstack"></a>
### Development stack

* Go lang
* Travis CI [tests]

<a name="clone"></a>
### Clone

```bash
  $ git clone [this repo] $GOPATH/github.com/vladpereskokov
```

<a name="lrun"></a>
### Local run

```go
  $ go run ./src/main.go
```  
*or*  
```bash
  $ make run
```

<a name="drun"></a>
### Docker run

```bash
  $ docker build -t [NAME] [THIS REPO OR .]
  $ docker run -p 80:80 --name [COMTAINER NAME] -t [NAME]
```  

<a name="utest"></a>
### Unit tests

```bash
  $ make
```

<a name="htest"></a>
### Http test suite

[Tests repo](https://github.com/init/http-test-suite)  

**All tests passed**  

```bash
  $ ./httptest.py
```

<a name="author"></a>
### Author  
[Pereskokov Vladislav](https://vladpereskokov.github.io/vladislav_pereskokov/)  
