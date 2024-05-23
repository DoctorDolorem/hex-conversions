package main

import (
	"fmt"
	"log"
	"unsafe"
)

var input = []int{0x63, 0x69, 0x61, 0x6f, 0x63, 0x6f, 0x6d, 0x65, 0x65, 0x73, 0x74, 0x69, 0x62, 0x65, 0x6e, 0x65}
var sizeInput int = len(input)

func convertToIPv4(rawByteOne int, rawByteTwo int, rawByteThree int, rawByteFour int) string {
	output := fmt.Sprintf("%d.%d.%d.%d", rawByteOne, rawByteTwo, rawByteThree, rawByteFour)
	return output
}

func generateIPv4(pointer unsafe.Pointer, size int) (bool, string) {
	if pointer == nil || size == 0 || size%4 != 0 {
		log.Fatal("Invalid input, not a multiple of 4, requires padding")
		return false, ""
	}

	c := 4
	counter := 0
	output := ""

	for i := 0; i < size; i++ {
		if c == 4 {
			counter++

			IP := convertToIPv4(input[i], input[i+1], input[i+2], input[i+3])

			if i == size-4 {
				output = output + fmt.Sprintf("\"%s\"", IP)
				break
			} else {
				output = output + fmt.Sprintf("\"%s\", ", IP)
			}

			c = 1

		} else {
			c++
		}
	}

	return true, output
}
func main() {

	_, output := generateIPv4(unsafe.Pointer(&input[0]), sizeInput)
	fmt.Printf("%s};", output)
}
