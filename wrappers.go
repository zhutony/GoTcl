//wrappers.go
package main

/*
#include <tcl.h>
#include <tclDecls.h>

int Hello_Cmd_cgo(ClientData cdata, Tcl_Interp *interp, int objc,
        Tcl_Obj *const objv[]) {
        int Hello_Cmd(ClientData cdata, Tcl_Interp *interp, int objc,
                Tcl_Obj *const objv[]);
        return Hello_Cmd(cdata, interp, objc, objv);
}

int Square_Cmd_cgo(ClientData cdata, Tcl_Interp *interp, int objc,
        Tcl_Obj *const objv[]) {
        int Square_Cmd(ClientData cdata, Tcl_Interp *interp, int objc,
                Tcl_Obj *const objv[]);
        return Square_Cmd(cdata, interp, objc, objv);
}
*/
import "C"
