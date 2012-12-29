// Copyright 2012 Mikio Hara. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build freebsd linux openbsd

package tunnel_test

import (
	"github.com/mikioh/tunnel"
	"os"
	"testing"
)

func TestNullOverIPv4Tunneling(t *testing.T) {
	if os.Getuid() != 0 {
		t.Logf("skipping test; must be root")
		return
	}

	c, err := tunnel.New()
	if err != nil {
		t.Fatalf("tunnel.New failed: %v", err)
	}
	ifi, err := c.Interface()
	if err != nil {
		t.Fatalf("tunnel.Interface failed: %v", err)
	}
	if err := setup(ifi.Name); err != nil {
		t.Fatalf("platform dependent setup failed: %v", err)
	}
	defer func() {
		c.Close()
		teardown(ifi.Name)
	}()
	b := make([]byte, 2048)
	for i := 0; i < 3; i++ {
		n, err := c.Read(b)
		if err != nil {
			t.Errorf("tunnel.Read failed: %v", err)
			return
		}
		t.Logf("tunnel.Read: %v bytes read", n)
		//if _, err := c.Write(b[:n]); err != nil {
		//	t.Errorf("tunnel.Write failed: %v", err)
		//	return
		//}
		//t.Logf("tunnel.Write: %v bytes written", n)
	}
}
