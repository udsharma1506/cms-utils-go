package util

import (
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// This prepends the data with the padding value until the total length is reached
func LeftPad(data, padding string, totalLength int) string {
	currLength := len(data)
	for i := 0; i < (totalLength-currLength)/len(padding); i++ {
		data = padding + data
	}
	return data
}

// This appends the data with the padding value until the total length is reached
func RightPad(data, padding string, totalLength int) string {
	currLength := len(data)
	for i := 0; i < (totalLength-currLength)/len(padding); i++ {
		data += padding
	}
	return data
}

// converts the data string into multiples of mod bits and prepends the start index of data
func StrToHex(data string, mod int) string {
	dataLen := len(data)
	x := dataLen % mod
	reqdBits := mod - x
	if x != 0 {
		fillLen := dataLen + reqdBits
		data = LeftPad(data, "0", fillLen)
		reqdBits += mod
	}
	reqdBitsStr := strconv.Itoa(reqdBits)
	dummy := LeftPad(reqdBitsStr, "0", mod)
	retData := dummy + data
	return retData
}

// create a v4 UUID string without any "-" (hyphen)
func CreateUUID() string {
	//id := uuid.NewV4().String()
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return strings.Replace(id.String(), "-", "", -1)
	//return strings.Replace(id, "-", "", -1)
}