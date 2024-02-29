// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #cgo LDFLAGS: -L/usr/local/lib -L${SRCDIR}/lib
// #cgo CFLAGS: -I${SRCDIR}/include -I/usr/local/include
// #cgo darwin LDFLAGS: -lhdf5_darwin -lhdf5_hl_darwin -lsz_darwin
// #cgo linux,!arm64 LDFLAGS: -lhdf5_x86 -lhdf5_hl_x86 -lsz_x86
// #cgo linux,arm64 LDFLAGS: -L/usr/lib/aarch64-linux-gnu  -lhdf5_hl_arm64 -lhdf5_arm64 -lsz_arm64 -lz -lm -lpthread -ldl
// #include "hdf5.h"
// #include "hdf5_hl.h"
import "C"
