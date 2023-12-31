# itertools for Go (1.22+)

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.22-%23007d9c)
[![Go Reference](https://pkg.go.dev/badge/github.com/j178/it.svg)](https://pkg.go.dev/github.com/j178/it)
[![CI](https://github.com/j178/it/actions/workflows/ci.yaml/badge.svg)](https://github.com/j178/it/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/j178/it)](https://goreportcard.com/report/github.com/j178/it)
[![codecov](https://codecov.io/gh/j178/it/graph/badge.svg?token=Q0G5O7DF3G)](https://codecov.io/gh/j178/it)

Go1.22 will support [`range over function`](https://github.com/golang/go/issues/61405) and introduce the [`iter`](https://github.com/golang/go/issues/61897) std package.
(Behind the `GOEXPERIMENT=rangefunc` gate.)

See more at [RangefuncExperiment](https://github.com/golang/go/wiki/RangefuncExperiment).

This package provides a set of iterator functions borrowed from Python and Rust.

Some code are copied from Russ Cox's proposals [61898](https://github.com/golang/go/issues/61898), [61899](https://github.com/golang/go/issues/61899) and [61900](https://github.com/golang/go/issues/61900).

## Installation

```bash
go get github.com/j178/it
```

## Usage

To use this package, you need to install at least Go 1.22(not released yet) and set `GOEXPERIMENT=rangefunc`.

```bash
go install golang.org/dl/gotip@latest
gotip download

GOEXPERIMENT=rangefunc gotip run main.go
```
