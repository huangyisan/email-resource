package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/huangyisan/email-resource/out"
)

var (
	//VERSION -
	VERSION string
)

func main() {
	VERSION = "v1.1.0"
	sourceRoot := os.Args[1]
	inbytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("%s - Error reading stdin - %s", VERSION, err.Error()))
		os.Exit(1)
	}
	output, err := out.Execute(sourceRoot, VERSION, inbytes)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("%s - Error during execute - %s", VERSION, err.Error()))
		os.Exit(1)
	}
	fmt.Println(output)
}
