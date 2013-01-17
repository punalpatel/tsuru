// Copyright 2013 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package heal

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestRegisterAndGetHealer(c *C) {
	var h Healer
	Register("my-healer", h)
	got, err := Get("my-healer")
	c.Assert(err, IsNil)
	c.Assert(got, DeepEquals, h)
}
