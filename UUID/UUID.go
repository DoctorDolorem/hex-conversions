package main

import (
	"fmt"
	"log"
	"unsafe"
)

var input = []int{0x63, 0x69, 0x61, 0x6f, 0x63, 0x6f, 0x6d, 0x65, 0x65, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x6f, 0x6f, 0x63, 0x69, 0x61, 0x6f, 0x63, 0x6f, 0x6d, 0x65, 0x65, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x6f, 0x6f}
var sizeInput = len(input)

func convertToUUID(rawByteOne int, rawByteTwo int, rawByteThree int, rawByteFour int, rawByteFive int, rawByteSix int, rawByteSeven int, rawByteEight int, rawByteNine int, rawByteTen int, rawByteEleven int, rawByteTwelve int, rawByteThirteen int, rawByteFourteen int, rawByteFifteen int, rawByteSixteen int) string {

	output0 := fmt.Sprintf("%02X%02X%02X%02X", rawByteFour, rawByteThree, rawByteTwo, rawByteOne)
	output1 := fmt.Sprintf("%02X%02X%02X%02X", rawByteSix, rawByteFive, rawByteEight, rawByteSeven)
	output2 := fmt.Sprintf("%02X%02X%02X%02X", rawByteNine, rawByteTen, rawByteEleven, rawByteTwelve)
	output3 := fmt.Sprintf("%02X%02X%02X%02X", rawByteThirteen, rawByteFourteen, rawByteFifteen, rawByteSixteen)
	result := fmt.Sprintf("%s-%s-%s-%s", output0, output1, output2, output3)
	return result
}

func generateUUID(pointer unsafe.Pointer, size int) (bool, string) {
	if pointer == nil || size == 0 || size%16 != 0 {
		log.Fatal("Invalid input, not a multiple of 16, requires padding")
		return false, ""
	}

	c := 16
	counter := 0
	output := ""

	for i := 0; i < size; i++ {
		if c == 16 {
			counter++

			UUID := convertToUUID(input[i], input[i+1], input[i+2], input[i+3], input[i+4], input[i+5], input[i+6], input[i+7], input[i+8], input[i+9], input[i+10], input[i+11], input[i+12], input[i+13], input[i+14], input[i+15])

			if i == size-16 {
				output = output + fmt.Sprintf("\"%s\"", UUID)
				break
			} else {
				output = output + fmt.Sprintf("\"%s\", ", UUID)
			}
			c = 1

		} else {
			c++
		}

	}

	return true, output

}

func main() {

	_, output := generateUUID(unsafe.Pointer(&input[0]), sizeInput)
	fmt.Printf("%s};", output)

}
