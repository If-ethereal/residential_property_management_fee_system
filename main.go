package main

import (
	"graduation_project/api"
)

func main() {
	router := api.SetupRouters()
	router.Run() // 默认监听 0.0.0.0:8080
}

//
//import (
//	"go.uber.org/zap"
//	"graduation_project/config"
//	"graduation_project/db_conn"
//)
//
//func main() {
//	// 初始化配置
//	cfg := config.Init("config.yaml")
//
//	// 初始化日志
//	logger := core.InitLogger(cfg.Env)
//	defer logger.Sync()
//	zap.ReplaceGlobals(logger)
//
//	// 初始化MySQL
//	mysqlDB, err := db.InitMySQL(cfg.MySQL)
//	if err != nil {
//		zap.L().Fatal("Failed to connect to MySQL", zap.Error(err))
//	}
//	defer func() {
//		sqlDB, _ := mysqlDB.DB()
//		sqlDB.Close()
//	}()
//
//	// 初始化Redis
//	redisClient, err := db.InitRedis(cfg.Redis)
//	if err != nil {
//		zap.L().Fatal("Failed to connect to Redis", zap.Error(err))
//	}
//	defer redisClient.Close()
//
//	// 创建HTTP服务器
//	srv := server.NewServer(cfg, mysqlDB, redisClient)
//
//	// 启动服务
//	zap.L().Info("Starting server",
//		zap.String("address", cfg.Server.Address),
//		zap.Int("port", cfg.Server.Port))
//
//	if err := srv.Run(); err != nil {
//		zap.L().Fatal("Server run failed", zap.Error(err))
//	}
//}
