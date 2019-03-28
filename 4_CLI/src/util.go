package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"log"
	"os"
)

//Int2Byte int64转[]byte
func Int2Byte(num int64) (ret []byte) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	ret = buff.Bytes()

	return
}

//Serialize 序列化结构体
func Serialize(v interface{}) (ret []byte) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(v)
	if err != nil {
		log.Panic("序列化失败！")
	}

	ret = result.Bytes()
	return

}

//Deserialize 反序列化结构体
func Deserialize(data []byte, v interface{}) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(v)
	if err != nil {
		log.Panic("反序列化失败！", err)
	}

	return
}

func CreateFileIfNotExist(path string) (err error) {
	_, err = os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			log.Printf("mkdir failed![%v]\n", err)
			return err
		}
		defer f.Close()
		return nil
	}

	return nil
}
