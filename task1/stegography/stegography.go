package stegography

import (
	"bytes"

	"errors"
	"fmt"
	"ibizi/task1/getMessage"
	"ibizi/task1/putMessage"
	"os"
)

func readMessage(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	return file
}

func readStegocontainer(path string) ([]byte, error) {
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
				message[RowByUtf] += 1 << (7 - countRows%8)
			}
			countRows++
		}
	}
	if stegoContainer[len(stegoContainer)-1] == ' ' {
		RowByUtf := countRows / 8
		if RowByUtf >= len(message) {
			message = append(message, 0)
		}
		message[RowByUtf] += 1 << (7 - countRows%8)
	}
	fmt.Printf("message: %v\n", message)
	return message, nil
}

func GetMessage(conf *getMessage.Config) error {
	message, err := readStegocontainer(conf.StegocontainerPath)
	if err != nil {
		return err
	}

	err = os.WriteFile(conf.MessagePath,
		message, 0777)

	return err
}

func bitRepresentation(mess []byte) (string, int) {
	var res string

	for _, v := range mess {
		tmp := fmt.Sprintf("%b", v)
		for len(tmp) < 8 {
			tmp = "0" + tmp
		}
		res += tmp
	}

	return res, len(res)
}

func PutMessage(conf *putMessage.Config) error {
	container := readMessage(conf.ContainerPath)
	messageToBits, lenMessage := bitRepresentation(readMessage(conf.MessagePath))
	if bytes.Count(container, []byte{'\n'})+1 < len(messageToBits) {
		println(bytes.Count(container, []byte{'\n'}))
		return errors.New("containers is small")
	}

	countRows := 0

	var stego []byte
	for _, v := range container {
		if v == '\n' && countRows < lenMessage {
			if messageToBits[countRows] == '1' {
				stego = append(stego, ' ')
			}
			countRows++
		}
		stego = append(stego, v)
	}
	if countRows < lenMessage && messageToBits[countRows] == '1' {
		stego = append(stego, ' ')
	}
	err := os.WriteFile(conf.StegocontainerPath,
		stego, 0777)

	return err
}
