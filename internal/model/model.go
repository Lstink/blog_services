package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lstink/blog/global"
	"github.com/lstink/blog/pkg/setting"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreateBy   string `json:"created_by"`
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
	DeletedOn  string `json:"deleted_on"`
	IsDel      string `json:"is_del"`
}

func NewDBEngine(databaseSeting *setting.DatabaseSettings) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", databaseSeting.Username,
		databaseSeting.Password,
		databaseSeting.Host,
		databaseSeting.DBName,
		databaseSeting.Charset,
		databaseSeting.ParseTime)
	db, err := gorm.Open(databaseSeting.DBType, s)
	if err != nil {
		return nil, err
	}
	// 如果运行模式为 debug 模式，则启动日志记录
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSeting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSeting.MaxOpenConns)

	return db, nil
}
