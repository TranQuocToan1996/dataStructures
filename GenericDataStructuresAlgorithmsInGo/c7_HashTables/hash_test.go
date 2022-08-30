package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestHashMD5AndSHA256(t *testing.T) {
	name1 := "Richard"
	name2 := "Richards"
	md5hash := md5.Sum([]byte(name1))
	sha256hash := sha256.Sum256([]byte(name1))
	fmt.Println(" MD5: ", md5hash)
	fmt.Println("SHA256: ", sha256hash)
	md5hash = md5.Sum([]byte(name2))
	sha256hash = sha256.Sum256([]byte(name2))
	fmt.Println(" MD5: ", md5hash)
	fmt.Println("SHA256: ", sha256hash)
}
