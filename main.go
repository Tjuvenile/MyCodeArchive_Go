package main

import (
	_ "MyCodeArchive_Go/logging"
	"encoding/json"
	"fmt"
)

func main() {
	t1, err := s()
	t2 := t1.(name)
	fmt.Println(t1, "***", err)
	fmt.Println(t2.N2)
	fmt.Println(t2.TypeName)
	fmt.Println(t2.N4)

}

type name struct {
	N  string
	N2 int
	TypeName
}

type TypeName struct {
	N3 string
	N4 string
}
type name2 struct {
	N  string
	N2 string
}

func s() (interface{}, error) {
	t1 := name2{
		N:  "123",
		N2: "aaaa",
	}
	t2, _ := json.Marshal(t1)

	params := name{}
	err := json.Unmarshal(t2, &params)

	var s3 interface{}
	s3 = params
	return s3, err
}
