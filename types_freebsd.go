// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package tunnel

/*
#include <sys/ioctl.h>
#include <sys/socket.h>

#include <net/if.h>
#include <net/if_tun.h>
*/
import "C"

const (
	tunMTU = C.TUNMTU
	tunMRU = C.TUNMRU
)

const (
	tunSDEBUG  = C.TUNSDEBUG
	tunGDEBUG  = C.TUNGDEBUG
	tunSIFINFO = C.TUNSIFINFO
	tunGIFINFO = C.TUNGIFINFO
	tunSLMODE  = C.TUNSLMODE
	tunSIFMODE = C.TUNSIFMODE
	tunSIFPID  = C.TUNSIFPID
	tunSIFHEAD = C.TUNSIFHEAD
	tunGIFHEAD = C.TUNGIFHEAD
)

type tunInfo C.struct_tuninfo
