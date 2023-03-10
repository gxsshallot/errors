# errors

[![status](https://github.com/gxsshallot/errors/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gxsshallot/errors/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gxsshallot/errors/branch/master/graph/badge.svg?token=AOXNUDXAS7)](https://codecov.io/gh/gxsshallot/errors)
[![gover](https://img.shields.io/badge/Go-v1.2+-blue)](https://go.dev/)
[![godoc](https://pkg.go.dev/badge/github.com/gxsshallot/errors?status.svg)](https://pkg.go.dev/github.com/gxsshallot/errors)
[![apache](https://img.shields.io/badge/License-Apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0)

[中文文档](README_cn.md)

This package is used for generating error.

* Support main level code and message.
* Support sub level code and message.
* Support global main level codes and sub level codes map.
* Support logging runtime stack when you call `NewXXX` method.
* Support option to enable or disable runtime stack logging globally.
* Compatible with go version from 1.2+.

You should import the package in your source file first:

```go
import "github.com/gxsshallot/errors"
```

(optional) Enable/Disable the runtime stack option globally:

```go
errors.EnableStack = true/false
```

(optional) Set the global main level codes or sub level codes map:

```go
const (
    CodeError1 = 1000
    ...
    SubCodeError11 = 10001
    ...
)

errors.GlobalCodes.Add(CodeError1, "err1")
errors.GlobalSubCodes.Add(SubCodeError11, "suberr1")

// delete if you want
errors.GlobalCodes.Del(CodeError1)
errors.GlobalSubCodes.Del(SubCodeError11)

// you can get error message manuanlly
msg := errors.GlobalCodes.Get(CodeError1)
subMsg := errors.GlobalSubCodes.Get(SubCodeError11)
```

Use `NewXXX` method in program as `error` interface:

```go
// auto generate message/submessage from global map
err = errors.New(CodeError1)
err = errors.NewSub(CodeError1, SubCodeError11)

// use custom message/submessage instead of global map
err = errors.Newf(CodeError1, "c1"),
err = errors.Newf(CodeError1, "%s %d", arg1, arg2)
err = errors.NewSubLevel(CodeError1, "c1", SubCodeError11, "s1")

// use other error object as message
err = errors.NewWithError(CodeError1, otherError)

// full support
err = errors.NewFull(CodeError1, "c1", enableStack)
err = errors.NewSubLevelFull(CodeError1, msg, SubCodeError11, "s1", enableStack)

// set/get attach info
err.Attach["key"] = "value"
v := err.Attach["key"]
```

Is/Revert error object to errorBase:

```go
func XXX(err error) {
    ok := errors.Is(err)
    ...
    errObj, ok := errors.Revert(err)
    if ok {
        ...
    } else {
        ...
    }
}
```