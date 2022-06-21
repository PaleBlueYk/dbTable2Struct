package yutils

import "strings"

func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// Transform2CodeType ck类型转换成golang代码类型
func Transform2CodeType(s string) string {
	s = strings.ToLower(s)

	if strings.Contains(s, "datetime64") || strings.Contains(s, "datetime") || strings.Contains(s, "date") {
		s = "time.Time"
	}
	if strings.Contains(s, "fixedstring") || strings.Contains(s, "object") {
		s = "string"
	}
	return s
}

// ListRemoveDuplication 数组去重
func ListRemoveDuplication(s []string) []string {
	set := make(map[string]interface{})
	j := 0
	for _, v := range s {
		if _, exist := set[v]; exist {
			continue
		}
		set[v] = struct {}{}
		s[j] = v
		j++
	}
	return s[:j]
}

// SnakeString 大驼峰转下划线
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}