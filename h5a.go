package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"

	"reflect"
	"runtime"
	"unsafe"
)

type Attribute struct {
	Location
}

func newAttribute(id C.hid_t) *Attribute {
	d := &Attribute{Location{Identifier{id}}}
	runtime.SetFinalizer(d, (*Attribute).finalizer)
	return d
}

func createAttribute(id C.hid_t, name string, dtype *Datatype, dspace *Dataspace, acpl *PropList) (*Attribute, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	hid := C.H5Acreate2(id, c_name, dtype.id, dspace.id, acpl.id, P_DEFAULT.id)
	if err := h5err(C.herr_t(int(hid))); err != nil {
		return nil, err
	}
	return newAttribute(hid), nil
}

func openAttribute(id C.hid_t, name string) (*Attribute, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Aopen(id, c_name, P_DEFAULT.id)
	if err := h5err(C.herr_t(int(hid))); err != nil {
		return nil, err
	}
	return newAttribute(hid), nil
}

func (s *Attribute) finalizer() {
	err := s.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing attr: %s", err))
	}
}

func (s *Attribute) Id() int {
	return int(s.id)
}

// Access the type of an attribute
func (s *Attribute) GetType() Location {
	ftype := C.H5Aget_type(s.id)
	return Location{Identifier{ftype}}
}

// Close releases and terminates access to an attribute.
func (s *Attribute) Close() error {
	if s.id > 0 {
		err := C.H5Aclose(s.id)
		s.id = 0
		return h5err(err)
	}
	return nil
}

// Space returns an identifier for a copy of the dataspace for a attribute.
func (s *Attribute) Space() *Dataspace {
	hid := C.H5Aget_space(s.id)
	if int(hid) > 0 {
		return newDataspace(hid)
	}
	return nil
}

// Read reads raw data from a attribute into a buffer.
func (s *Attribute) Read(data interface{}, dtype *Datatype) error {
	var addr uintptr
	v := reflect.ValueOf(data)

	switch v.Kind() {

	case reflect.Array:
		addr = v.UnsafeAddr()

	case reflect.String:
		str := (*reflect.StringHeader)(unsafe.Pointer(v.UnsafeAddr()))
		addr = str.Data

	case reflect.Ptr:
		addr = v.Pointer()

	default:
		addr = v.UnsafeAddr()
	}

	rc := C.H5Aread(s.id, dtype.id, unsafe.Pointer(addr))
	err := h5err(rc)
	return err
}

// Write writes raw data from a buffer to an attribute.
func (s *Attribute) Write(data interface{}, dtype *Datatype) error {
	var addr uintptr
	v := reflect.ValueOf(data)
	switch v.Kind() {

	case reflect.Array:
		addr = v.UnsafeAddr()

	case reflect.String:
		str := (*reflect.StringHeader)(unsafe.Pointer(v.UnsafeAddr()))
		addr = str.Data

	case reflect.Ptr:
		addr = v.Pointer()

	default:
		addr = v.Pointer()
	}

	rc := C.H5Awrite(s.id, dtype.id, unsafe.Pointer(addr))
	err := h5err(rc)
	return err
}
