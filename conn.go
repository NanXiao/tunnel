// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tunnel implements tunnel software network interface for
// Unix variants.
package tunnel

import (
	"errors"
	"net"
	"os"
	"syscall"
)

var errNotSupported = errors.New("not supported")

// A Conn represents a tunnel software network interface endpoint.
type Conn struct {
	file    *os.File // tunnel device file
	ifindex int      // associated network interface index
}

func (c *Conn) ok() bool { return c != nil && c.file != nil }

// File returns the underlying os.File.
func (c *Conn) File() (*os.File, error) {
	if !c.ok() {
		return nil, syscall.EINVAL
	}
	return c.file, nil
}

// Interface returns the associated network interface.
func (c *Conn) Interface() (*net.Interface, error) {
	if !c.ok() {
		return nil, syscall.EINVAL
	}
	ifi, err := net.InterfaceByIndex(c.ifindex)
	if err != nil {
		return nil, err
	}
	return ifi, nil
}

// Read reads data from the tunnel endpoint.
func (c *Conn) Read(b []byte) (int, error) {
	if !c.ok() {
		return 0, syscall.EINVAL
	}
	n, err := c.file.Read(b)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// Write writes data to the tunnel endpoint.
func (c *Conn) Write(b []byte) (int, error) {
	if !c.ok() {
		return 0, syscall.EINVAL
	}
	return c.file.Write(b)
}

// Close closes the tunnel endpoint.
func (c *Conn) Close() error {
	if !c.ok() {
		return syscall.EINVAL
	}
	return c.file.Close()
}

// New returns a new tunnel endpoint.
func New() (*Conn, error) {
	c := &Conn{}
	if err := open(c); err != nil {
		return nil, err
	}
	return c, nil
}
