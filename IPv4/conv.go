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

func ipv4ToHex(ipv4s []string) ([]string, error) {
	hexValues := make([]string, len(ipv4s))

	for i, ipv4 := range ipv4s {
		ipv4CString, err := windows.UTF16PtrFromString(ipv4)
		if err != nil {
			return nil, err
		}

		var buffer [4]byte
		_, _, err = procRtlIpv4StringToAddressA.Call(uintptr(unsafe.Pointer(ipv4CString)), 1, uintptr(unsafe.Pointer(&buffer[0])), 0)
		if err != nil && err.(windows.Errno) == 0 {
			hexValues[i] = fmt.Sprintf("%02x%02x%02x%02x", buffer[0], buffer[1], buffer[2], buffer[3])
		} else {
			return nil, err
		}
	}

	return hexValues, nil
}

func main() {
	//ipv4s := []string{"192.168.1.1", "10.0.0.1"}
	hexValues, err := ipv4ToHex(inputE)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, hexValue := range hexValues {
		fmt.Printf("IPv4: %s, Hex: %s\n", inputE[i], hexValue)
	}
}
