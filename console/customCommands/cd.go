package customCommands

import (
	"fmt"
	"strings"

	"github.com/horvatic/ratnicka-konzola/console/path"
)

func getParentPath(currentPwd string) string {
	if currentPwd == "/" {
		return currentPwd
	}
	newPwd := strings.TrimSuffix(currentPwd, "/")
	currentDirIndex := strings.LastIndex(newPwd, "/")
	if currentDirIndex == 0 {
		return "/"
	}
	return newPwd[:currentDirIndex]
}

func Cd(userInput string, currentPwd string) string {
	pwd := strings.TrimPrefix(userInput, "cd ")
	pwd = strings.ReplaceAll(pwd, "\\ ", " ")
	if pwd == "." {
		return currentPwd
	} else if pwd == ".." {
		pwd = getParentPath(currentPwd)

	} else if pwd[0] != '/' {
		pwd = currentPwd + pwd
	}

	pwd = path.AddTermination(pwd)

	if path.IsValidPath(pwd) {
		return pwd
	}

	fmt.Printf("Not a valid path: %s", pwd)
	return currentPwd

}
