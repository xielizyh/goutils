package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// FileMd5 returns md5 of a file
func FileMd5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("md5.go hash.FileMd5 os open error: %v", err)
	}
	h := md5.New()
	_, err = io.Copy(h, file)
	if err != nil {
		return "", fmt.Errorf("md5.go hash.FileMd5 io copy error: %v", err)
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// StringMd5 returns md5 of a string
func StringMd5(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	return hex.EncodeToString(md5.Sum(nil))
}
