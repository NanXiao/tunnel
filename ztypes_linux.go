// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs ./types_linux.go

package tunnel

const (
	tun_READQ_SIZE = 0x1f4

	tun_TUN_DEV   = 0x1
	tun_TAP_DEV   = 0x2
	tun_TYPE_MASK = 0xf

	tun_FASYNC     = 0x10
	tun_NOCHECKSUM = 0x20
	tun_NO_PI      = 0x40
	tun_ONE_QUEUE  = 0x80
	tun_PERSIST    = 0x100
	tun_VNET_HDR   = 0x200

	tunSETNOCSUM    = 0x400454c8
	tunSETDEBUG     = 0x400454c9
	tunSETIFF       = 0x400454ca
	tunSETPERSIST   = 0x400454cb
	tunSETOWNER     = 0x400454cc
	tunSETLINK      = 0x400454cd
	tunSETGROUP     = 0x400454ce
	tunGETFEATURES  = 0x800454cf
	tunSETOFFLOAD   = 0x400454d0
	tunSETTXFILTER  = 0x400454d1
	tunGETIFF       = 0x800454d2
	tunGETSNDBUF    = 0x800454d3
	tunSETSNDBUF    = 0x400454d4
	tunATTACHFILTER = 0x401054d5
	tunDETACHFILTER = 0x401054d6
	tunGETVNETHDRSZ = 0x800454d7
	tunSETVNETHDRSZ = 0x400454d8

	iff_TUN       = 0x1
	iff_TAP       = 0x2
	iff_NO_PI     = 0x1000
	iff_ONE_QUEUE = 0x2000
	iff_VNET_HDR  = 0x4000
	iff_TUN_EXCL  = 0x8000

	tun_PKT_STRIP = 0x1
)

type tunPI struct {
	Flags uint16
	Proto uint16
}

type ifReq struct {
	Name  [0x10]byte
	Flags uint16
	Pad   [0x28 - 0x10 - 2]byte
}
