// Copyright 2012 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin plan9 windows

package tunnel

func open(c *Conn) error {
	return errNotSupported
}
