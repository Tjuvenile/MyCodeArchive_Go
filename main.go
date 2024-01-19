package main

import (
	"MyCodeArchive_Go/utils/strings_"
)

func init() {
	//cmd.Execute()
}

func main() {
	strings_.EncryptExample("admin")
	strings_.EncryptExample("123456")
}

// User 是一个示例 GORM 模型
type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
