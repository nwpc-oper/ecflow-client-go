/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 4.0.1
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: ecflow_util.i

package ecflow_client

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
typedef _gostring_ swig_type_6;
typedef _gostring_ swig_type_7;
typedef _gostring_ swig_type_8;
typedef long long swig_type_9;
typedef long long swig_type_10;
typedef long long swig_type_11;
typedef long long swig_type_12;
extern void _wrap_Swig_free_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_ecflow_client_149cbf1b1d6d7ca4(swig_intgo arg1);
extern void _wrap_NodeStatusRecord_path__set_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1, swig_type_1 arg2);
extern swig_type_2 _wrap_NodeStatusRecord_path__get_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern void _wrap_NodeStatusRecord_status__set_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1, swig_type_3 arg2);
extern swig_type_4 _wrap_NodeStatusRecord_status__get_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern uintptr_t _wrap_new_NodeStatusRecord_ecflow_client_149cbf1b1d6d7ca4(void);
extern void _wrap_delete_NodeStatusRecord_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern uintptr_t _wrap_new_EcflowClientWrapper_ecflow_client_149cbf1b1d6d7ca4(swig_type_5 arg1, swig_type_6 arg2);
extern void _wrap_delete_EcflowClientWrapper_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern void _wrap_EcflowClientWrapper_setConnectTimeout_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1, swig_intgo arg2);
extern swig_intgo _wrap_EcflowClientWrapper_sync_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern uintptr_t _wrap_EcflowClientWrapper_statusRecords_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern swig_type_7 _wrap_EcflowClientWrapper_statusRecordsJson_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern swig_type_8 _wrap_EcflowClientWrapper_errorMessage_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern uintptr_t _wrap_new_NodeStatusRecordVector__SWIG_0_ecflow_client_149cbf1b1d6d7ca4(void);
extern uintptr_t _wrap_new_NodeStatusRecordVector__SWIG_1_ecflow_client_149cbf1b1d6d7ca4(swig_type_9 arg1);
extern uintptr_t _wrap_new_NodeStatusRecordVector__SWIG_2_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern swig_type_10 _wrap_NodeStatusRecordVector_size_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern swig_type_11 _wrap_NodeStatusRecordVector_capacity_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern void _wrap_NodeStatusRecordVector_reserve_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1, swig_type_12 arg2);
extern _Bool _wrap_NodeStatusRecordVector_isEmpty_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern void _wrap_NodeStatusRecordVector_clear_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
extern void _wrap_NodeStatusRecordVector_add_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_NodeStatusRecordVector_get_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_NodeStatusRecordVector_set_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1, swig_intgo arg2, uintptr_t arg3);
extern void _wrap_delete_NodeStatusRecordVector_ecflow_client_149cbf1b1d6d7ca4(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"

type _ unsafe.Pointer

var Swig_escape_always_false bool
var Swig_escape_val interface{}

type _swig_fnptr *byte
type _swig_memberptr *byte

type _ sync.Mutex

type swig_gostring struct {
	p uintptr
	n int
}

func swigCopyString(s string) string {
	p := *(*swig_gostring)(unsafe.Pointer(&s))
	r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
	Swig_free(p.p)
	return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_ecflow_client_149cbf1b1d6d7ca4(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrNodeStatusRecord uintptr

func (p SwigcptrNodeStatusRecord) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNodeStatusRecord) SwigIsNodeStatusRecord() {
}

func (arg1 SwigcptrNodeStatusRecord) SetPath_(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_NodeStatusRecord_path__set_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0), *(*C.swig_type_1)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrNodeStatusRecord) GetPath_() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_NodeStatusRecord_path__get_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
	swig_r_1 = swigCopyString(swig_r)
	return swig_r_1
}

func (arg1 SwigcptrNodeStatusRecord) SetStatus_(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_NodeStatusRecord_status__set_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrNodeStatusRecord) GetStatus_() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_NodeStatusRecord_status__get_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
	swig_r_1 = swigCopyString(swig_r)
	return swig_r_1
}

func NewNodeStatusRecord() (_swig_ret NodeStatusRecord) {
	var swig_r NodeStatusRecord
	swig_r = (NodeStatusRecord)(SwigcptrNodeStatusRecord(C._wrap_new_NodeStatusRecord_ecflow_client_149cbf1b1d6d7ca4()))
	return swig_r
}

func DeleteNodeStatusRecord(arg1 NodeStatusRecord) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_NodeStatusRecord_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0))
}

type NodeStatusRecord interface {
	Swigcptr() uintptr
	SwigIsNodeStatusRecord()
	SetPath_(arg2 string)
	GetPath_() (_swig_ret string)
	SetStatus_(arg2 string)
	GetStatus_() (_swig_ret string)
}

type SwigcptrEcflowClientWrapper uintptr

func (p SwigcptrEcflowClientWrapper) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrEcflowClientWrapper) SwigIsEcflowClientWrapper() {
}

func NewEcflowClientWrapper(arg1 string, arg2 string) (_swig_ret EcflowClientWrapper) {
	var swig_r EcflowClientWrapper
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (EcflowClientWrapper)(SwigcptrEcflowClientWrapper(C._wrap_new_EcflowClientWrapper_ecflow_client_149cbf1b1d6d7ca4(*(*C.swig_type_5)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_6)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func DeleteEcflowClientWrapper(arg1 EcflowClientWrapper) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_EcflowClientWrapper_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrEcflowClientWrapper) SetConnectTimeout(arg2 int) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_EcflowClientWrapper_setConnectTimeout_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))
}

