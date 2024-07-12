package clients

import (
	"fmt"

	"github.com/huynhtrongtien/dove/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/tracing"
)

type MySQLConfig struct {
	Address  string
	DBName   string
	Username string
	Password string
}

var MySQLClient *gorm.DB

func NewMySQLClient(config *MySQLConfig) (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True",
		config.Username, config.Password, config.Address, config.DBName)

	var err error
	client, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		PrepareStmt:                              false,
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	err = client.Use(tracing.NewPlugin())
	if err != nil {
		return nil, err
	}

	client = client.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 auto_increment=1")

	// set data models many to many relationship
	/*
		err = client.SetupJoinTable(&Product{}, "Factories", &ProductFactory{})
		if err != nil {
			return nil, err
		}
	*/

	return client, nil
}

func AutoMigrate() error {

	err := MySQLClient.AutoMigrate(&entities.User{})
	if err != nil {
		return err
	}

	err = MySQLClient.AutoMigrate(&entities.Category{})
	if err != nil {
		return err
	}

	err = MySQLClient.AutoMigrate(&entities.Product{})
	if err != nil {
		return err
	}

	return nil
}
