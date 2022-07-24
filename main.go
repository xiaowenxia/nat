package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	Id        int64     `xorm:"id" json:"id"`
	Ip        string    `xorm:"ip" json:"ip"`
	Message   string    `xorm:"message" json:"message"`
	CreatedAt time.Time `xorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at" json:"updated_at"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", "xxw:3edcMJU&@rm-bp1c7tg12zu2f5j95ko.mysql.rds.aliyuncs.com/?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}
}
