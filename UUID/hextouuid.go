package main

import (
	"fmt"
)

var input = []int{0x63, 0x69, 0x61, 0x6f, 0x63, 0x6f, 0x6d, 0x65, 0x65, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x6f, 0x6f}

func convertToUUID(rawByteOne int, rawByteTwo int, rawByteThree int, rawByteFour int, rawByteFive int, rawByteSix int, rawByteSeven int, rawByteEight int, rawByteNine int, rawByteTen int, rawByteEleven int, rawByteTwelve int, rawByteThirteen int, rawByteFourteen int, rawByteFifteen int, rawByteSixteen int) string {
	output0 := fmt.Sprintf("%02X%02X%02X%02X", rawByteFour, rawByteThree, rawByteTwo, rawByteOne)
	output1 := fmt.Sprintf("%02X%02X%02X%02X", rawByteSix, rawByteFive, rawByteEight, rawByteSeven)
	output2 := fmt.Sprintf("%02X%02X%02X%02X", rawByteNine, rawByteTen, rawByteEleven, rawByteTwelve)
	output3 := fmt.Sprintf("%02X%02X%02X%02X", rawByteThirteen, rawByteFourteen, rawByteFifteen, rawByteSixteen)
	result := fmt.Sprintf("%s-%s-%s-%s", output0, output1, output2, output3)
	return result
}

func main() {
	UUID := convertToUUID(input[0], input[1], input[2], input[3], input[4], input[5], input[6], input[7], input[8], input[9], input[10], input[11], input[12], input[13], input[14], input[15])
	fmt.Print(UUID)

}
