package main

import (
	"fmt"
	"ibizi/task1/putMessage"
	"ibizi/task1/stegography"
)

func main() {
	conf := putMessage.Execute()
	err := stegography.PutMessage(conf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
