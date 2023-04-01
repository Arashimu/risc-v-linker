package utils

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func Fatal(v any) {
	fmt.Printf("main \033[0;1;31m fatal:\033[0m %v\b", v)
	os.Exit(1)
}

func MustNo(err error) {
	if err != nil {
		Fatal(err)
	}
}

func Read[T any](data []byte) (val T) {
	reader := bytes.NewReader(data)
	err := binary.Read(reader, binary.LittleEndian, &val) //读取val类型T大小的字节
	MustNo(err)
	return val
}
