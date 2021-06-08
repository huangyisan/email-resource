package main

import (
	"fmt"

	"github.com/huangyisan/email-resource/check"
)

func main() {
	output, err := check.Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
