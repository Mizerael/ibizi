package main

import (
	"fmt"
	"ibizi/task1/getMessage"
)

func main() {
	var conf *getMessage.Config = getMessage.Execute()
	fmt.Println(conf)
}
