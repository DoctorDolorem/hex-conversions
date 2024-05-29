package main

import (
	"fmt"
	"syscall"
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

func Ipv4Deobfuscation(Ipv4Array []string) ([]int, error) {
	var deobfuscated []int

	for _, ipv4 := range Ipv4Array {
		var terminator *uint16
		var address uint32

		ipv4Ptr, err := syscall.UTF16PtrFromString(ipv4)
		if err != nil {
			return nil, err
		}

		r0, _, e1 := procRtlIpv4StringToAddressA.Call(uintptr(unsafe.Pointer(ipv4Ptr)), 0, uintptr(unsafe.Pointer(&address)), uintptr(unsafe.Pointer(&terminator)))
		if r0 != 0 {
			return nil, fmt.Errorf("RtlIpv4StringToAddressA failed: %s", e1)
		}

		// Reverse the byte order of the address
		address = (address>>24)&0xff | // move byte 3 to byte 0
			(address<<8)&0xff0000 | // move byte 1 to byte 2
			(address>>8)&0xff00 | // move byte 2 to byte 1
			(address<<24)&0xff000000 // byte 0 to byte 3

		deobfuscated = append(deobfuscated, int(address))
	}

	return deobfuscated, nil
}

func main() {

	ipv4, err := Ipv4Deobfuscation(inputE)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	for _, val := range ipv4 {
		fmt.Printf("0x%X ", val)
	}
	fmt.Println()

}
