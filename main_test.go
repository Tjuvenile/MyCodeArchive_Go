package main

import (
	"github.com/agiledragon/gomonkey"
	"testing"
)

func TestFetchDataFromDatabase(t *testing.T) {
	// 创建一个Monkey实例
	patches := gomonkey.NewPatches()

	// 打桩，替换queryDatabase函数的实现
	patches.ApplyFunc(queryDatabase, func() (string, error) {
		return "mocked data", nil
	})
	defer patches.Reset()

	patches.ApplyFunc(sum, func() string {
		return "sdfsdfsfsd"
	})

	// 调用被测试的函数
	result, err := fetchDataFromDatabase()

	// 验证结果
	expected := "Processed: mocked data"
	if result != expected || err != nil {
		t.Errorf("Expected %v, got %v with error: %v", expected, result, err)
	}
}
