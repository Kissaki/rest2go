# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=github.com/Kissaki/rest.go
GOFMT=gofmt

GOFILES=\
	resource-interfaces.go\
	server.go\
	client.go\

include $(GOROOT)/src/Make.pkg

format:
	${GOFMT} -w server.go
	${GOFMT} -w client.go
