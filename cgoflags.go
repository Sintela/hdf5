// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Use the prebuilt bundled libraries for HDF5 which have been compiled with
// multithreading enabled.
// If Linux or Darwin then use the bundled libraries with multi threading enabled.
// If using Alpine just use the default libraries.
package hdf5

// #cgo !alpine LDFLAGS: -L${SRCDIR}/lib
// #cgo alpine LDFLAGS: -L/usr/local/lib
// #cgo CFLAGS: -I${SRCDIR}/include -I/usr/local/include
// #cgo darwin,!arm64 CFLAGS: -I/usr/local/Cellar/gcc/13.2.0/include/c++/13/parallel
// #cgo darwin,!arm64 LDFLAGS: -lhdf5_darwin_x86 -lhdf5_hl_darwin_x86 -lsz_darwin_x86 -lz -lm -ldl
// #cgo darwin,arm64 CFLAGS: -I/opt/homebrew/Cellar/gcc/13.2.0/include/c++/13/parallel
// #cgo darwin,arm64 LDFLAGS: -lhdf5_darwin_arm64 -lhdf5_hl_darwin_arm64 -lsz_darwin_arm64 -lz -lm -ldl
// #cgo linux,!arm64,!alpine LDFLAGS: -L/lib/x86_64-linux-gnu  -lhdf5_hl_x86 -lhdf5_x86 -lsz_x86 -lz -lm -ldl
// #cgo linux,arm64,!alpine LDFLAGS: -L/usr/lib/aarch64-linux-gnu  -lhdf5_hl_arm64 -lhdf5_arm64 -lsz_arm64 -lz -lm -ldl
// #cgo alpine LDFLAGS: -lhdf5_hl -lhdf5 -lz -lz -lm -ldl
// #include "hdf5.h"
// #include "hdf5_hl.h"
import "C"
