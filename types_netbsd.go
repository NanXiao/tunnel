// Copyright 2012 Mikio Hara. All rights reserved.
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
)

const (
	tunSDEBUG  = C.TUNSDEBUG
	tunGDEBUG  = C.TUNGDEBUG
	tunSIFMODE = C.TUNSIFMODE
	tunSLMODE  = C.TUNSLMODE
	tunSIFHEAD = C.TUNSIFHEAD
	tunGIFHEAD = C.TUNGIFHEAD
)
