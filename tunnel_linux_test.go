// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tunnel_test

import "os/exec"

func setup(name string) error {
	cmd := exec.Command("ifconfig", name, "inet", "169.254.0.1", "dstaddr", "169.254.0.254")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func teardown(name string) error {
	cmd := exec.Command("ifconfig", name, "delete")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
