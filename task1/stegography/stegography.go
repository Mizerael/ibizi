package stegography

import (
	"fmt"
	"ibizi/task1/getMessage"
	"ibizi/task1/putMessage"
	"io/fs"
	"os"
)

func readMessage(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return file
}

func readStegocontainer(path string) []byte {
	stegoContainer := readMessage(path)
	countRows := 0
	var message []byte

	for i, x := range stegoContainer {
		if x == '\n' {
			RowByUtf := countRows / 8
			if RowByUtf >= len(message) {
				message = append(message, 0)
			}

			if stegoContainer[i-1] == ' ' {
				message[RowByUtf] += 1 << (countRows % 8)
			}
			countRows++
		}
	}
	if stegoContainer[len(stegoContainer)-1] == ' ' {
		RowByUtf := countRows / 8
		if RowByUtf >= len(message) {
			message = append(message, 0)
		}
		message[RowByUtf] += 1 << (countRows % 8)
	}
	return message
}

func GetMessage(conf *getMessage.Config) error {
	message := readStegocontainer(conf.StegocontainerPath)
	err := os.WriteFile(conf.MessagePath,
		message, fs.FileMode(os.O_WRONLY)|fs.FileMode(os.O_CREATE))
	return err
}

func bitRepresentation(mess []byte) {
	// var res []int
	for _, v := range mess {
		bitString := fmt.Sprintf("%b", v)
		println(bitString)
	}
}

func PutMessage(conf *putMessage.Config) error {
	// container := readMessage(conf.ContainerPath)
	message := readMessage(conf.MessagePath)
	bitRepresentation(message)
	// var stego []byte
	return nil
}
