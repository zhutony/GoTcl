//example.go v0.2.0
package main

/*
#include <stdlib.h>
#include <tcl.h>
#include <tclDecls.h>
#cgo LDFLAGS: -ltcl8.5

int Hello_Cmd_cgo(ClientData cdata, Tcl_Interp *interp, int objc,
        Tcl_Obj *const objv[]);
int Square_Cmd_cgo(ClientData cdata, Tcl_Interp *interp, int objc,
        Tcl_Obj *const objv[]);
*/
import "C"

import (
	"reflect"
	"unsafe"
)

const (
	TCL_OK       = 0
	TCL_ERROR    = 1
	TCL_RETURN   = 2
	TCL_BREAK    = 3
	TCL_CONTINUE = 4
)

func (interp *C.struct_Tcl_Interp) createCommand(name string,
	f *C.Tcl_ObjCmdProc) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.Tcl_CreateObjCommand(interp, cName, f, nil, nil)
}

func (interp *C.struct_Tcl_Interp) wrongNumArgs(objc C.int,
	objv **C.Tcl_Obj, message string) {
	var cMessage *C.char
	if message == "" {
		cMessage = nil
	} else {
		cMessage = C.CString(message)
		defer C.free(unsafe.Pointer(cMessage))
	}
	C.Tcl_WrongNumArgs(interp, objc, objv, cMessage)
}

func slicify(objc C.int, objv **C.Tcl_Obj) (objs []*C.Tcl_Obj) {
	// http://stackoverflow.com/a/14828189
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&objs))
	sliceHeader.Cap = (int)(objc)
	sliceHeader.Len = (int)(objc)
	sliceHeader.Data = uintptr(unsafe.Pointer(objv))
	return
}

//export Hello_Cmd
func Hello_Cmd(cdata C.ClientData, interp *C.struct_Tcl_Interp,
	objc C.int, objv **C.Tcl_Obj) C.int {
	if objc != 1 {
		interp.wrongNumArgs(1, objv, "")
		return TCL_ERROR
	}
	result := C.CString("Hello, World!")
	defer C.free(unsafe.Pointer(result))
	C.Tcl_SetObjResult(interp, C.Tcl_NewStringObj(result, -1))
	return TCL_OK
}

//export Square_Cmd
func Square_Cmd(cdata C.ClientData, interp *C.struct_Tcl_Interp,
	objc C.int, objv **C.Tcl_Obj) C.int {
	if objc != 2 {
		interp.wrongNumArgs(1, objv, "value")
		return TCL_ERROR
	}

	objs := slicify(objc, objv)

	var i C.int
	if C.Tcl_GetIntFromObj(interp, objs[1], &i) != TCL_OK {
		return TCL_ERROR
	}
	C.Tcl_SetObjResult(interp, C.Tcl_NewIntObj((C.int)(i*i)))
	return TCL_OK
}

//export Tclgoexample_Init
func Tclgoexample_Init(interp *C.struct_Tcl_Interp) C.int {
	interp.createCommand("::hello", (*C.Tcl_ObjCmdProc)(C.Hello_Cmd_cgo))
	interp.createCommand("::square", (*C.Tcl_ObjCmdProc)(C.Square_Cmd_cgo))

	return TCL_OK
}

func main() {}
