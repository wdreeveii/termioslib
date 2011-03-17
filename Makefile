# Copyright 2010 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=termioslib

CGOFILES=\
	termioslib.go

#ifeq ($(GOOS),darwin)
#CGO_LDFLAGS=/usr/lib/libsqlite3.0.dylib
#else
#CGO_LDFLAGS=-L/usr/local/lib/oeis -lsqlite3
#endif

include $(GOROOT)/src/Make.pkg
