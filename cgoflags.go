// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #cgo LDFLAGS: -L/usr/local/lib -L${SRCDIR}/lib
// #cgo CFLAGS: -I${SRCDIR}/include -I/usr/local/include
// #cgo darwin,!arm64 LDFLAGS: -lhdf5_darwin_x86 -lhdf5_hl_darwin_x86 -lsz_darwin_x86 -lz -lm -ldl
// #cgo darwin,!arm64 CFLAGS: -I/usr/local/Cellar/gcc/13.2.0/include/c++/13/paralle
// #cgo darwin,arm64 LDFLAGS: -lhdf5_darwin_arm64 -lhdf5_hl_darwin_arm64 -lsz_darwin_arm64 -lz -lm -ldl
// #cgo darwin,arm64 CFLAGS: -I/opt/homebrew/Cellar/gcc/13.2.0/include/c++/13/parallel
// #cgo linux,!arm64 LDFLAGS: -L/lib/x86_64-linux-gnu  -lhdf5_hl_x86 -lhdf5_x86 -lsz_x86 -lz -lm -ldl
// #cgo linux,arm64 LDFLAGS: -L/usr/lib/aarch64-linux-gnu  -lhdf5_hl_arm64 -lhdf5_arm64 -lsz_arm64 -lz -lm -ldl
// #include "hdf5.h"
// #include "hdf5_hl.h"
import "C"
