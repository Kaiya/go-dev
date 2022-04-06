package main

import (
	"fmt"
	"time"
	"unsafe"
)

const (
	offsetARM64HasATOMICS = unsafe.Offsetof(ARM64.HasATOMICS)
)

const CacheLinePadSize = 64

// CacheLinePad is used to pad structs to avoid false sharing.
type CacheLinePad struct{ _ [CacheLinePadSize]byte }

// CacheLineSize is the CPU's assumed cache line size.
// There is currently no runtime detection of the real cache line size
// so we use the constant per GOARCH CacheLinePadSize as an approximation.
var CacheLineSize uintptr = CacheLinePadSize

// The booleans in ARM64 contain the correspondingly named cpu feature bit.
// The struct is padded to avoid false sharing.
var ARM64 struct {
	_            CacheLinePad
	HasAES       bool
	HasPMULL     bool
	HasSHA1      bool
	HasSHA2      bool
	HasCRC32     bool
	HasATOMICS   bool
	HasCPUID     bool
	IsNeoverseN1 bool
	IsZeus       bool
	_            CacheLinePad
}

//go:noescape
func Cas64(ptr *uint64, old, new uint64) bool
func main() {
	fmt.Println("new fmt.Println, should have message above this line")
	go testNewProc()
}

func testNewProc() {
	time.Sleep(3)
}
func osInit() {
	ARM64.HasATOMICS = sysctlEnabled([]byte("hw.optional.armv8_1_atomics\x00"))
}

//go:noescape
func getsysctlbyname(name []byte) (int32, int32)

func sysctlEnabled(name []byte) bool {
	ret, value := getsysctlbyname(name)
	if ret < 0 {
		return false
	}
	return value > 0
}
