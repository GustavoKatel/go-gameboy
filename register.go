package main

import "strconv"

// Register16 16-bit register abstraction
type Register16 uint16

// Register16_Index
type Register16_Index uint8

const (
	Register16_Index_LEFT  Register16_Index = 8
	Register16_Index_RIGHT Register16_Index = 0
)

func (reg Register16) GetByte(index Register16_Index) byte {
	value := reg >> index
	return uint8(value)
}

func (reg Register16) PutByte(index Register16_Index, val byte) Register16 {
	mask := uint16(0xff)

	if index == Register16_Index_RIGHT {
		mask = mask << 8
	}

	v16 := uint16(val)

	if index == Register16_Index_LEFT {
		v16 = v16 << 8
	}

	ret := uint16(reg)
	ret = ret & mask
	return Register16(ret | v16)
}

func (reg Register16) HexString() string {
	return strconv.FormatUint(uint64(reg), 16)
}
