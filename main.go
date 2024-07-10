package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"project/controller"
	"project/ioc"
	"project/repository"
	prcache "project/repository/cache"
	prdao "project/repository/dao"
	"project/server"
)

func main() {
	initViper()
	StartServer()
}

func initViper() {
	file := pflag.String("config", "config/config.yaml", "配置文件路径")
	pflag.Parse()

	viper.SetConfigType("yaml")
	viper.SetConfigFile(*file)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func StartServer() {
	log := ioc.InitLog()
	db := ioc.InitDB(log)
	cmd := ioc.InitRedis()
	dao := prdao.NewGORMDAO(db)
	cache := prcache.NewRedisCache(cmd)
	repo := repository.NewCachedUserRepository(cache, dao, log)
	ser := server.NewServer(repo)
	con := controller.NewContrpller(ser, log)
	p := controller.NewPage()
	service := controller.NewService(con, p)

	if err := service.InitServer(); err != nil {
		panic(err)
		return
	}
}
