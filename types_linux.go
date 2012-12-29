// Copyright 2012 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package tunnel

/*
#include <sys/ioctl.h>
#include <sys/socket.h>

#include <linux/if.h>
#include <linux/if_tun.h>

#define IFREQSIZ sizeof(struct ifreq)
*/
import "C"

const (
	tun_READQ_SIZE = C.TUN_READQ_SIZE
	tun_TUN_DEV    = C.TUN_TUN_DEV
	tun_TAP_DEV    = C.TUN_TAP_DEV
	tun_TYPE_MASK  = C.TUN_TYPE_MASK
	tun_FASYNC     = C.TUN_FASYNC
	tun_NOCHECKSUM = C.TUN_NOCHECKSUM
	tun_NO_PI      = C.TUN_NO_PI
	tun_ONE_QUEUE  = C.TUN_ONE_QUEUE
	tun_PERSIST    = C.TUN_PERSIST
	tun_VNET_HDR   = C.TUN_VNET_HDR
	tun_PKT_STRIP  = C.TUN_PKT_STRIP
)

const (
	tunSETNOCSUM    = C.TUNSETNOCSUM
	tunSETDEBUG     = C.TUNSETDEBUG
	tunSETIFF       = C.TUNSETIFF
	tunSETPERSIST   = C.TUNSETPERSIST
	tunSETOWNER     = C.TUNSETOWNER
	tunSETLINK      = C.TUNSETLINK
	tunSETGROUP     = C.TUNSETGROUP
	tunGETFEATURES  = C.TUNGETFEATURES
	tunSETOFFLOAD   = C.TUNSETOFFLOAD
	tunSETTXFILTER  = C.TUNSETTXFILTER
	tunGETIFF       = C.TUNGETIFF
	tunGETSNDBUF    = C.TUNGETSNDBUF
	tunSETSNDBUF    = C.TUNSETSNDBUF
	tunATTACHFILTER = C.TUNATTACHFILTER
	tunDETACHFILTER = C.TUNDETACHFILTER
	tunGETVNETHDRSZ = C.TUNGETVNETHDRSZ
	tunSETVNETHDRSZ = C.TUNSETVNETHDRSZ
)

const (
	iff_TUN       = C.IFF_TUN
	iff_TAP       = C.IFF_TAP
	iff_NO_PI     = C.IFF_NO_PI
	iff_ONE_QUEUE = C.IFF_ONE_QUEUE
	iff_VNET_HDR  = C.IFF_VNET_HDR
	iff_TUN_EXCL  = C.IFF_TUN_EXCL
)

type tunPI C.struct_tun_pi

type ifReq struct {
	Name  [C.IFNAMSIZ]byte
	Flags uint16
	Pad   [C.IFREQSIZ - C.IFNAMSIZ - 2]byte
}
