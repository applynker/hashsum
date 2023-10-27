package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

type hashFunc func([]byte) string

var method2Func map[string]hashFunc = map[string]hashFunc{
	"md5":    md5Hex,
	"sha1":   sha1Hex,
	"sha256": sha256Hex,
	"sha512": sha512hex,
}

func main() {
	if len(os.Args) != 3 {
		help()
		return
	}
	method := os.Args[1]
	content := []byte(os.Args[2])
	hashF, ok := method2Func[method]
	if !ok {
		help()
		return
	}
	fmt.Println(hashF(content))
}

func help() {
	info := "hashsum [md5|sha1|sha256|sha512] content-to-hash"
	fmt.Println(info)
}

func md5Hex(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func sha1Hex(b []byte) string {
	return fmt.Sprintf("%x", sha1.Sum(b))
}

func sha256Hex(b []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

func sha512hex(b []byte) string {
	return fmt.Sprintf("%x", sha512.Sum512(b))
}
