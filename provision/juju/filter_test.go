// Copyright 2012 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package juju

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestfilterOutputWithPythonWarnings(c *C) {
	output := []byte(`2012-11-28 16:00:35,615 WARNING Ubuntu Cloud Image lookups encrypted but not authenticated
2012-11-28 16:00:35,616 INFO Connecting to environment...
/usr/local/lib/python2.7/dist-packages/txAWS-0.2.3-py2.7.egg/txaws/client/base.py:208: UserWarning: The client attribute on BaseQuery is deprecated and will go away in future release.
warnings.warn('The client attribute on BaseQuery is deprecated and'
2012-11-28 16:00:36,787 INFO Connected to environment.
2012-11-28 16:00:37,110 INFO Connecting to machine 23 at 10.19.2.195
pre-restart:
  - python manage.py dbmigrate
  - python manage.py collectstatic --noinput)
	`)
	expected := []byte(`pre-restart:
  - python manage.py dbmigrate
  - python manage.py collectstatic --noinput)
	`)
	got := filterOutput(output)
	c.Assert(string(got), Equals, string(expected))
}

func (s *S) TestfilterOutputWithJujuLog(c *C) {
	output := []byte(`/usr/lib/python2.6/site-packages/juju/providers/ec2/files.py:8: DeprecationWarning: the sha module is deprecated; use the hashlib module instead
  import sha
2012-06-05 17:26:15,881 WARNING ssl-hostname-verification is disabled for this environment
2012-06-05 17:26:15,881 WARNING EC2 API calls not using secure transport
2012-06-05 17:26:15,881 WARNING S3 API calls not using secure transport
2012-06-05 17:26:15,881 WARNING Ubuntu Cloud Image lookups encrypted but not authenticated
2012-06-05 17:26:15,891 INFO Connecting to environment...
2012-06-05 17:26:16,657 INFO Connected to environment.
2012-06-05 17:26:16,860 INFO Connecting to machine 0 at 10.170.0.191
; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	expected := []byte(`; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	got := filterOutput(output)
	c.Assert(string(got), Equals, string(expected))
}

func (s *S) TestfilterOutputWithoutJujuLog(c *C) {
	output := []byte(`/usr/lib/python2.6/site-packages/juju/providers/ec2/files.py:8: DeprecationWarning: the sha module is deprecated; use the hashlib module instead
  import sha
; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	expected := []byte(`; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	got := filterOutput(output)
	c.Assert(string(got), Equals, string(expected))
}

func (s *S) TestFiterOutputWithSshWarning(c *C) {
	output := []byte(`2012-06-20 16:54:09,922 WARNING ssl-hostname-verification is disabled for this environment
2012-06-20 16:54:09,922 WARNING EC2 API calls not using secure transport
2012-06-20 16:54:09,922 WARNING S3 API calls not using secure transport
2012-06-20 16:54:09,922 WARNING Ubuntu Cloud Image lookups encrypted but not authenticated
2012-06-20 16:54:09,924 INFO Connecting to environment...
2012-06-20 16:54:10,549 INFO Connected to environment.
2012-06-20 16:54:10,664 INFO Connecting to machine 3 at 10.170.0.166
Warning: Permanently added '10.170.0.121' (ECDSA) to the list of known hosts.
total 0`)
	expected := []byte("total 0")
	got := filterOutput(output)
	c.Assert(string(got), Equals, string(expected))
}

func (s *S) TestfilterOutputWithoutJujuLogAndWarnings(c *C) {
	output := []byte(`; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	expected := []byte(`; generated by /sbin/dhclient-script
search novalocal
nameserver 192.168.1.1`)
	got := filterOutput(output)
	c.Assert(string(got), Equals, string(expected))
}

func (s *S) TestfilterOutputRSA(c *C) {
	output := []byte(`/usr/lib/python2.6/site-packages/juju/providers/ec2/files.py:8: DeprecationWarning: the sha module is deprecated; use the hashlib module instead
  import sha
2012-08-22 14:39:18,211 WARNING ssl-hostname-verification is disabled for this environment
2012-08-22 14:39:18,211 WARNING EC2 API calls not using secure transport
2012-08-22 14:39:18,212 WARNING S3 API calls not using secure transport
2012-08-22 14:39:18,212 WARNING Ubuntu Cloud Image lookups encrypted but not authenticated
2012-08-22 14:39:18,222 INFO Connecting to environment...
2012-08-22 14:39:18,854 INFO Connected to environment.
2012-08-22 14:39:18,989 INFO Connecting to machine 4 at 10.170.1.193
Warning: Permanently added '10.170.1.193' (RSA) to the list of known hosts.
Last login: Wed Aug 15 16:08:40 2012 from 10.170.1.239`)
	expected := []byte("Last login: Wed Aug 15 16:08:40 2012 from 10.170.1.239")
	got := filterOutput(output)
	c.Assert(string(got), Equals, string(expected))
}
