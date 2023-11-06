package main

import (
	"fmt"
	"ibizi/task1/getMessage"
	"ibizi/task1/stegography"
)

func main() {
	conf := getMessage.Execute()
	err := stegography.GetMessage(conf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
