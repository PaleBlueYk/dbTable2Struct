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

	if strings.Contains(s, "datetime64") {
		s = "time.Time"
	}
	if strings.Contains(s, "fixedstring") {
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