package main

import (
	"MyCodeArchive_Go/request/dcs"
	"MyCodeArchive_Go/utils/logging"
	"encoding/json"
	"fmt"
)

func main() {
	//formatter := txtfmt.NewTableFormatter("HAGroupName", "GatewayNum", "VipNum", "RecordNum")
	//var table []txtfmt.TableRow
	//row := txtfmt.TableRow{
	//	"HAGroupName": "123",
	//	"GatewayNum":  "456",
	//	"VipNum":      "789",
	//	"RecordNum":   "101112",
	//}
	//table = append(table, row)
	//fmt.Println(formatter.Format(table))
	err := dcs.CreateStrategyExe()
	if err != nil {
		logging.Log.Errorf(fmt.Sprintf("%+v", err))
	}
}

// 假设有一个用于查询数据库的函数
func queryDatabase() (string, error) {
	fmt.Println("查询表1")
	fmt.Println("查询表2")
	// 实际情况中会有数据库查询逻辑
	return "real data", nil
}

func sum() string {
	return "haha1"
}

// 要测试的函数，依赖于queryDatabase
func fetchDataFromDatabase() (string, error) {
	fmt.Println("获取数据")
	fmt.Println("进行操作")
	s := sum()
	fmt.Println(s)
	data, err := queryDatabase()
	if err != nil {
		return "", err
	}
	fmt.Println("返回数据", data)
	// 处理data的逻辑
	return "Processed: " + data, nil
}

type test3 struct {
	bytesList []byte
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

func s(test3 *test3) (interface{}, error) {
	params := name{}
	return params, json.Unmarshal(test3.bytesList, &params)
}
