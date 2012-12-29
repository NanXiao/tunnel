// Copyright 2012 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package tunnel

/*
#include <sys/types.h>

#include <net/if_tun.h>
*/
import "C"

const (
	tun_OPEN    = C.TUN_OPEN
	tun_INITED  = C.TUN_INITED
	tun_RCOLL   = C.TUN_RCOLL
	tun_IASET   = C.TUN_IASET
	tun_DSTADDR = C.TUN_DSTADDR
	tun_RWAIT   = C.TUN_RWAIT
	tun_ASYNC   = C.TUN_ASYNC
	tun_NBIO    = C.TUN_NBIO
	tun_BRDADDR = C.TUN_BRDADDR
	tun_STAYUP  = C.TUN_STAYUP
	tun_LAYER2  = C.TUN_LAYER2
	tun_READY   = C.TUN_READY
)

const (
	tunMTU = C.TUNMTU
	tunMRU = C.TUNMRU
)

const (
	tunSDEBUG  = C.TUNSDEBUG
	tunGDEBUG  = C.TUNGDEBUG
	tunSIFINFO = C.TUNSIFINFO
	tunGIFINFO = C.TUNGIFINFO
)

type tunInfo C.struct_tuninfo
