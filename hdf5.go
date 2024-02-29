// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hdf5 provides access to the HDF5 C library.
package hdf5 // import "gonum.org/v1/hdf5"

// #include "hdf5.h"
import "C"

import (
	"fmt"
	"unsafe"
)

// init initializes the hdf5 library
func init() {
	err := h5err(C.H5open())
	if err != nil {
		err_str := fmt.Sprintf("pb calling H5open(): %s", err)
		panic(err_str)
	}
}

// hdferror wraps hdf5 int-based error codes
type h5error struct {
	code int
}

func (h h5error) Error() string {
	return fmt.Sprintf("code %d", h.code)
}

func h5err(herr C.herr_t) error {
	if herr < 0 {
		return h5error{code: int(herr)}
	}
	return nil
}

func checkID(hid C.hid_t) error {
	if hid < 0 {
		return h5error{code: int(hid)}
	}
	return nil
}

// Close flushes all data to disk, closes all open identifiers, and cleans up memory.
// It should generally be called before your application exits.
func Close() error {
	return h5err(C.H5close())
}

// Version represents the currently used hdf5 library version
type Version struct {
	Major   uint
	Minor   uint
	Release uint
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Release)
}

// LibVersion returns version information for the HDF5 library.
func LibVersion() (Version, error) {
	var maj, min, rel C.uint
	var v Version
	err := h5err(C.H5get_libversion(&maj, &min, &rel))
	if err == nil {
		v.Major = uint(maj)
		v.Minor = uint(min)
		v.Release = uint(rel)
	}
	return v, err
}

// GarbageCollect collects garbage on all free-lists of all types.
func GarbageCollect() error {
	return h5err(C.H5garbage_collect())
}

// Object represents an hdf5 object.
type Object interface {
	Name() string
	Id() int
	File() *File
}

func PrintVersionInfo() {
	var majNum, minNum, relNum int
	err := C.H5get_libversion((*C.uint)(unsafe.Pointer(&majNum)), (*C.uint)(unsafe.Pointer(&minNum)), (*C.uint)(unsafe.Pointer(&relNum)))

	if err == 0 {
		fmt.Printf("HDF5 Library Version = %d.%d  Release %d\n", majNum, minNum, relNum)
	} else {
		fmt.Println("HDF5 Library Version : Error reading ")
	}
}

// SetExtent grows the extent of a dataset - because the hdf5 library doesn't include it
func SetExtent(d hdf5.Identifier, dims []uint) (err error) {
	cDims := (*C.hsize_t)(unsafe.Pointer(&dims[0]))
	rc := C.H5Dset_extent(C.hid_t(d.ID()), cDims)
	e := C.herr_t(rc)
	if e < 0 {
		err = fmt.Errorf("HDF5 error in SetExtent %d", e)
	}
	return err
}
