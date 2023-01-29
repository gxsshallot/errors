# errors

[![status](https://github.com/gaoxiaosong/errors/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gaoxiaosong/errors/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gaoxiaosong/errors/branch/master/graph/badge.svg?token=AOXNUDXAS7)](https://codecov.io/gh/gaoxiaosong/errors)
[![gover](https://img.shields.io/badge/Go-v1.2+-blue)](https://go.dev/)
[![godoc](https://pkg.go.dev/badge/github.com/gaoxiaosong/errors?status.svg)](https://pkg.go.dev/github.com/gaoxiaosong/errors)
[![apache](https://img.shields.io/badge/License-Apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0)

这是一个用于产生错误信息对象的库。

* 支持记录主业务的错误码和错误消息。
* 支持记录子业务的错误码和错误消息。
* 支持全局的错误码和错误消息的映射关系。
* 支持使用`NewXXX`方法时生成调用栈信息。
* 支持配置全局开启/禁用调用栈信息。
* 兼容Go 1.2+版本。

首先，需要导入包名到程序中：

```go
import "github.com/gaoxiaosong/errors"
```

(可选) 全局开启/禁用调用栈信息:

```go
errors.EnableStack = true/false
```

(可选) 设置全局的错误码和错误消息的映射关系：

```go
const (
    CodeError1 = 1000
    ...
    SubCodeError11 = 10001
    ...
)

errors.GlobalCodes.Add(CodeError1, "err1")
errors.GlobalSubCodes.Add(SubCodeError11, "suberr1")

// 如果需要删除的话
errors.GlobalCodes.Del(CodeError1)
errors.GlobalSubCodes.Del(SubCodeError11)

// 可以直接获取错误码对应的错误消息
msg := errors.GlobalCodes.Get(CodeError1)
subMsg := errors.GlobalSubCodes.Get(SubCodeError11)
```

在程序中使用`NewXXX`方法，生成`error`对象：

```go
// 从全局映射关系中，自动生成错误消息
err = errors.New(CodeError1)
err = errors.NewSub(CodeError1, SubCodeError11)

// 自定义错误消息
err = errors.Newf(CodeError1, "c1"),
err = errors.Newf(CodeError1, "%s %d", arg1, arg2)
err = errors.NewSubLevel(CodeError1, "c1", SubCodeError11, "s1")

// 使用其他的error对象作为错误消息
err = errors.NewWithError(CodeError1, otherError)

// 全部参数支持
err = errors.NewFull(CodeError1, "c1", enableStack)
err = errors.NewSubLevelFull(CodeError1, msg, SubCodeError11, "s1", enableStack)

// 设置/获取附加数据信息
err.Attach["key"] = "value"
v := err.Attach["key"]
```

error对象判断/还原为errorBase对象:

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
