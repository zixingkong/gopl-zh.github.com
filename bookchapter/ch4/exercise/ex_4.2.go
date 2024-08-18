package main

//练习 4.2： 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

func main() {
	hashType := flag.String("type", "sha256", "指定哈希算法: sha256, sha384, sha512")
	flag.Parse()

	var hasher hash.Hash
	switch *hashType {
	case "sha256":
		hasher = sha256.New()
	case "sha384":
		hasher = sha512.New384()
	case "sha512":
		hasher = sha512.New()
	default:
		fmt.Println("不支持的哈希算法")
		os.Exit(1)
	}

	if _, err := io.Copy(hasher, os.Stdin); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hash := hex.EncodeToString(hasher.Sum(nil))
	fmt.Println(hash)
}
