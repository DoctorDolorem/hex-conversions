package main

import (
	"fmt"
)

var input = []int{0x63, 0x69, 0x61, 0x6F}

func convertToIPv4(rawByteOne int, rawByteTwo int, rawByteThree int, rawByteFour int) string {
	output := fmt.Sprintf("%d.%d.%d.%d", rawByteOne, rawByteTwo, rawByteThree, rawByteFour)
	return output
}

func main() {
	IP := convertToIPv4(input[0], input[1], input[2], input[3])
	fmt.Println("HEX to IP: ", IP)
}
