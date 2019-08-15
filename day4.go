package main

import (
	"fmt"
	"crypto/md5"
	"strconv"
	"encoding/hex"
)

func main() {
	input := "yzbqklnj"
	count := 0
	flag := false
	for !flag {
		newinp := input + strconv.Itoa(count)

		if isAnswer(newinp) {
			flag = true
			break;
		}
		count++
	}

	fmt.Println(count)
}

func isAnswer(input string) bool {
	md := GetMD5Hash(input)
	fmt.Println(md)
	fmt.Println(input, " - ", string(md[:5]))

	if string(md[:5]) == "00000" {
		return true
	}

	return false
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
