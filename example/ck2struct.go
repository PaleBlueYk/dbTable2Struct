package example

import (
	"dbTable2Struct/yconfig"
	"dbTable2Struct/ydb"
	"dbTable2Struct/ymodel"
	"dbTable2Struct/yutils"
	"fmt"
	"github.com/PaleBlueYk/logger"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type col struct {
	CodecExpression   string `json:"codec_expression"`
	Comment           string `json:"comment"`
	DefaultExpression string `json:"default_expression"`
	DefaultType       string `json:"default_type"`
	Name              string `json:"name"`
	TtlExpression     string `json:"ttl_expression"`
	Type              string `json:"type"`
}

func CKDB2Struct() {
	tmp := template.Must(template.ParseFiles("template/obj.go"))
	tables := getTables()
	var o ymodel.Objs
	o.PkgName = yconfig.Config.Package
	o.Imp = append(o.Imp, yconfig.Config.Clickhouse.Imp...)
	for _, table := range tables {
		var st ymodel.St
		st.ObjName = yconfig.Config.Clickhouse.ObjPre+yutils.UnderscoreToUpperCamelCase(table)
		// TODO 结构体集成类,项目需要data的tag
		st.ObjExtFrom = yconfig.Config.Clickhouse.Ext + fmt.Sprintf(" `data:\"db:clickhouse;key:%s\"`", table)
		colList := getCols(table)
		var filedList []ymodel.Filed
		for _, c := range colList {
			fieldType := yutils.Transform2CodeType(c.Type)
			if fieldType == "time.Time" {
				o.Imp = append(o.Imp, "time")
			}
			filedList = append(filedList, ymodel.Filed{
				FieldName: yutils.UnderscoreToUpperCamelCase(c.Name),
				FieldType: fieldType,
				FieldTag:  fmt.Sprintf("`json:\"%s\"`", c.Name),
			})
		}
		st.FieldList = filedList
		o.Objs = append(o.Objs, st)
	}
	o.Imp = yutils.ListRemoveDuplication(o.Imp)
	outFile := filepath.FromSlash(yconfig.Config.Package + string(filepath.Separator) + yconfig.Config.CodeOutFile)
	_, err := os.Stat(outFile)
	if err != nil {
		if !os.IsExist(err) {
			os.Mkdir(yconfig.Config.Package, os.ModePerm)
			//os.Create(outFile)
		}
	}
	f, err := os.Create(outFile)
	if err != nil {
		logger.Error(err)
	}
	if err := tmp.Execute(f, o); err != nil {
		logger.Error(err)
	}
	cmd := exec.Command("gofmt", "-w", "-s", outFile)
	err = cmd.Run()
	if err != nil {
		logger.Error(err)
	}
}

func getTables() []string {
	var tableNames []string
	if err := ydb.DB.Raw("show tables;").Find(&tableNames).Error; err != nil {
		logger.Error(err)
		return nil
	}
	return tableNames
}

func getCols(tableName string) []col {
	var ColList []col
	if err := ydb.DB.Raw(fmt.Sprintf("desc %s;", tableName)).Find(&ColList).Error; err != nil {
		logger.Error(err)
		return nil
	}
	return ColList
}
