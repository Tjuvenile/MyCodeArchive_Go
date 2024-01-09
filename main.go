package main

import (
	"MyCodeArchive_Go/cmd"
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
	s := [3]int{1, 2}
	s2 := s[:1]
	s2 = append(s2, 444)
	fmt.Println(cap(s))
	fmt.Println(len(s))
}

// User 是一个示例 GORM 模型
type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
