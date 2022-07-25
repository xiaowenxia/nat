package server

import (
	"fmt"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/xiaowenxia/nat/config"
	"github.com/xiaowenxia/nat/log"
	"github.com/xiaowenxia/nat/models"
)

// APIError API Error message
type CommonMessage struct {
	ErrorCode    int
	ErrorMessage string
}

func Run() {
	models.Init(config.GConfig.DBUser,
		config.GConfig.DBPassword,
		config.GConfig.DBURL,
		config.GConfig.DBName,
		config.GConfig.DBPort)

	r := gin.New()

	log.Info("redirect gin output")
	gin.DefaultWriter = log.AlogWriter
	gin.DefaultErrorWriter = log.AlogWriter
	r.Use(gin.LoggerWithWriter(log.AlogWriter))

	// any panic error return 500 to client
	r.Use(gin.Recovery())
	r.Static("/", path.Join(config.GConfig.Server.Statics, "statics/dist"))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/ip_history/get", IpHistoryGet)
		v1.POST("/ip_history/add", IpHistoryAdd)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", config.GConfig.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Errorw("run server failed\n", "addr", addr)
		os.Exit(2)
	}

	log.Errorw("Will not get here")
	os.Exit(1)
}
