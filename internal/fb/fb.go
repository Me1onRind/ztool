package fb

import (
	"hash/crc32"
	"strconv"
)

func FbByMod(number string, mod int) int {
	n, _ := strconv.Atoi(number)
	return n % mod
}

func FbByCrc32(hashStr string, mod int) int {
	n := int(crc32.ChecksumIEEE([]byte(hashStr)))
	return n % mod
}
