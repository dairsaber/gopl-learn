package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	fmt.Println(strings.Join(os.Args[1:], " "))
	t2 := time.Now().UnixNano()

	fmt.Println(t2 - t)
}
