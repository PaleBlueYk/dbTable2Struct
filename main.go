package main

import (
	"dbTable2Struct/example"
	"dbTable2Struct/yconfig"
	"dbTable2Struct/ydb"
	"github.com/PaleBlueYk/logger"
)

func main() {
	if err := yconfig.ReadConf(); err != nil {
		logger.Error(err)
		return
	}
	if err := ydb.ConnDB(); err != nil {
		logger.Error(err)
		return
	}
	example.CKDB2Struct()
	logger.Info("all-down!")
}

