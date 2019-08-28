package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := "yzbqklnj"
	count := 0

	resflag := make(chan bool)
	resval := make(chan int)
	defer close(resflag)
	defer close(resval)

	for {
		select {
		case <-resval:
			exitWithResult(<-resval)
		default:
			count++
			go isAnswer(resflag, resval, input, count)
		}
	}

}

func isAnswer(resflag chan bool, resval chan int, input string, num int) {
	newinp := input + strconv.Itoa(num)
	md := GetMD5Hash(newinp)
	fmt.Println(md)
	fmt.Println(input, " - ", string(md[:6]))

	// Part 1 Condition
	// if string(md[:5]) == "00000" {
	// 	return true
	// }

	// Part 2 Condition
	if string(md[:6]) == "000000" {
		resval <- num
		resflag <- true
	}
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func exitWithResult(result int) {
	fmt.Println("Result is : ", result)
	os.Exit(1)
}
