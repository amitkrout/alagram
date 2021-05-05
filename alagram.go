package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	input := os.Args[1]

	split := strings.Split(input, "")

	count := 0

	for i := 0; i <= len(split)-1; i++ {
		fmt.Print(split[i] + " ")
	}

	for i := 0; i <= len(split)-1; i++ {
		fmt.Print(split[i] + split[count+1] + " ")
		count = count + 1
		if count == len(split)-1 {
			break
		}
	}

	for i := 0; i <= len(split)-1; i++ {
		count1 := 0
		fmt.Print(split[i] + split[count1+1] + split[count1+2] + " ")
		count1 = count1 + 1
		if count == len(split)-1 {
			break
		}
	}

	for i := 0; i <= len(split)-1; i++ {
		count1 := 0
		fmt.Print(split[i] + split[count1+1] + split[count1+2] + split[count1+2] + " ")
		count1 = count1 + 1
		if count == len(split)-1 {
			break
		}

	}

	fmt.Print(input + " ")

	for i := 0; i <= len(split)-1; i++ {
		count5 := 0
		fmt.Print(split[i+1] + split[count5+2] + split[count5+3] + " ")
		count5 = count5 + 1
		if count == len(split)-1 {
			break
		}

	}

	for i := 0; i <= len(split)-1; i++ {
		count5 := 0
		fmt.Print(split[i+1] + split[count5+2] + split[count5+3] + split[count5+4] + " ")
		count5 = count5 + 1
		if count == len(split)-1 {
			break
		}

	}

	for i := 0; i <= len(split)-1; i++ {
		count5 := 0
		fmt.Print(split[i+2] + split[count5+3] + split[count5+4] + " ")
		count5 = count5 + 1
		if count == len(split)-1 {
			break
		}
	}
	fmt.Print("\n")
}
