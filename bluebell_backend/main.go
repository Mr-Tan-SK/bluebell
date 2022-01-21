package main

import (
	"bluebell_backend/controller"
	"bluebell_backend/dao/mysql"
	"bluebell_backend/dao/redis"
	"bluebell_backend/logger"
	"bluebell_backend/pkg/snowflake"
	"bluebell_backend/routers"
	"bluebell_backend/settings"
	"fmt"
)

func main() {
	// 1. 加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}

	// 2. 加载初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// 3. 加载初始化mysql
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接

	// 4. 加载初始化redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	// 5. 加载加密id分配算法
	if err := snowflake.Init("2021-01-19", 1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 6. 注册路由并运行
	err := controller.InitTrans("zh")
	if err != nil {
		fmt.Printf("InitTrans failed, err: %v\n", err.Error())
		return
	}
	r := routers.Setup()
	err = r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
