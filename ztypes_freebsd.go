// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs types_freebsd.go

package tunnel

const (
	tunSDEBUG  = 0x8004745a
	tunGDEBUG  = 0x40047459
	tunSIFINFO = 0x8008745b
	tunGIFINFO = 0x4008745c
	tunSLMODE  = 0x8004745d
	tunSIFMODE = 0x8004745e
	tunSIFPID  = 0x2000745f
	tunSIFHEAD = 0x80047460
	tunGIFHEAD = 0x40047461
)

const (
	tunMTU = 0x5dc
	tunMRU = 0x4000
)

type tunInfo struct {
	Baudrate int32
	Mtu      int16
	Type     uint8
	Dummy    uint8
}
