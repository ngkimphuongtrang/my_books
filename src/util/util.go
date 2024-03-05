package util

import (
	"path"
	"path/filepath"
	"runtime"
)

var (
	_, thisFilePath, _, _ = runtime.Caller(0)
	thisDirPath           = filepath.Dir(thisFilePath)
	projectRoot           = path.Join(thisDirPath, "..", "..")
)

func GetProjectRoot() string {
	return projectRoot
}
