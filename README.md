# Random token

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/goten.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/goten/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/goten/actions/workflows/go.yml) [![CodeQL](https://github.com/gowizzard/goten/actions/workflows/codeql.yml/badge.svg)](https://github.com/gowizzard/goten/actions/workflows/codeql.yml) [![CompVer](https://github.com/gowizzard/goten/actions/workflows/compver.yml/badge.svg)](https://github.com/gowizzard/goten/actions/workflows/compver.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/goten)](https://goreportcard.com/report/github.com/gowizzard/goten) [![Go Reference](https://pkg.go.dev/badge/github.com/gowizzard/goten.svg)](https://pkg.go.dev/github.com/gowizzard/goten) [![GitHub issues](https://img.shields.io/github/issues/gowizzard/goten)](https://github.com/gowizzard/goten/issues) [![GitHub forks](https://img.shields.io/github/forks/gowizzard/goten)](https://github.com/gowizzard/goten/network) [![GitHub stars](https://img.shields.io/github/stars/gowizzard/goten)](https://github.com/gowizzard/goten/stargazers) [![GitHub license](https://img.shields.io/github/license/gowizzard/goten)](https://github.com/gowizzard/goten/blob/master/LICENSE)  

To create a random token or password in GO. Based on letters, numbers and symbols. Create via Seed & UnixNano with Intn.

## Install

First you have to install the package. You can do this as follows:

```console
go get github.com/gowizzard/goten
```

## How to use

Here is a small example how to create a random password without numbers and symbols. Only letters.

```go
token := goten.Generate(50, nil)
log.Println(token)
```

And here is a small example how to create a random token with numbers and symbols.

```go
options := goten.Options{
    Uppercase:  true,
    Numbers:    true,
    Symbols:    true,
}

token := goten.Generate(50, &options)
log.Println(token)
```


