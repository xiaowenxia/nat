package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaowenxia/nat/log"
	"github.com/xiaowenxia/nat/models"
)

type IpHistoryMessage struct {
	CommonMessage
	Total int64 `json:"total"`
	List  []models.IpHistory
}

func IpHistoryGet(c *gin.Context) {
	filter := &models.IpHistoryListFilter{}
	err := c.BindJSON(filter)
	if err != nil || filter.Limit == 0 {
		c.JSON(http.StatusBadRequest, &CommonMessage{
			ErrorMessage: "missing params",
		})
		return
	}

	log.Infow("get ip history list", "filter", filter)

	list, total, err := models.IpHistoryGetByFilter(*filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, &CommonMessage{
			ErrorMessage: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &IpHistoryMessage{
		Total: total,
		List:  list,
	})
}

func IpHistoryAdd(c *gin.Context) {
	record := &models.IpHistory{}
	err := c.BindJSON(record)
	if err != nil {
		c.JSON(http.StatusBadRequest, &CommonMessage{
			ErrorMessage: "missing params",
		})
		return
	}

	record.Ip = c.Request.RemoteAddr
	log.Infow("add ip history", "filter", record)

	ret, err := models.IpHistoryAdd(record.Ip, record.Message)
	if err != nil {
		c.JSON(http.StatusBadRequest, &CommonMessage{
			ErrorMessage: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &ret)
}