func (arg1 SwigcptrEcflowClientWrapper) Sync() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_EcflowClientWrapper_sync_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrEcflowClientWrapper) StatusRecords() (_swig_ret NodeStatusRecordVector) {
	var swig_r NodeStatusRecordVector
	_swig_i_0 := arg1
	swig_r = (NodeStatusRecordVector)(SwigcptrNodeStatusRecordVector(C._wrap_EcflowClientWrapper_statusRecords_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func (arg1 SwigcptrEcflowClientWrapper) StatusRecordsJson() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_EcflowClientWrapper_statusRecordsJson_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
	swig_r_1 = swigCopyString(swig_r)
	return swig_r_1
}

func (arg1 SwigcptrEcflowClientWrapper) ErrorMessage() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_EcflowClientWrapper_errorMessage_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
	swig_r_1 = swigCopyString(swig_r)
	return swig_r_1
}

type EcflowClientWrapper interface {
	Swigcptr() uintptr
	SwigIsEcflowClientWrapper()
	SetConnectTimeout(arg2 int)
	Sync() (_swig_ret int)
	StatusRecords() (_swig_ret NodeStatusRecordVector)
	StatusRecordsJson() (_swig_ret string)
	ErrorMessage() (_swig_ret string)
}

type SwigcptrNodeStatusRecordVector uintptr

func (p SwigcptrNodeStatusRecordVector) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrNodeStatusRecordVector) SwigIsNodeStatusRecordVector() {
}

func NewNodeStatusRecordVector__SWIG_0() (_swig_ret NodeStatusRecordVector) {
	var swig_r NodeStatusRecordVector
	swig_r = (NodeStatusRecordVector)(SwigcptrNodeStatusRecordVector(C._wrap_new_NodeStatusRecordVector__SWIG_0_ecflow_client_149cbf1b1d6d7ca4()))
	return swig_r
}

func NewNodeStatusRecordVector__SWIG_1(arg1 int64) (_swig_ret NodeStatusRecordVector) {
	var swig_r NodeStatusRecordVector
	_swig_i_0 := arg1
	swig_r = (NodeStatusRecordVector)(SwigcptrNodeStatusRecordVector(C._wrap_new_NodeStatusRecordVector__SWIG_1_ecflow_client_149cbf1b1d6d7ca4(C.swig_type_9(_swig_i_0))))
	return swig_r
}

func NewNodeStatusRecordVector__SWIG_2(arg1 NodeStatusRecordVector) (_swig_ret NodeStatusRecordVector) {
	var swig_r NodeStatusRecordVector
	_swig_i_0 := arg1.Swigcptr()
	swig_r = (NodeStatusRecordVector)(SwigcptrNodeStatusRecordVector(C._wrap_new_NodeStatusRecordVector__SWIG_2_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0))))
	return swig_r
}

func NewNodeStatusRecordVector(a ...interface{}) NodeStatusRecordVector {
	argc := len(a)
	if argc == 0 {
		return NewNodeStatusRecordVector__SWIG_0()
	}
	if argc == 1 {
		if _, ok := a[0].(int64); !ok {
			goto check_2
		}
		return NewNodeStatusRecordVector__SWIG_1(a[0].(int64))
	}
check_2:
	if argc == 1 {
		return NewNodeStatusRecordVector__SWIG_2(a[0].(NodeStatusRecordVector))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrNodeStatusRecordVector) Size() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_NodeStatusRecordVector_size_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrNodeStatusRecordVector) Capacity() (_swig_ret int64) {
	var swig_r int64
	_swig_i_0 := arg1
	swig_r = (int64)(C._wrap_NodeStatusRecordVector_capacity_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrNodeStatusRecordVector) Reserve(arg2 int64) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_NodeStatusRecordVector_reserve_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0), C.swig_type_12(_swig_i_1))
}

func (arg1 SwigcptrNodeStatusRecordVector) IsEmpty() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_NodeStatusRecordVector_isEmpty_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrNodeStatusRecordVector) Clear() {
	_swig_i_0 := arg1
	C._wrap_NodeStatusRecordVector_clear_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrNodeStatusRecordVector) Add(arg2 NodeStatusRecord) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	C._wrap_NodeStatusRecordVector_add_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))
}

func (arg1 SwigcptrNodeStatusRecordVector) Get(arg2 int) (_swig_ret NodeStatusRecord) {
	var swig_r NodeStatusRecord
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (NodeStatusRecord)(SwigcptrNodeStatusRecord(C._wrap_NodeStatusRecordVector_get_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrNodeStatusRecordVector) Set(arg2 int, arg3 NodeStatusRecord) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	C._wrap_NodeStatusRecordVector_set_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.uintptr_t(_swig_i_2))
}

func DeleteNodeStatusRecordVector(arg1 NodeStatusRecordVector) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_NodeStatusRecordVector_ecflow_client_149cbf1b1d6d7ca4(C.uintptr_t(_swig_i_0))
}

type NodeStatusRecordVector interface {
	Swigcptr() uintptr
	SwigIsNodeStatusRecordVector()
	Size() (_swig_ret int64)
	Capacity() (_swig_ret int64)
	Reserve(arg2 int64)
	IsEmpty() (_swig_ret bool)
	Clear()
	Add(arg2 NodeStatusRecord)
	Get(arg2 int) (_swig_ret NodeStatusRecord)
	Set(arg2 int, arg3 NodeStatusRecord)
}
