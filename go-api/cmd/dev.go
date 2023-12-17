package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 64userID:  OWT/HYPCwnCw42gvZKVNvZalZ0Lz2UHR950RPd3YRJOtElUioLwSE7Fyxxqz5vbk53uYfNa7x3pS5noPXenKsA==
	// res     :  OWT_HYPCwnCw42gvZKVNvZalZ0Lz2UHR950RPd3YRJOtElUioLwSE7Fyxxqz5vbk53uYfNa7x3pS5noPXenKsA

	// b64Encoding := base64.StdEncoding.WithPadding(base64.NoPadding)
	b64Encoding := base64.URLEncoding.WithPadding(base64.NoPadding)

	decoded, err := b64Encoding.DecodeString("OWT_HYPCwnCw42gvZKVNvZalZ0Lz2UHR950RPd3YRJOtElUioLwSE7Fyxxqz5vbk53uYfNa7x3pS5noPXenKsA")
	if err != nil {
		panic(err)
	}
	fmt.Println("decoded: ", decoded)
}
