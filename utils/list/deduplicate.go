package list

import (
	"MyCodeArchive_Go/utils/fault"
	"MyCodeArchive_Go/utils/request_model/db"
	"fmt"
)

// CheckDuplicateByList 插入一个手机号，就比对一下计数器和map之间是否匹配，如果不匹配，意味着有重复，即可报错。
func CheckDuplicateByList(name string) *fault.Fault {
	var example db.ExampleDb
	exList, err := example.QueryByName(name)
	if err != nil {
		return err
	}

	var phoneMap map[string]bool
	countMap := 0
	for _, ex := range exList {
		phoneMap[ex.Phone], countMap = true, countMap+1
		if len(phoneMap) != countMap {
			return fault.Wrap(fmt.Sprintf("The phone %q duplicates with the existing %q phone", ex.Phone, ex.Name),
				fmt.Sprintf("手机号 %q 与已有的 %q 手机号重复", ex.Phone, ex.Name))
		}
	}
	return nil
}
