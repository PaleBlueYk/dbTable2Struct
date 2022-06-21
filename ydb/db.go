package ydb

import (
	"dbTable2Struct/yconfig"
	"fmt"
	"github.com/PaleBlueYk/logger"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDB() error {
	var err error
	dsn := fmt.Sprintf("tcp://%s:%d?database=%s&username=%s&password=%s&read_timeout=10&write_timeout=20", yconfig.Config.Clickhouse.Host, yconfig.Config.Clickhouse.Port, yconfig.Config.Clickhouse.DB, yconfig.Config.Clickhouse.User, yconfig.Config.Clickhouse.Pwd)
	DB, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
