package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func ToBytes(i interface{}) []byte {
	var b bytes.Buffer
	encoder := gob.NewEncoder(&b)
	err := encoder.Encode(i)
	HandleErr(err)
	return b.Bytes()
}

func FromBytes(i interface{}, data []byte) {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
}
