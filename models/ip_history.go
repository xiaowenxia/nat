package models

import (
	"time"

	"github.com/xiaowenxia/nat/log"
)

type IpHistory struct {
	Id        int64     `xorm:"id" json:"id"`
	Ip        string    `xorm:"ip" json:"ip"`
	Message   string    `xorm:"message" json:"message"`
	CreatedAt time.Time `xorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at" json:"updated_at"`
}

type IpHistoryListFilter struct {
	ID     int64  `json:"id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Ip     string `json:"ip"`
}

func IpHistoryGetByFilter(filter IpHistoryListFilter) ([]IpHistory, int64, error) {
	var list []IpHistory

	session := Engine.Table("ip_history")

	if filter.ID != 0 {
		session = session.Where("id = ?", filter.ID)
	}

	if filter.Ip != "" {
		session = session.And("ip = ?", filter.Ip)
	}

	count, err := session.Limit(filter.Limit, filter.Offset).Desc("id").FindAndCount(&list)
	if err != nil {
		log.Errorw("get ip history list failed", "filter", filter)
		return nil, 0, err
	}

	return list, count, nil
}

func IpHistoryAdd(ip, message string) (record IpHistory, err error) {
	record.Ip = ip
	record.Message = message
	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Now()
	_, err = Engine.Table("ip_history").InsertOne(&record)
	if err != nil {
		log.Error("create ip history failed")
		return
	}
	return
}
