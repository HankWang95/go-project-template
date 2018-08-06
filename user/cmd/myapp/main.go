package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/smartwalle/dbr"
	"github.com/smartwalle/dbs"
	"github.com/HankWang95/go-project-template/user/service"
	"github.com/HankWang95/go-project-template/user/service/repository/mysql"
	"github.com/HankWang95/go-project-template/user/service/repository/redis"
	"github.com/HankWang95/go-project-template/user/transport/restful"
)

func main() {
	var db, _ = dbs.NewSQL("mysql", "root:yangfeng@tcp(192.168.1.99:3306)/v3?parseTime=true", 30, 5)
	var rPool = dbr.NewRedis("192.168.1.99:6379", "", 10, 10, 5)

	var uRepo = redis.NewUserRepository(rPool, mysql.NewUserRepository(db))
	var uServ = service.NewUserService(uRepo)
	var uHandler = restful.NewUserHandler(uServ)

	var s = gin.Default()
	uHandler.Run(s)
	s.Run(":8888")
}
