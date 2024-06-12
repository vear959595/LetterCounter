package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func encode(s string) string {
	count := 1
	result := ""
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count += 1
		} else {
			result += string(s[i-1])
			if count > 1 {
				result += strconv.Itoa(count)
				count = 1
			}
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')

	fmt.Println(encode(s))
}
