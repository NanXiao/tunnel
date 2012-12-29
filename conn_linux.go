// Copyright 2012 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tunnel

import (
	"net"
	"os"
	"syscall"
	"unsafe"
)

func open(c *Conn) error {
	var err error
	c.file, err = os.OpenFile("/dev/net/tun", os.O_RDWR, 0)
	if err != nil {
		return err
	}
	ifr := &ifReq{}
	ifr.Flags = iff_TUN | iff_NO_PI | iff_TUN_EXCL
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, c.file.Fd(), tunSETIFF, uintptr(unsafe.Pointer(ifr)))
	if e != 0 {
		return syscall.Errno(e)
	}
	i := 0
	for ifr.Name[i] != 0 {
		i++
	}
	s := string(ifr.Name[:i])
	ifi, err := net.InterfaceByName(s)
	if err != nil {
		c.file.Close()
		return err
	}
	c.ifindex = ifi.Index
	return nil
}
