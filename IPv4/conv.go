package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	ntdll                       = windows.NewLazySystemDLL("ntdll.dll")
	procRtlIpv4StringToAddressA = ntdll.NewProc("RtlIpv4StringToAddressA")

	//input = "99.105.97.111"
	inputE = []string{"99.105.97.111", "99.111.109.101", "101.115.116.105", "98.101.110.101"}
)

func RtlIpv4StringToAddressA(hFunction *windows.LazyProc, String *uint16, Strict uint8, Address *uint32, Terminator **uint16) error {
	r0, _, e1 := hFunction.Call(uintptr(unsafe.Pointer(String)), uintptr(Strict), uintptr(unsafe.Pointer(Address)), uintptr(unsafe.Pointer(Terminator)))
	if r0 != uintptr(windows.STATUS_SUCCESS) {
		return fmt.Errorf("RtlIpv4StringToAddressA failed: %s", e1)
	}
	return nil
}

func deob(input []string, elements int, output []uint32) []uint32 {
	for i := 0; i < elements; i++ {
		var address uint32
		var terminator *uint16

		for i := 0; i < elements; i++ {
			err := RtlIpv4StringToAddressA(procRtlIpv4StringToAddressA, windows.StringToUTF16Ptr(input[i]), 1, &address, &terminator)
			if err != nil {
				fmt.Println(err)
			}

			output[i] = address
			//fmt.Printf("%X\n", output[i])

		}

	}
	return output
}
func main() {

	var output = make([]uint32, len(inputE))
	o := deob(inputE, len(inputE), output)

	fmt.Printf("%x", o)

}
