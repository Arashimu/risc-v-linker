package linker

import (
	_ "bytes"
	"debug/elf"
	_ "encoding/binary"
	"fmt"
	_ "fmt"
	"riscv_link/pkg/utils"
)

type InputFile struct {
	File         *File
	ElfSections  []Shdr
	ElfSyms      []Sym
	FirstGlobal  int64 //第一个global的下标(symtab里先把保存local再保存global)
	ShStrtab     []byte
	SymbolStrtab []byte
}

func NewInputFile(file *File) InputFile {
	f := InputFile{File: file}
	if len(file.Contents) < EhdrSize {
		utils.Fatal("file too small")
	}
	if !CheckMagic(file.Contents) {
		utils.Fatal("not an ELF")
	}
	ehdr := utils.Read[Ehdr](f.File.Contents)
	contents := file.Contents[ehdr.ShOff:] //节表开始
	shdr := utils.Read[Shdr](contents)     //节表第一个元素(节表头) Section Header

	numSections := int64(ehdr.ShNum) //由于ShNum是16位，而节的数量可能很大，无法存下，此时ShNum为0，节的数量需要在节表的第一个节表头的Size字段获取
	if numSections == 0 {
		numSections = int64(shdr.Size)
	}
	f.ElfSections = []Shdr{shdr}
	for numSections > 1 {
		contents = contents[ShdrSize:]
		f.ElfSections = append(f.ElfSections, utils.Read[Shdr](contents))
		numSections--
	}
	shstrndx := int64(ehdr.ShStrndx) //同上，如果节区很多，Shstrtab的下标超过16位的表示范围，那么默认为elf.SHN_XINDEX
	if ehdr.ShStrndx == uint16(elf.SHN_XINDEX) {
		shstrndx = int64(shdr.Link)
	}

	f.ShStrtab = f.GetBytesFromIdx(shstrndx)

	return f

}
func (f *InputFile) GetBytesFromShdr(s *Shdr) []byte {
	end := s.Offset + s.Size
	if uint64(len(f.File.Contents)) < end {
		utils.Fatal(fmt.Sprintf("section header is out of range: %d", s.Offset))
	}
	return f.File.Contents[s.Offset:end]
}

func (f *InputFile) GetBytesFromIdx(idx int64) []byte {
	return f.GetBytesFromShdr(&f.ElfSections[idx])
}
func (f *InputFile) FillUpElfSyms(s *Shdr) {
	bs := f.GetBytesFromShdr(s)
	nums := len(bs) / SymSize
	f.ElfSyms = make([]Sym, 0, nums)
	for nums > 0 {
		f.ElfSyms = append(f.ElfSyms, utils.Read[Sym](bs))
		bs = bs[SymSize:]
		nums--
	}
}
func (f *InputFile) FindSection(ty uint32) *Shdr {
	for i := 0; i < len(f.ElfSections); i++ {
		shdr := &f.ElfSections[i]
		if shdr.Type == ty {
			return shdr
		}
	}
	return nil
}
