package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/huynhtrongtien/dove/clients"
	"github.com/huynhtrongtien/dove/controllers/v3/category"
	"github.com/huynhtrongtien/dove/controllers/v3/product"
	"github.com/huynhtrongtien/dove/global"
	"github.com/huynhtrongtien/dove/middlewares"
	"github.com/huynhtrongtien/dove/pkg/log"
	"github.com/huynhtrongtien/dove/pkg/tracing"
	"github.com/huynhtrongtien/dove/pkg/utilities"
	"github.com/huynhtrongtien/dove/services"
	"github.com/spf13/viper"
)

func main() {
	initViper()
	initLogger()
	initGlobalSetting()
	initDatabase()
	initRedis()
	initServices()
	middlewares.InitMiddlewares()

	// start HTTP Server
	StartHTTPServer()
}

func StartHTTPServer() {

	router := gin.Default()
	router.Use(gin.Recovery())

	// setup tracing
	jaegerConfig := &tracing.JaegerHTTPConfig{
		Environment: global.Environment(),
		ServiceName: global.ServiceName(),
		Endpoint:    viper.GetString("trace.end_point"),
		URLPath:     viper.GetString("trace.url_path"),
	}
	tracing.StartOpenTelemetryV2(jaegerConfig)
	tracing.SetupMiddleware(router, jaegerConfig.ServiceName)
	log.Bg().Info("[start-http-server] connect to jaeger udp success")

	categoryHandler := category.NewHandler()
	productHandler := product.NewHandler()
	apiV3 := router.Group("/api/v3")
	{
		categoryClient := apiV3.Group("/categories")
		{
			categoryClient.GET("/:category_uuid", categoryHandler.Read)
		}

		productClient := apiV3.Group("/categories/:category_uuid/products")
		{
			productClient.GET("", productHandler.List)
		}
	}

	host := fmt.Sprintf("localhost:%d", viper.GetInt("service.port"))

	log.Bg().Info("[start service] start service", log.Field("host", host))

	router.Run(host)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func initViper() {

	cfgFile := "./config.toml"
	fmt.Printf("Read config file from env: [%s] \n", cfgFile)

	folder, fileName, ext, err := utilities.ExtractFilePath(cfgFile)
	if err != nil {
		fmt.Printf("Extract config file failed %s err: %s \n", viper.ConfigFileUsed(), err.Error())
		os.Exit(-1)
	}
	fmt.Printf("Extract config file success folder[%s] fileName[%s] ext[%s] \n", folder, fileName, ext)

	// Setting
	viper.AddConfigPath(folder)
	viper.SetConfigName(fileName)
	viper.AutomaticEnv()
	viper.SetConfigType(ext)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("FATAL: Viper using config file failed %s err: %s \n", viper.ConfigFileUsed(), err.Error())
		os.Exit(-1)
	}

	fmt.Printf("Service using config file: %s \n", viper.ConfigFileUsed())
	//watch on config change
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s", e.Name)
	})

	fmt.Println("Start initialize config success.")
}

func initLogger() {
	viper.SetDefault("log.filepath", "./log.log")
	log.InitLogger(&log.Configuration{
		JSONFormat:      true,
		LogLevel:        log.DebugLevel,
		StacktraceLevel: log.FatalLevel,
		Console:         &log.ConsoleConfiguration{},
		File: &log.FileConfiguration{
			Filename:   viper.GetString("log.filepath"),
			MaxSize:    10,
			MaxAge:     14,
			MaxBackups: 10,
		},
	})
}

func initGlobalSetting() {
	if err := global.InitSetting(); err != nil {
		log.Bg().Fatal("[init-global-setting] read global config failed", log.Err(err))
		os.Exit(-1)
	}
}

func initDatabase() {
	var err error
	cfg := &clients.MySQLConfig{
		Address:  viper.GetString("database.address"),
		DBName:   viper.GetString("database.dbname"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
	}
	clients.MySQLClient, err = clients.NewMySQLClient(cfg)
	if err != nil {
		log.Bg().Fatal("[init-database] create database connection failed", log.Field("file", viper.ConfigFileUsed()), log.Err(err))
		os.Exit(-1)
		return
	}

	err = clients.AutoMigrate()
	if err != nil {
		log.Bg().Fatal("[init-database] auto migrate failed", log.Field("file", viper.ConfigFileUsed()),
			log.Field("address", cfg.Address), log.Field("dbname", cfg.DBName), log.Err(err))
		os.Exit(-1)
	}

	log.Bg().Info("[init-database] create database connection success", log.Field("address", cfg.Address), log.Field("dbname", cfg.DBName))
}

func initRedis() {
	viper.SetDefault("redis.max_retry", 3)
	cfg := &clients.RedisConfig{
		Address:  viper.GetString("redis.address"),
		MaxRetry: viper.GetInt("redis.max_retry"),
		Password: viper.GetString("redis.password"),
	}

	var err error
	clients.RedisClient, err = clients.NewRedisClient(cfg)
	if err != nil {
		log.Bg().Fatal("[init-redis] create database connection failed", log.Field("file", viper.ConfigFileUsed()), log.Field("address", cfg.Address), log.Err(err))
		os.Exit(-1)
		return
	}

	log.Bg().Info("[init-redis] create database connection success", log.Field("address", cfg.Address))
}

func initServices() {
	err := services.InitServices()
	if err != nil {
		log.Bg().Fatal("[init-services] process failed", log.Err(err))
		os.Exit(-1)
	}

	log.Bg().Info("[init-services] process success")
}