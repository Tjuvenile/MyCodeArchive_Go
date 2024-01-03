package main

import (
	"MyCodeArchive_Go/cmd"
)

func init() {
	cmd.Execute()
}

func main() {
	//err := dcs.CreateRelationExe()
	//if err != nil {
	//	logging.Log.Errorf(fmt.Sprintf("%+v", err))
	//}

}

// User 是一个示例 GORM 模型
type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
