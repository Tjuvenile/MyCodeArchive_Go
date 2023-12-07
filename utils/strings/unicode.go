package strings

import (
	"strconv"
	"strings"
)

// UnquoteString '&'在经过json.Unmarshal之后会变成'\u0026'，本函数的目的在于还原为Unmarshal之前的字符串
// 如果发现第一个字母是'u'，尝试着转换一下，如果转换成，就追加字符串。
// 如果尝试转换失败，说明虽然以u开头，但不是正常的unicode字符，不需要进行额外转换，直接追加。
// 由于split的分割特性，所以i为第0位时，不需要加'\'。
func UnquoteString(str string) string {
	var res strings.Builder
	for i, v := range strings.Split(str, `\`) {
		if len(v) != 0 && v[0] == 'u' {
			tryUnqStr := `\` + v
			if unStr, err := strconv.Unquote(`"` + tryUnqStr + `"`); err == nil {
				res.WriteString(unStr)
				continue
			}
		}
		if i != 0 {
			res.WriteString(`\`)
		}
		res.WriteString(v)
	}
	return res.String()
}
