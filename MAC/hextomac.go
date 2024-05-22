package main

import (
	"fmt"
)

var input = []int{0x63, 0x69, 0x61, 0x6f, 0x63, 0x6f}

func convertToMAC(rawByteOne int, rawByteTwo int, rawByteThree int, rawByteFour int, rawByteFive int, rawByteSix int) string {
	output := fmt.Sprintf("%X-%X-%X-%X-%X-%X", rawByteOne, rawByteTwo, rawByteThree, rawByteFour, rawByteFive, rawByteSix)
	return output
}

func main() {
	MAC := convertToMAC(input[0], input[1], input[2], input[3], input[4], input[5])
	fmt.Println("HEX to MAC: ", MAC)
}
