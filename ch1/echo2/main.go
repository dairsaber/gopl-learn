package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now().UnixNano()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	t2 := time.Now().UnixNano()
	fmt.Println(t2 - t)

}
