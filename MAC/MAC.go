package main

import (
	"fmt"
	"log"
	"unsafe"
)

var input = []int{0x63, 0x69, 0x61, 0x6f, 0x63, 0x6f, 0x6d, 0x65, 0x65, 0x73, 0x74, 0x69}
var sizeInput int = len(input)

func convertToMAC(rawByteOne int, rawByteTwo int, rawByteThree int, rawByteFour int, rawByteFive int, rawByteSix int) string {
	output := fmt.Sprintf("%X-%X-%X-%X-%X-%X", rawByteOne, rawByteTwo, rawByteThree, rawByteFour, rawByteFive, rawByteSix)
	return output
}

func generateMAC(pointer unsafe.Pointer, size int) (bool, string) {
	if pointer == nil || size == 0 || size%6 != 0 {
		log.Fatal("Invalid input, not a multiple of 6, requires padding")
		return false, ""
	}

	c := 6
	counter := 0
	output := ""

	for i := 0; i < size; i++ {
		if c == 6 {
			counter++

			MAC := convertToMAC(input[i], input[i+1], input[i+2], input[i+3], input[i+4], input[i+5])

			if i == size-4 {
				output = output + fmt.Sprintf("\\%s\\", MAC)
				break
			} else {
				output = output + fmt.Sprintf("\\%s\\, ", MAC)
			}

			c = 1

		} else {
			c++
		}
	}

	return true, output
}
func main() {

	_, output := generateMAC(unsafe.Pointer(&input[0]), sizeInput)
	fmt.Printf("%s};", output)
}
