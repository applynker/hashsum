package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"os"
)

const version = "2.0.0"

type hashFunc func([]byte) string

var method2Func map[string]hashFunc = map[string]hashFunc{
	"md5":    md5Hex,
	"sha1":   sha1Hex,
	"sha256": sha256Hex,
	"sha512": sha512hex,
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}
	method := os.Args[1]
	if method == "-v" || method == "--version" {
		fmt.Println(version)
		return
	}
	hashF, ok := method2Func[method]
	if !ok {
		help()
		return
	}
	var content []byte
	var err error
	if len(os.Args) == 3 {
		content = []byte(os.Args[2])
	} else {
		content, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("failed to read stdin. error:", err)
			return
		}
	}
	fmt.Println(hashF(content))
}

func help() {
	info := `hashsum [md5|sha1|sha256|sha512] content-to-hash
  Hash content. If there is no content provided, it will read input from stdin.
  Example:
  - hashsum md5 "hello world"
  - echo "hello world" | hashsum md5
  - cat hello.txt | hashsum md5

hashsum -v|--version
  Show hashsum current version
`
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
