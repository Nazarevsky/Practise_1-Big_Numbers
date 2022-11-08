package cryption

import (
	"fmt"
	"strconv"
	"strings"
)

var blocks [][][]byte
var stateKey [][]byte
var sbox [][]byte = [][]byte{
	{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76},
	{0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0},
	{0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15},
	{0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75},
	{0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84},
	{0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf},
	{0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8},
	{0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2},
	{0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73},
	{0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb},
	{0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79},
	{0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08},
	{0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a},
	{0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e},
	{0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf},
	{0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16}}
var mulMatr [][]int = [][]int{
	{2, 3, 1, 1},
	{1, 2, 3, 1},
	{1, 1, 2, 3},
	{3, 1, 1, 2}}

func divMesIntoBlocks(bitMes string, countBlocks int) {
	blocks = make([][][]byte, countBlocks)
	ind := 0
	for i := 0; i < countBlocks; i++ {
		blocks[i] = make([][]byte, 4)
		for y := 0; y < 4; y++ {
			blocks[i][y] = make([]byte, 4)
			for x := 0; x < 4; x++ {
				num, _ := strconv.ParseInt(bitMes[ind:ind+8], 2, 8)
				blocks[i][y][x] = byte(num)
				ind += 8
			}
		}
	}
}

func intitKey(bitMes string) {
	stateKey = make([][]byte, 4)
	ind := 0
	for y := 0; y < 4; y++ {
		stateKey[y] = make([]byte, 4)
		for x := 0; x < 4; x++ {
			num, _ := strconv.ParseInt(bitMes[ind:ind+8], 2, 8)
			stateKey[y][x] = byte(num)
			ind += 8
		}
	}
}

func complete(mes string, to int) string {
	return strings.Repeat("0", to-len(mes)) + mes
}

func pad(mes string, to int) string {
	return strings.Repeat("0", to-(len(mes)%to)) + mes
}

func mesToBits(mes string) string {
	str := ""
	for _, c := range mes {
		str += fmt.Sprintf("%b", c)
	}
	return str
}

func printBlock() {
	for i := 0; i < len(blocks); i++ {
		for y := 0; y < 4; y++ {
			line := ""
			for x := 0; x < 4; x++ {
				line += fmt.Sprintf("%d ", blocks[i][y][x])
			}
			println(line)
		}
		println()
	}
}

func addKey() {
	for i := 0; i < len(blocks); i++ {
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				blocks[i][y][x] ^= stateKey[y][x]
			}
		}
	}
}

func getDecByHex(val string) int {
	num, _ := strconv.ParseInt(val, 16, 8)
	return int(num)
}

func subBytes() {
	for i := 0; i < len(blocks); i++ {
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				hex := complete(fmt.Sprintf("%x", blocks[i][y][x]), 2)
				blocks[i][y][x] = sbox[getDecByHex(string(hex[0]))][getDecByHex(string(hex[1]))]
			}
		}
	}
}

func shiftArr(arr []byte, shift int) []byte {
	return append(arr[shift:], arr[:shift]...)
}

func shiftBlock() {
	for i := 0; i < len(blocks); i++ {
		for y := 1; y < 4; y++ {
			blocks[i][y] = shiftArr(blocks[i][y], y)
		}
	}
}

func mixColumns() {
	var m [][]int = [][]int{
		{0xd4, 0xe0, 0xb8, 0x1e},
		{0xbf, 0xb4, 0x41, 0x27},
		{0x5d, 0x52, 0x11, 0x98},
		{0x30, 0xae, 0xf1, 0xe5}}
	var res [][]int = [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}

	res[0][0] = m[0][0]*mulMatr[0][0] ^ m[1][0]*mulMatr[0][1] ^ m[2][0]*mulMatr[0][2] ^ m[3][0]*mulMatr[0][3]
	println(res[0][0])
	// for y := 0; y < 4; y++ {
	// 	for x := 0; x < 4; x++ {
	// 		line := ""
	// 		for k := 0; k < 4; k++ {
	// 			res[y][x] ^= int(m[k][x]) * int(mulMatr[y][k])
	// 			line += fmt.Sprintf("%x * %x + ", int(m[k][x]), int(mulMatr[y][k]))
	// 		}
	// 		println(line, fmt.Sprintf("%d", res[y][x]))
	// 	}
	// }

	// for y := 0; y < 4; y++ {
	// 	line := ""
	// 	for x := 0; x < 4; x++ {
	// 		line += fmt.Sprintf("%x ", res[y][x])
	// 	}
	// 	println(line)
	// }

}

func xor(a, b byte) string {
	if a != b {
		return "1"
	}
	return "0"
}

func multHex(hex byte, mul byte) byte {
	if mul == 1 {
		return hex
	}
	bitHex := complete(fmt.Sprintf("%b", hex), 8)
	bitMul := fmt.Sprintf("%b", mul)
	var addMatr [][]byte

	for i := 0; i < len(bitHex); i++ {
		if bitHex[7-i] == 49 {
			addMatr = append(addMatr, []byte{})
			for j := len(bitMul) - 1; j >= 0; j-- {
				if bitMul[j] == 49 {
					addMatr[len(addMatr)-1] = append(addMatr[len(addMatr)-1], byte(i+(1-j)))
				}
			}
		}
	}

	println(addMatr[len(addMatr)-1][1])
	var str []string = make([]string, addMatr[len(addMatr)-1][1])
	for i := 0; i < len(addMatr); i++ {
		println(str[i])
	}
	// for i := 0; i < len(addMatr); i++ {
	// 	for j := 0; j < len(addMatr[0]); j++ {

	// 	}
	// }

	return 0
}

func AES_crypt(mes string, key string) string {
	bitMes := pad(mesToBits(mes), 128)
	// divMesIntoBlocks(bitMes, len(bitMes)/128)

	// bitKey := pad(mesToBits(key), 128)
	// intitKey(bitKey)

	//addKey()
	//subBytes()
	//shiftBlock()
	//printBlock()
	//mixColumns()
	//printBlock()

	multHex(0x68, 3)
	//println(fmt.Sprintf("%b", 0x1554^0x11b))
	return bitMes
}
