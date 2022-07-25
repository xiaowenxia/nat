package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xiaowenxia/nat/log"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var Engine *xorm.Engine

func Init(user, password, host, dbName string, port int64) (err error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, password, host, port, dbName)
	log.Infow("Got DB Config", "url", url)

	Engine, err = xorm.NewEngine("mysql", url)
	if err != nil {
		log.Errorw("new db engine error", "err", err)
		return
	}
	Engine.SetMapper(names.GonicMapper{})
	Engine.ShowSQL(true)
	return nil
}
