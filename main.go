package main

import (
	_ "MyCodeArchive_Go/logging"
	"fmt"
)

func main() {
	s := new(bool)
	s2 := new(int64)

	fmt.Println(*s, *s2)
}
