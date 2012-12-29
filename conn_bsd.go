// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build freebsd openbsd

package tunnel

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func open(c *Conn) error {
	var err error
	for i := 0; i < 10; i++ { // assume a few user processes want to use tun/tap
		c.file, err = os.OpenFile(fmt.Sprintf("/dev/tun%d", i), os.O_RDWR, 0)
		if err == nil {
			break
		}
	}
	if err != nil {
		return err
	}
	ss := strings.Split(c.file.Name(), "/")
	ifi, err := net.InterfaceByName(ss[len(ss)-1])
	if err != nil {
		c.file.Close()
		return err
	}
	c.ifindex = ifi.Index
	return nil
}
