package gbc

const (
	A = iota
	B
	C
	D
	E
	H
	L
	F
)

const (
	AF = iota
	BC
	DE
	HL
	HLI
	HLD
	SP
	PC
)

// Register Z80
type Register struct {
	R   [8]byte
	SP  uint16
	PC  uint16
	IME bool
}

func (r *Register) R16(i int) uint16 {
	switch i {
	case AF:
		return r.AF()
	case BC:
		return r.BC()
	case DE:
		return r.DE()
	case HL:
		return r.HL()
	case HLD:
		hl := r.HL()
		r.setHL(hl - 1)
		return hl
	case HLI:
		hl := r.HL()
		r.setHL(hl + 1)
		return hl
	case SP:
		return r.SP
	case PC:
		return r.PC
	}
	panic("invalid register16")
}

func (r *Register) setR16(i int, val uint16) {
	switch i {
	case AF:
		r.setAF(val)
	case BC:
		r.setBC(val)
	case DE:
		r.setDE(val)
	case HL:
		r.setHL(val)
	case SP:
		r.SP = val
	case PC:
		r.PC = val
	}
}

func (r *Register) AF() uint16 {
	return (uint16(r.R[A]) << 8) | uint16(r.R[F])
}
func (r *Register) setAF(value uint16) {
	r.R[A], r.R[F] = byte(value>>8), byte(value)
}

func (r *Register) BC() uint16 {
	return (uint16(r.R[B]) << 8) | uint16(r.R[C])
}
func (r *Register) setBC(value uint16) {
	r.R[B], r.R[C] = byte(value>>8), byte(value)
}

func (r *Register) DE() uint16 {
	return (uint16(r.R[D]) << 8) | uint16(r.R[E])
}
func (r *Register) setDE(value uint16) {
	r.R[D], r.R[E] = byte(value>>8), byte(value)
}

func (r *Register) HL() uint16 {
	return (uint16(r.R[H]) << 8) | uint16(r.R[L])
}
func (r *Register) setHL(value uint16) {
	r.R[H], r.R[L] = byte(value>>8), byte(value)
}

func (cpu *CPU) getRegister(s string) uint16 {
	switch s {
	case "A":
		return uint16(cpu.Reg.R[A])
	case "F":
		return uint16(cpu.Reg.R[F])
	case "B":
		return uint16(cpu.Reg.R[B])
	case "C":
		return uint16(cpu.Reg.R[C])
	case "D":
		return uint16(cpu.Reg.R[D])
	case "E":
		return uint16(cpu.Reg.R[E])
	case "H":
		return uint16(cpu.Reg.R[H])
	case "L":
		return uint16(cpu.Reg.R[L])
	case "AF":
		return cpu.Reg.AF()
	case "BC":
		return cpu.Reg.BC()
	case "DE":
		return cpu.Reg.DE()
	case "HL":
		return cpu.Reg.HL()
	case "SP":
		return cpu.Reg.SP
	}

	return 0
}
