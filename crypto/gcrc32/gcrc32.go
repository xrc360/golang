// Package gcrc32 provides useful API for CRC32 encryption algorithms.
package gcrc32

import (
	"hash/crc32"

	"github.com/xrc360/golang/util/gconv"
)

// Encrypt encrypts any type of variable using CRC32 algorithms.
// It uses gconv package to convert `v` to its bytes type.
func Encrypt(v interface{}) uint32 {
	return crc32.ChecksumIEEE(gconv.Bytes(v))
}
