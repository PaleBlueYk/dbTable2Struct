package ymodel

type Objs struct {
	PkgName    string   // 包名
	Imp        []string // 导入包
	Objs       []St     // 结构体
	TableNames []TBName
}

type St struct {
	ObjName    string  // 结构体名
	ObjExtFrom string  // 结构体继承自
	FieldList  []Filed // 字段
}

type Filed struct {
	FieldName string // 字段名
	FieldType string // 字段类型
	FieldTag  string // 字段tag
}

type TBName struct {
	ObjName   string
	TableName string
}
