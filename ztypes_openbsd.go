// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs ./types_openbsd.go

package tunnel

const (
	tun_OPEN    = 0x1
	tun_INITED  = 0x2
	tun_RCOLL   = 0x4
	tun_IASET   = 0x8
	tun_DSTADDR = 0x10
	tun_RWAIT   = 0x40
	tun_ASYNC   = 0x80
	tun_NBIO    = 0x100
	tun_BRDADDR = 0x200
	tun_STAYUP  = 0x400
	tun_LAYER2  = 0x800
	tun_READY   = 0x3

	tunMTU = 0xbb8
	tunMRU = 0x4000
)

const (
	tunSDEBUG  = 0x8004745e
	tunGDEBUG  = 0x4004745f
	tunSIFINFO = 0x800c745b
	tunGIFINFO = 0x400c745c
)

type tunInfo struct {
	Mtu      uint32
	Type     uint16
	Flags    uint16
	Baudrate uint32
}
