package main

import (
	"MyCodeArchive_Go/cmd"
	"encoding/json"
	"fmt"
)

func init() {
	cmd.Execute()
}

func main() {
	//err := dcs.CreateRelationExe()
	//if err != nil {
	//	logging.Log.Errorf(fmt.Sprintf("%+v", err))
	//}
	type sdf struct {
		S int
	}
	s2 := sdf{}
	err := json.Unmarshal([]byte(""), &s2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ending")
}
