package main

// GBCpu gameboy cpu abstraction
type GBCpu struct {
	sp  Register16
	pc  Register16
	mem *GBMem

	AF Register16
	BC Register16
	DE Register16
	HL Register16

	ticks uint64
}

// GBCpuNew creates a new gameboy cpu
func GBCpuNew(mem *GBMem) *GBCpu {
	cpu := &GBCpu{
		pc:  0x0,
		mem: mem,

		ticks: 0,
	}

	return cpu
}

func (cpu *GBCpu) PC() uint16 {
	return uint16(cpu.pc)
}

func (cpu *GBCpu) SP() uint16 {
	return uint16(cpu.sp)
}

func (cpu *GBCpu) Step() {
	op := cpu.fetch()
	cpu.exec(op)
}

func (cpu *GBCpu) fetch() uint16 {
	// TODO: fetch CB prefix
	pos := int(cpu.pc)
	cpu.pc += 0x01
	return uint16(cpu.mem.Get(pos))
}

func (cpu *GBCpu) exec(op uint16) {
	switch op {
	case 0x0000:
		cpu.ticks += 4
	case 0x0001:
		cpu.ticks += 12
		cpu.BC = Register16(cpu.mem.GetUint16(int(cpu.pc)))
		cpu.pc += 2
	case 0x0002:
		cpu.ticks += 8
		pos := cpu.mem.Get(int(cpu.BC))
		cpu.mem.Put(int(pos), cpu.AF.GetByte(Register16_Index_LEFT))
	case 0x0003:
		cpu.ticks += 8
		cpu.BC += 0x1
	case 0x0004:
		cpu.ticks += 4
		b := cpu.BC.GetByte(Register16_Index_LEFT)
		b += 0x1
		cpu.BC.PutByte(Register16_Index_LEFT, b)

	case 0x0031:
		cpu.ticks += 12
		cpu.sp = Register16(cpu.mem.GetUint16(int(cpu.pc)))
		cpu.pc += 2
	}
}
