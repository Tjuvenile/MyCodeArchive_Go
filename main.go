package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello main")
	//base_tool.AutoWiki("202401232400", "2602", "20240124025442", "tag-CeaStor-CeaStor3.2.1-2201-20240102111033")
	//s := ""
	//s2 := strings.Split(s, ",")
	//fmt.Println(reflect.TypeOf(s2[0]))
	//fmt.Printf("%v", s2)
	//fmt.Println([]byte(strings.Join(s2, "")))
	//fmt.Println(len(s2))
	//fmt.Println(cap(s2))
	//var par = LinkParam{}
	//if err := json.Unmarshal(respStruct.Payload, &par); err != nil {
	//	t.Errorf("parse data failed, resp: %s", string(resp))
	//}
	//
	//jsonData, _ := json.MarshalIndent(par, "", "  ")
	//// 打印漂亮格式的 JSON 字符串
	//t.Logf("jsonData: %s", string(jsonData))
}

type Resource struct {
	data int
}

func test1() *Resource {
	return &Resource{data: 3}
}
