package logger

import (
	"path"
	"runtime"
	"strings"
)

type callInfo struct {
	packageFullName string
	packageName     string
	fileName        string
	funcName        string
	line            int
}

func retrieveCallInfo() *callInfo {
	pc, file, line, _ := runtime.Caller(2)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageFullName := ""
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageFullName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageFullName = strings.Join(parts[0:pl-1], ".")
	}

	segs := strings.Split(packageFullName, "/")
	packageName = segs[len(segs)-1]

	return &callInfo{
		packageFullName: packageFullName,
		packageName:     packageName,
		fileName:        fileName,
		funcName:        funcName,
		line:            line,
	}
}
