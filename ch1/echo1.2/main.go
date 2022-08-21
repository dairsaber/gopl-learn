package main

import (
	"fmt"
	"os"
)

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	for i, arg := range os.Args {
// 		s := strconv.Itoa(i) + ": " + arg
// 		fmt.Println(s)
// 	}
// }

func main() {
	for i, arg := range os.Args {
		s := fmt.Sprintf("%d: %s", i, arg)
		fmt.Println(s)
	}
}
