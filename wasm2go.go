package wasm2go

import (
	base "github.com/goccy/spidermonkeywasm2go/base"
	"fmt"
	"unsafe"
	_ "github.com/goccy/spidermonkeywasm2go/p4"
	_ "embed"
)

func NewWithWASIReserve(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports, reserveBytes int) *base.Module {
	m := &base.Module{Wasi_snapshot_preview1: wasi_snapshot_preview1}
	__memcap := reserveBytes
	if __memcap < 8847360 {
		__memcap = 8847360
	}
	m.Memory = make([]byte, 8847360, __memcap)
	m.M = unsafe.Pointer(unsafe.SliceData(m.Memory))
	m.MaxMem = 4294967296
	m.T0 = make([]any, 3580)
	m.G0 = int32(8388608)
	InitElemSeg_0_0(m)
	InitElemSeg_1_0(m)
	InitElemSeg_2_0(m)
	InitElemSeg_3_0(m)
	InitElemSeg_3_1(m)
	InitElemSeg_4_0(m)
	InitElemSeg_4_1(m)
	InitElemSeg_4_2(m)
	InitElemSeg_4_3(m)
	InitElemSeg_4_4(m)
	InitElemSeg_4_5(m)
	InitElemSeg_4_6(m)
	InitElemSeg_4_7(m)
	InitElemSeg_4_8(m)
	InitElemSeg_4_9(m)
	InitElemSeg_4_10(m)
	InitElemSeg_4_11(m)
	initData_0(m)
	return m
}

// NewWithWASI constructs a *Module with a custom
// wasi_snapshot_preview1 implementation and a default initial
// linear-memory reservation. Use NewWithWASIReserve to pre-size
// the reservation (e.g. to cover an interpreter's whole boot and
// avoid reallocating/copying linear memory on the first grow).
func NewWithWASI(wasi_snapshot_preview1 base.Wasi_snapshot_preview1Imports) *base.Module {
	return NewWithWASIReserve(wasi_snapshot_preview1, 11059200)
}

// New constructs a *Module using DefaultWASI() for the
// wasi_snapshot_preview1 import. Use NewWithWASI to plug in a
// custom implementation (sandboxed FS, captured stdout, ...).
func New() *base.Module {
	return NewWithWASI(base.DefaultWASI())
}
func initData_0(m *base.Module) {
	copy(m.Memory[8388608:], wasm2goData_data_bin[0:413848])
}
func Initialize(m *base.Module) {
	Fn16(m)
}
func WasmAlloc(m *base.Module, l0 int32) int32 {
	return Fn17(m, l0)
}
func WasmFree(m *base.Module, l0 int32) {
	Fn18(m, l0)
}
func WasmifyGetTypeName(m *base.Module, l0 int32, l1 int32) int64 {
	return Fn36(m, l0, l1)
}
func WasmInit(m *base.Module) int32 {
	return Fn38(m)
}
func WasmShutdown(m *base.Module) {
	Fn39(m)
}
func Inv_0_0(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn19(m, l0, l1)
	return
}
func Inv_0_1(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn26(m, l0, l1)
	return
}
func Inv_0_2(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn31(m, l0, l1)
	return
}
func Inv_0_3(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn33(m, l0, l1)
	return
}
func Inv_0_4(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn34(m, l0, l1)
	return
}
func Inv_0_5(m *base.Module, l0, l1 int32) (packed int64, err error) {
	savedG0 := m.G0
	defer func() {
		r := recover()
		if r != nil {
			m.G0 = savedG0
			err = fmt.Errorf("wasm trap: %v", r)
		}
	}()
	packed = Fn35(m, l0, l1)
	return
}
func Memory(m *base.Module) []byte {
	return m.Memory
}

//go:embed data.bin
var wasm2goData_data_bin []byte
