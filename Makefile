# Copyright 2009 The nntp-go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=nzb
GOFILES=\
	nzb.go

include $(GOROOT)/src/Make.pkg
