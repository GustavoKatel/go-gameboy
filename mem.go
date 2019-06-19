package main

import "os"

// GBMem gameboy base ram
type GBMem struct {
	dat []byte
}

// GBMemNew creates a new managed memory
func GBMemNew() *GBMem {
	mem := &GBMem{
		dat: make([]byte, 1024*64), // 64kb
	}

	return mem
}

func (mem *GBMem) Put(pos int, val byte) {
	mem.dat[pos] = val
}

func (mem *GBMem) Get(pos int) byte {
	return mem.dat[pos]
}

func (mem *GBMem) GetUint16(pos int) uint16 {
	upper := uint16(mem.Get(pos))

	lower := uint16(mem.Get(pos + 1))
	lower = lower << 8

	return upper | lower
}

func (mem *GBMem) Dump(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(mem.dat)

	return err
}
