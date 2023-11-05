package main

import (
	"ibizi/task1/putMessage"
	"ibizi/task1/stegography"
)

func main() {
	conf := putMessage.Execute()
	stegography.PutMessage(conf)
}
