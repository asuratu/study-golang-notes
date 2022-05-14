package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) (result string) {
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return
}

func printFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(13),
		convertToBin(123),
		convertToBin(5),
		convertToBin(0),
	)

	printFile("/Users/asura/Code/go/taobao/01、basic/abc.txt")
}
