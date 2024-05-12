package util

import (
	"path"
	"runtime"
)

func FullFuncName(skip int) (funcName string) {
	pc, _, _, _ := runtime.Caller(skip + 1) //nolint:dogsled
	return runtime.FuncForPC(pc).Name()
}

func FuncName(skip int) (funcName string) {
	return path.Base(FullFuncName(skip + 1))
}
