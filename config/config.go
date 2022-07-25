package config

var GConfig Config

type Config struct {
	DBName     string
	DBUser     string
	DBURL      string
	DBPassword string
	DBPort     int64
	Server     struct {
		Port    int64
		Statics string
	}
}
