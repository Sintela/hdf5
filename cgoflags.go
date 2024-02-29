// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #cgo LDFLAGS: -L/home/sintela/hdf5-1.14.3/hdf5/lib -lhdf5 -lhdf5_hl -L/home/sintela/szip-2.1.1/szip/lib -lsz -lz
// #cgo CFLAGS: -I/home/sintela/hdf5-1.14.3/hdf5/include -I/home/sintela/szip-2.1.1/szip/include
// #include "hdf5.h"
import "C"
