package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex converts a int64 number into a byte array
func IntToHex(num int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}