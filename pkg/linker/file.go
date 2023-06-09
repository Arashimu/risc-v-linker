package linker

import (
	"os"
	"riscv_link/pkg/utils"
)

type File struct {
	Name     string
	Contents []byte
}

func MustNewFile(filename string) *File {
	contents, err := os.ReadFile(filename)
	utils.MustNo(err)
	return &File{
		Name:     filename,
		Contents: contents,
	}
}
