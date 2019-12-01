package path

import (
	"os"
)

func IsValidPath(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func AddTermination(path string) string {
	pwd := path
	if pwd[0] != '/' {
		pwd = "/" + pwd
	}

	if pwd[len(pwd)-1] != '/' {
		pwd = pwd + "/"
	}
	return pwd
}

func InitPath() string {
	pwd, _ := os.Getwd()
	return AddTermination(pwd)
}
