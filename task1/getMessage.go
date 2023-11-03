package main

import (
	"ibizi/task1/getMessage"
	"ibizi/task1/stegography"
)

func main() {
	var conf *getMessage.Config = getMessage.Execute()

	stegography.GetMessage(conf)
}
