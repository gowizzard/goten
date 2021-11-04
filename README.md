# Random token

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/goten.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/goten/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/goten/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/goten)](https://goreportcard.com/report/github.com/gowizzard/goten) [![Go Reference](https://pkg.go.dev/badge/github.com/gowizzard/goten.svg)](https://pkg.go.dev/github.com/gowizzard/goten)  

To create a random token or password in GO Lang.

## Install

First you have to install the package:

```console
go get github.com/gowizzard/goten
```

## How to use

Here is a small example how to create a random password:

```go
token := goten.New(55)
```

