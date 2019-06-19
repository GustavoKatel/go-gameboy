package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {

	mem := GBMemNew()

	romFile, err := os.Open("etc/boot.bin")
	if err != nil {
		panic(err)
	}

	romData, err := ioutil.ReadAll(romFile)
	if err != nil {
		panic(err)
	}

	for i, b := range romData {
		mem.Put(i, b)
	}

	romFile.Close()

	cpu := GBCpuNew(mem)

	for {
		fmt.Printf("PC: %s\n", strconv.FormatUint(uint64(cpu.PC()), 16))
		fmt.Printf("SP: %s\n", strconv.FormatUint(uint64(cpu.SP()), 16))
		fmt.Printf("OP: %s\n", strconv.FormatUint(uint64(mem.Get(int(cpu.PC()))), 16))
		fmt.Printf("---------\n")

		cpu.Step()

		<-time.After(1 * time.Second)
	}
}
