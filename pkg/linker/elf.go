package linker

import (
	"bytes"
	"unsafe"
)

type Ehdr struct {
	Ident     [16]uint8
	Type      uint16
	Machine   uint16
	Version   uint32
	Entry     uint64
	PhOff     uint64 //段表偏移
	ShOff     uint64 //节表偏移
	Flags     uint32
	EhSize    uint16
	PhEntSize uint16
	PhNum     uint16 //段个数
	ShEntSize uint16
	ShNum     uint16 //节个数
	ShStrndx  uint16 //记录节名称的表在节表中的下标
}

const EhdrSize = int(unsafe.Sizeof(Ehdr{}))

type Shdr struct {
	Name      uint32 //保存节区的名字在shstrtab的下标
	Type      uint32
	Flags     uint64
	Addr      uint64
	Offset    uint64
	Size      uint64
	Link      uint32
	Info      uint32
	AddrAlign uint64
	EntSize   uint64
}

const ShdrSize = int(unsafe.Sizeof(Shdr{}))

func ElfGetName(strTab []byte, offset uint32) string {
	length := uint32(bytes.Index(strTab[offset:], []byte{0}))
	return string(strTab[offset : offset+length])
}

const SymSize = int(unsafe.Sizeof(Sym{}))

type Sym struct {
	Name  uint32
	Info  uint32
	Other uint32
	Shndx uint32
	Val   uint32
	Size  uint32
}
